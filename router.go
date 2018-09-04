package router

import (
	"net/http"
)

type HttpCallback func(resp http.ResponseWriter, req *http.Request)

type HttpRouter struct {
	http.ServeMux
	tree TrieTree
}

func NewHttpRouter() *HttpRouter {
	hr := HttpRouter{}
	hr.ServeMux.HandleFunc("/", hr.onRouter)
	return &hr
}

func (hr *HttpRouter) onRouter(resp http.ResponseWriter, req *http.Request) {
	url := req.RequestURI
	v := hr.tree.GetValue(url)
	if v == nil {
		resp.WriteHeader(http.StatusInternalServerError)
	}
	cb := v.(HttpCallback)
	cb(resp, req)
}

func (hr *HttpRouter) HandleFunc(url string, cb HttpCallback) error {
	return hr.tree.AddPath(url, cb)
}
