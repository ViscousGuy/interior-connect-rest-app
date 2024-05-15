package main

import (
	"fmt"

	_ "github.com/ViscousGuy/interior-connect-rest-app/configs" // Import your configs package
	_ "github.com/ViscousGuy/interior-connect-rest-app/routers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors" // Import the CORS filter
)

func main() {
	// Disable app.conf auto-loading
	web.BConfig.RunMode = "dev" // Also sets development mode

	// Console output
	fmt.Println("Hello World from the Interior Design API (Console Test!)")

	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed methods
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Simple Beego Route
	web.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("Hello World from the Interior Design API!"))
	})
	web.Run()
}
