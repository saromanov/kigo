package main

import (
	"context"
	"github.com/saromanov/kigo/internal/rest"
)

func main(){
	ctx := context.Background()
	rest.Init(ctx, rest.Config{})
}