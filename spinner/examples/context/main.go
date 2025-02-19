package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nick-popovic/huh/spinner"
)

func main() {
	action := func() { time.Sleep(5 * time.Second) }
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	go action()
	spinner.New().Context(ctx).Run()
	fmt.Println("Done!")
}
