package main

import (
	"log"

	"github.com/Psakine/auth/internal/v1/app"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	if err := a.Run(); err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
