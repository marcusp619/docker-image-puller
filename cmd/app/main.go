package main

import (
	"github.com/docker/docker/client"
)

func main() {
	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
}
