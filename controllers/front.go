package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	userCtl := newUserCtl()

	http.Handle("/users", *userCtl)
	http.Handle("/users/", *userCtl)
}

func encodeResponseAsJSON(data interface{}, writer io.Writer) {
	enc := json.NewEncoder(writer)
	enc.Encode(data)
}
