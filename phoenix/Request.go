package phoenix

import "net/http"

type Request struct {
	W http.ResponseWriter
	R *http.Request
	A *Application
}

func (r *Request) Html(html string) *Response {
	return NewResponse(r.W).Html(html)
}

func NewRequest(w http.ResponseWriter, r *http.Request, a *Application) *Request {
	return &Request{W: w, R: r, A: a}
}
