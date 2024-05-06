package function

import (
	"context"
	"path/filepath"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/oneee-playground/r2d2-image-builder/config"
	"github.com/oneee-playground/r2d2-image-builder/git"
	"github.com/oneee-playground/r2d2-image-builder/image"
	"github.com/oneee-playground/r2d2-image-builder/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type BuildImageRequest struct {
	Repository string `json:"repository"`
	CommitHash string `json:"commitHash"`
	Platform   string `json:"platform"`
}

func HandleBuildImage(ctx context.Context, req BuildImageRequest) error {
	defer func() {
		if err := util.CleanupFS(); err != nil {
			logrus.Fatal(err)
		}
	}()

	logrus.Info("Creating source directory")
	fs := osfs.New(filepath.Join(config.Tmpfs, "source"))

	logrus.Infof("Fetching source from: %s:%s", req.Repository, req.CommitHash)
	err := git.FetchSource(ctx, fs, req.Repository, req.CommitHash)
	if err != nil {
		return errors.Wrap(err, "failed to fetch source")
	}

	logrus.Info("Building image from source")
	img, err := image.Build(ctx, image.BuildOpts{
		ContextPath: fs.Root(),
		Platform:    req.Platform,
	})
	if err != nil {
		return errors.Wrap(err, "failed to build image")
	}

	logrus.Info("Pushing image")
	err = image.Push(ctx, img, image.PushOpts{
		Auth:         config.RegistryAuth,
		RegistryName: config.RegistryAddr,
		RegistryUser: config.RegistryUser,
		Repository:   config.Repository,
		Tag:          config.DefaultTag,
	})
	if err != nil {
		return errors.Wrap(err, "failed to push image")
	}

	return nil
}
