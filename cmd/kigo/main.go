package main

import (
	"github.com/saromanov/kigo/internal/rest"
)

func main(){
	ctx := context.Background()
	rest.Init(ctx, rest.Config{})
}