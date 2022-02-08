package main

import (
	"net/http"

	"github.com/PabloRS/collections/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe("localhost:3000", nil)
}
