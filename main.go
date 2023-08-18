package main

import (
	"context"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// create a Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	src := client.Git("github.com/kpenfound/greetings-api").Branch("main").Tree()

	builder := client.Container().From("golang:1.21").
		WithMountedCache("/go/pkg/mods", client.CacheVolume("gomods")).
		WithMountedCache("/root/.cache/go-build", client.CacheVolume("gobuild")).
		WithMountedDirectory("/src", src).
		WithWorkdir("/src").
		WithExec([]string{"go", "run", "./ci", "ci"}, dagger.ContainerWithExecOpts{ExperimentalPrivilegedNesting: true})

	_, err = builder.Sync(ctx)
	if err != nil {
		panic(err)
	}
}
