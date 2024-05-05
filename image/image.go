package image

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/GoogleContainerTools/kaniko/pkg/config"
	"github.com/GoogleContainerTools/kaniko/pkg/executor"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
)

const (
	imageTag          = "temporary-image:latest"
	imageRegistryAddr = "docker.io"
	imageRegistryUser = "oneeonly"
)

func Build(ctx context.Context, path string) (v1.Image, error) {
	o := &config.KanikoOptions{
		DockerfilePath: filepath.Join(path, "Dockerfile"),
		SrcContext:     path,
		SnapshotMode:   "full",
		SingleSnapshot: true,
		NoPushCache:    true,
		CustomPlatform: "linux/arm64/v8",
	}

	image, err := executor.DoBuild(o)
	if err != nil {
		return nil, errors.Wrap(err, "building image")
	}

	return image, nil
}

func Push(ctx context.Context, image v1.Image) error {
	tag, err := name.NewTag(createTag())
	if err != nil {
		return errors.Wrap(err, "creating new tag")
	}

	err = crane.Push(image, tag.Name(),
		crane.WithAuth(authn.FromConfig(authn.AuthConfig{
			Username: imageRegistryUser,
			Password: os.Getenv("DOCKERHUB_SECRET"),
		})),
	)
	if err != nil {
		return errors.Wrap(err, "pushing image to registry")
	}

	return nil
}

func createTag() string {
	return fmt.Sprintf("%s/%s/%s", imageRegistryAddr, imageRegistryUser, imageTag)
}
