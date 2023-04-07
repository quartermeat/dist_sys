package main

import (
	_ "image/png"

	"github.com/quartermeat/dist_sys/go_backend/app"
)

func main() {
	app.ListenAndServer()
}
