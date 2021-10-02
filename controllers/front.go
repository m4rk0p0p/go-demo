package controllers

import "net/http"

func RegisterControllers() {
	userCtl := newUserCtl()

	http.Handle("/users", *userCtl)
	http.Handle("/users/", *userCtl)
}
