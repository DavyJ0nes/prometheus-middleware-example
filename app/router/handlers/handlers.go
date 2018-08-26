package handlers

import (
	"io"
	"net/http"
)

// IndexHandler takes care of requests to /
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"message": "hey hey hey"}`)
}

// OtherHandler takes care of other requests
func OtherHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"message": "something else"}`)
}
