package main

import (
	routes "github.com/argilapp/core/router"
)

func main() {
	r := routes.SetupRouter()
	r.Run("0.0.0.0:8081")
}
