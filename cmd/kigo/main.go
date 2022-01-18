package main

import (
	"context"
	"github.com/saromanov/kigo/internal/rest"
	"github.com/saromanov/kigo/internal/generate/service"
)

func main(){
	ctx := context.Background()

	srv := service.New(service.Config{})
	rest.Init(ctx, rest.Config{
		Address: "localhost:8082",
	}, srv)
}