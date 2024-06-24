package main

import (
	ph "github.com/dracula707/phoenix"
)

func main() {
	app := ph.NewApplication().Configure()

	app.WithRouting(func(router *ph.Router) {
		router.Get("/", func(req *ph.Request) *ph.Response {
			return req.Html("<h1>Hello, World!</h1>")
		})

		router.Attach().Get("/test", func(req *ph.Request) *ph.Response {
			return req.Html("<h1>Hello, World!</h1>")
		})
	})

	app.Handle(":8080")
}
