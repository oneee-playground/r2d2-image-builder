package main

import (
	"context"
	"os"
	"path/filepath"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/oneee-playground/r2d2-image-builder/git"
	"github.com/oneee-playground/r2d2-image-builder/image"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger

	tmpfs = os.TempDir()
)

func init() {
	// Keep it this way for later use.
	logger = logrus.New()

	// Disable kankio logger.
	// TODO: Change output dst to show build logs to user.
	// logrus.SetOutput(io.Discard)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	err := os.MkdirAll(filepath.Join(tmpfs, "kaniko"), 0755)
	if err != nil {
		logger.Fatalf("could not create directory for kaniko: %v", err)
	}

	lambda.Start(HandleBuildImage)
}

type BuildImageRequest struct {
	Repository string `json:"repository"`
	CommitHash string `json:"commitHash"`
}

func HandleBuildImage(ctx context.Context, req BuildImageRequest) {
	defer cleanupFS()

	logger.Info("creating source directory")
	fs := osfs.New(filepath.Join(tmpfs, "source"))

	logger.Infof("fetching source from: %s:%s", req.Repository, req.CommitHash)
	err := git.FetchSource(ctx, fs, req.Repository, req.CommitHash)
	if err != nil {
		logger.Errorf("failed to fetch source: %v", err)
		return
	}

	logger.Info("building image from source")
	img, err := image.Build(ctx, fs.Root())
	if err != nil {
		logger.Errorf("failed to build image: %v", err)
		return
	}

	logger.Info("pushing image")
	if err := image.Push(ctx, img); err != nil {
		logger.Errorf("failed to push image : %v", err)
		return
	}
}

func cleanupFS() {
	d, err := os.Open(tmpfs)
	if err != nil {
		logger.Fatalf("could not open /tmp: %v", err)
	}

	names, err := d.Readdirnames(-1)
	if err != nil {
		logger.Fatalf("could not read children of /tmp: %v", err)
	}

	for _, name := range names {
		dir := filepath.Join(tmpfs, name)
		if err := os.RemoveAll(dir); err != nil {
			logger.Fatalf("could not remove %s: %v", dir, err)
		}
	}

	if err := os.MkdirAll(filepath.Join(tmpfs, "kaniko"), 0755); err != nil {
		logger.Fatalf("could not restore /tmp/kaniko: %v", err)
	}
}
