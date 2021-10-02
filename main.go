package main

import (
	"net/http"

	"github.com/m4rk0p0p/go-demo/controllers"
)

func main() {
	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
