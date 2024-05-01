package image

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/GoogleContainerTools/kaniko/pkg/config"
	"github.com/GoogleContainerTools/kaniko/pkg/executor"
	"github.com/pkg/errors"
)

const (
	baseCPUShare = 100000
	bytesPerMB   = 1000 * 1000

	imageTag          = "temporary-image"
	imageRegistryAddr = "registry.hub.docker.com"
	imageRegistryUser = "oneeonly"
)

type BuildOpts struct {
	NumCPU   int64
	MemoryMB int64
	Dir      string
}

func Build(ctx context.Context, opts BuildOpts) error {
	o := &config.KanikoOptions{
		DockerfilePath: filepath.Join(opts.Dir, "Dockerfile"),
		IgnoreVarRun:   true,
		NoPush:         true,
		SrcContext:     opts.Dir,
		SnapshotMode:   "full",
		CustomPlatform: "linux/amd64",
	}

	image, err := executor.DoBuild(o)
	if err != nil {
		return errors.Wrap(err, "building image")
	}

	fmt.Println(image)

	return nil
}

func createTag() string {
	return fmt.Sprintf("%s/%s/%s", imageRegistryAddr, imageRegistryUser, imageTag)
}
