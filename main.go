package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/go-git/go-billy/v5/osfs"
	"github.com/oneee-playground/r2d2-image-builder/git"
	"github.com/oneee-playground/r2d2-image-builder/image"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(io.Discard)
	os.MkdirAll("/tmp/kaniko", 0755)
	defer os.RemoveAll("/tmp")
	fs := osfs.New("/tmp/repo")

	startJob := time.Now()

	startFetch := time.Now()
	err := git.FetchSource(context.Background(), fs, "oneee-playground/hello-docker", "ea1c67aa43b30a22f4804c2d6d9fb9b7c65663ea")
	if err != nil {
		panic(err)
	}
	fmt.Printf("fetch took: %s\n", time.Since(startFetch))


	startBuild := time.Now()
	err = image.Build(context.Background(), image.BuildOpts{
		NumCPU:   1,
		MemoryMB: 4000,
		Dir:      fs.Root(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("build took: %s\n", time.Since(startBuild))



	fmt.Printf("job took: %s\n", time.Since(startJob))
}
