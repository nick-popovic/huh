package main

import "github.com/nick-popovic/huh"

func main() {
	f := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Dynamic Help"),
			huh.NewInput().Title("Dynamic Help"),
			huh.NewInput().Title("Dynamic Help"),
		),
	)
	f.Run()
}
