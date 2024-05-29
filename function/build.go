package function

import (
	"context"
	"path/filepath"
	"time"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/google/uuid"
	"github.com/oneee-playground/r2d2-image-builder/config"
	"github.com/oneee-playground/r2d2-image-builder/git"
	"github.com/oneee-playground/r2d2-image-builder/image"
	"github.com/oneee-playground/r2d2-image-builder/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type BuildImageRequest struct {
	ID         uuid.UUID `json:"id"`
	Repository string    `json:"repository"`
	CommitHash string    `json:"commitHash"`
	Platform   string    `json:"platform"`
}

type BuildImageResponse struct {
	ID      uuid.UUID     `json:"id"`
	Success bool          `json:"success"`
	Took    time.Duration `json:"took"`
	Extra   string        `json:"extra"`
}

func HandleBuildImage(ctx context.Context, req BuildImageRequest) (*BuildImageResponse, error) {
	defer func() {
		if err := util.CleanupFS(); err != nil {
			logrus.Fatal(err)
		}
	}()

	start := time.Now()

	logrus.Info("Creating source directory")
	fs := osfs.New(filepath.Join(config.Tmpfs, "source"))

	logrus.Infof("Fetching source from: %s:%s", req.Repository, req.CommitHash)
	err := git.FetchSource(ctx, fs, req.Repository, req.CommitHash)
	if err != nil {
		err = errors.Wrap(err, "failed to fetch source")
		return buildResponse(req.ID, since(start), err), nil
	}

	logrus.Info("Building image from source")
	img, err := image.Build(ctx, image.BuildOpts{
		ContextPath: fs.Root(),
		Platform:    req.Platform,
	})
	if err != nil {
		err = errors.Wrap(err, "failed to build image")
		return buildResponse(req.ID, since(start), err), nil
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
		err = errors.Wrap(err, "failed to push image")
		return buildResponse(req.ID, since(start), err), nil
	}

	return buildResponse(req.ID, since(start), nil), nil
}

func buildResponse(id uuid.UUID, took time.Duration, err error) *BuildImageResponse {
	res := BuildImageResponse{ID: id, Took: took, Success: true}
	if err != nil {
		res.Success = false
		res.Extra = err.Error()
	}
	return &res
}

func since(t time.Time) time.Duration {
	return time.Since(t)
}
