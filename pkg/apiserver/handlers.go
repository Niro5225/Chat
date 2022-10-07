package apiserver

import (
	"io"
	"net/http"
)

func (a *Api_server) handle_hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}
}
