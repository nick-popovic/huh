package main

import (
	"fmt"

	"github.com/nick-popovic/huh"
)

func main() {
	var file string
	huh.NewForm(
		huh.NewGroup(
			huh.NewFilePicker().
				Picking(true).
				Title("Code").
				Description("Select a .go file").
				AllowedTypes([]string{".go"}).
				Value(&file),
		),
	).WithShowHelp(true).Run()
	fmt.Println(file)
}
