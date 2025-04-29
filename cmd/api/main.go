package main

import "github.com/chriswith8/go-scaffolding-suggestion/internal/web"

func main() {
	app := web.NewApplication()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
