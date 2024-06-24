package phoenix

import "net/http"

type Middleware func(http.Handler) http.Handler

type Router struct {
	Mux *http.ServeMux
	app *Application
}

func NewRouter(app *Application) *Router {
	return &Router{Mux: http.NewServeMux(), app: app}
}

func (router *Router) Attach(middlewares ...Middleware) *Router {
	return router
}

func (router *Router) Group(pattern string) *Router {
	return router
}

func (router *Router) Get(path string, handler func(request *Request) *Response) {
	router.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		req := NewRequest(w, r, router.app)
		handler(req)
	})
}
