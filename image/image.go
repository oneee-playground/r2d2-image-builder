package image

import (
	"archive/tar"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	docker "github.com/docker/docker/client"
	"github.com/pkg/errors"
)

const (
	baseCPUShare = 100000
	bytesPerMB   = 1000 * 1000

	imageTag          = "temporary-image"
	imageRegistryAddr = "registry.hub.docker.com"
	imageRegistryUser = "oneeonly"
)

var encodedToken = createEncodedAuth()

type BuildOpts struct {
	NumCPU   int64
	MemoryMB int64
	Dir      string
}

func Build(ctx context.Context, client *docker.Client, opts BuildOpts) error {
	archive, err := archiveSource(opts.Dir)
	if err != nil {
		return errors.Wrap(err, "archiving dockerfile")
	}
	archive.Seek(0, 0)
	defer os.Remove(archive.Name())

	res, err := client.ImageBuild(ctx, archive, types.ImageBuildOptions{
		Tags:      []string{createTag()},
		NoCache:   true,
		Remove:    true,
		CPUShares: baseCPUShare * opts.NumCPU,
		Memory:    opts.MemoryMB * bytesPerMB,
	})
	if err != nil {
		return errors.Wrap(err, "building image")
	}

	io.Copy(io.Discard, res.Body)
	res.Body.Close()

	return nil
}

func archiveSource(path string) (*os.File, error) {
	file, err := os.CreateTemp(".", "tar")
	if err != nil {
		return nil, errors.Wrap(err, "creating archive file")
	}

	tw := tar.NewWriter(file)
	defer tw.Close()

	if err := tw.AddFS(os.DirFS(path)); err != nil {
		os.Remove(file.Name())
		return nil, errors.Wrap(err, "wtf")
	}

	return file, nil
}

func PushAndPrune(ctx context.Context, client *docker.Client) error {
	result, err := client.ImageList(ctx, image.ListOptions{
		Filters: filters.NewArgs(filters.Arg("reference", createTag())),
	})
	if err != nil {
		return errors.Wrap(err, "searcing image")
	}

	if len(result) != 1 {
		return fmt.Errorf("num images should be 1, but %d given", len(result))
	}

	id := result[0].ID

	res, err := client.ImagePush(ctx, createTag(), image.PushOptions{
		RegistryAuth: encodedToken,
	})
	if err != nil {
		return errors.Wrap(err, "pushing image")
	}
	io.Copy(io.Discard, res)
	res.Close()

	_, err = client.ImageRemove(ctx, id, image.RemoveOptions{})
	if err != nil {
		return errors.Wrap(err, "removing image")
	}

	_, err = client.ImagesPrune(ctx, filters.NewArgs(filters.Arg("dangling", "true")))
	if err != nil {
		return errors.Wrap(err, "pruning dangling images")
	}

	return nil
}

func createTag() string {
	return fmt.Sprintf("%s/%s/%s", imageRegistryAddr, imageRegistryUser, imageTag)
}

func createEncodedAuth() string {
	conf := registry.AuthConfig{
		Username:      imageRegistryUser,
		Password:      os.Getenv("DOCKERHUB_TOKEN"),
		ServerAddress: imageRegistryAddr,
	}

	b, _ := json.Marshal(conf)
	return base64.URLEncoding.EncodeToString(b)
}
