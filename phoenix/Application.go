package phoenix

import (
	"log"
	"net/http"
)

type Application struct {
	router *Router
}

func (app *Application) Configure() *Application {
	app.router = NewRouter(app)
	return app
}

func (app *Application) WithMiddleware(fn func(router *Router)) *Application {
	fn(app.router)
	return app
}

func (app *Application) WithRouting(fn func(router *Router)) *Application {
	fn(app.router)
	return app
}

func (app *Application) WithExceptions() *Application {
	return app
}

func (app *Application) Handle(port string) {
	log.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(port, app.router.Mux)
}

func NewApplication() *Application {
	return &Application{}
}
