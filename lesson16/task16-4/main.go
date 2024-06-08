package main

import (
	"context"
	"fmt"
)

type ctxKey int

func main() {
	ctx := context.Background()
	var key ctxKey = 1
	ctx = context.WithValue(ctx, key, "some value1")
	key = 2
	ctx = context.WithValue(ctx, key, "some value2")

	run(ctx)
}

func run(ctx context.Context) {
	var key ctxKey = 1
	fmt.Println("some key1:", ctx.Value(key))
	key = 2
	fmt.Println("some key2:", ctx.Value(key))
}
