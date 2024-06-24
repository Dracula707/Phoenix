package phoenix

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Ok    uint8       `json:"ok"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type Response struct {
	w http.ResponseWriter
}

func (r *Response) Html(value string) *Response {
	r.w.Header().Set("Content-Type", "text/html")
	r.w.Write([]byte(value))
	return r
}

func (r *Response) Err(err interface{}) {
	r.w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&ApiResponse{Ok: 1, Data: nil, Error: err})
	r.w.Write(b)
}

func (r *Response) Json(value interface{}) {
	r.w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(&ApiResponse{Ok: 0, Data: value, Error: nil})
	r.w.Write(b)
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{w}
}
