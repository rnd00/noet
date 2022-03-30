package controllers

import (
	"io"
	"net/http"
)

func TestWrite(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "testwrite called")
}
