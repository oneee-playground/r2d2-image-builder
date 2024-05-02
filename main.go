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
	defer os.MkdirAll("/tmp", 0755)
	defer os.RemoveAll("/tmp")
	fs := osfs.New("/tmp/repo")

	startJob := time.Now()

	startFetch := time.Now()
	err := git.FetchSource(context.Background(), fs, "oneee-playground/hello-docker", "745b05e587d5f3903c7622d19a4f41e42fbf5a6c")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("fetch took: %s\n", time.Since(startFetch))

	startBuild := time.Now()
	img, err := image.Build(context.Background(), fs.Root())
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("build took: %s\n", time.Since(startBuild))

	startPush := time.Now()
	if err := image.Push(img); err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("push took: %s\n", time.Since(startPush))

	fmt.Printf("job took: %s\n", time.Since(startJob))
}
