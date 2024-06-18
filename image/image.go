package image

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/GoogleContainerTools/kaniko/pkg/config"
	"github.com/GoogleContainerTools/kaniko/pkg/executor"
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/pkg/errors"
)

type BuildOpts struct {
	ContextPath string
	Platform    string
}

func Build(ctx context.Context, opts BuildOpts) (v1.Image, error) {
	o := &config.KanikoOptions{
		SrcContext:     opts.ContextPath,
		CustomPlatform: opts.Platform,
		DockerfilePath: filepath.Join(opts.ContextPath, "Dockerfile"),
		SnapshotMode:   "full",
		SingleSnapshot: true,
	}

	image, err := executor.DoBuild(o)
	if err != nil {
		return nil, errors.Wrap(err, "building image")
	}

	return image, nil
}

type PushOpts struct {
	Auth authn.Authenticator

	RegistryName string
	RegistryUser string
	Repository   string
	Tag          string
}

func Push(ctx context.Context, image v1.Image, opts PushOpts) error {
	dst := createDst(opts)

	err := crane.Push(image, dst, crane.WithAuth(opts.Auth))
	if err != nil {
		return errors.Wrap(err, "pushing image to registry")
	}

	if err := crane.Tag(dst, strings.Replace(opts.Tag, "/", "-", 1), crane.WithAuth(opts.Auth)); err != nil {
		return errors.Wrap(err, "tagging image")
	}

	return nil
}

func createDst(opts PushOpts) string {
	return fmt.Sprintf("%s/%s/%s",
		opts.RegistryName,
		opts.RegistryUser,
		opts.Repository,
	)
}
