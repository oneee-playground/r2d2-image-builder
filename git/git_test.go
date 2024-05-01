package git_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/oneee-playground/r2d2-image-builder/git"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchSource(t *testing.T) {
	dir, err := os.MkdirTemp(".", "git")
	require.Nil(t, err)
	fmt.Println(dir)

	defer os.RemoveAll(dir)

	fs := osfs.New(dir)

	err = git.FetchSource(context.Background(), fs, "oneee-playground/hello-docker", "054f12adb13e63ca238718d2a7a4693aedf9e8fe")
	require.Nil(t, err)

	info, err := fs.Stat("Dockerfile")
	assert.Nil(t, info)
	assert.Error(t, err)
}
