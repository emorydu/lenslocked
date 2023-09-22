package main

import (
	"context"
	"fmt"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, favoriteColorKey, "yellow")
	value := ctx.Value(favoriteColorKey)

	if sVal, ok := value.(string); ok {
		fmt.Println(sVal)
	} else {
		fmt.Println("not a string")
	}

}
