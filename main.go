package main

import (
	"context"
	"path/filepath"

	"github.com/docker/docker/client"
	"github.com/oneee-playground/r2d2-image-builder/image"
)

func main() {
	client, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}

	// fs := osfs.New("./tmp")
	// defer os.RemoveAll(fs.Root())

	// err = git.FetchSource(context.Background(), fs, "oneee-playground/hello-docker", "dc94744b9debac7fb14b164d1775f7e1423ad1a0")
	// if err != nil {
	// 	panic(err)
	// }

	dir, _ := filepath.Abs("./tmp")

	err = image.Build(context.Background(), client, image.BuildOpts{
		NumCPU:   1,
		MemoryMB: 4000,
		Dir:      dir,
	})
	if err != nil {
		panic(err)
	}

	err = image.PushAndPrune(context.Background(), client)
	if err != nil {
		panic(err)
	}
}
