package main

import (
	"fmt"

	"github.com/nick-popovic/huh/spinner"
)

func main() {
	_ = spinner.New().Title("Loading").Accessible(true).Run()
	fmt.Println("Done!")
}
