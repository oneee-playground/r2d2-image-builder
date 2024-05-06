package util

import (
	"os"
	"path/filepath"

	"github.com/oneee-playground/r2d2-image-builder/config"
	"github.com/pkg/errors"
)

func InitFS() error {
	err := os.MkdirAll(filepath.Join(config.Tmpfs, "kaniko"), 0755)
	if err != nil {
		return errors.Wrap(err, "could not create directory for kaniko: %v")
	}

	return nil
}

func CleanupFS() error {
	d, err := os.Open(config.Tmpfs)
	if err != nil {
		return errors.Wrap(err, "could not open /tmp: %v")
	}

	names, err := d.Readdirnames(-1)
	if err != nil {
		return errors.Wrap(err, "could not read children of /tmp")
	}

	for _, name := range names {
		dir := filepath.Join(config.Tmpfs, name)
		if err := os.RemoveAll(dir); err != nil {
			return errors.Wrapf(err, "could not remove %s", dir)
		}
	}

	if err := os.MkdirAll(filepath.Join(config.Tmpfs, "kaniko"), 0755); err != nil {
		return errors.Wrap(err, "could not restore /tmp/kaniko")
	}

	return nil
}
