package main

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"log"
)

func main() {
	app := pocketbase.New()

	// Custom routes and hooks go here
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// se.Router.GET("/api/custom", handleCustom)
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
