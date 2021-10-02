package controllers

import (
	"net/http"
	"regexp"
)

type UserCtl struct {
	userIdPattern *regexp.Regexp
}

func newUserCtl() *UserCtl {
	return &UserCtl{
		userIdPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (ctl UserCtl) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Foobar"))
}
