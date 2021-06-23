package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/api"
	"github.com/DENKweit/gqlgenc/clientgen"
	"github.com/DENKweit/gqlgenc/config"
	"github.com/DENKweit/gqlgenc/generator"
	"os"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cfg: %+v", err.Error())
		os.Exit(2)
	}

	clientGen := api.AddPlugin(clientgen.New(cfg.Query, cfg.Client, cfg.Generate))

	if err := generator.Generate(ctx, cfg, clientGen); err != nil {
		fmt.Fprintf(os.Stderr, "generate: %+v", err.Error())
		os.Exit(4)
	}
}
