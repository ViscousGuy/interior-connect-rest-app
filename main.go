package main

import (
	"fmt"

	_ "github.com/ViscousGuy/interior-connect-rest-app/configs" // Import your configs package
	_ "github.com/ViscousGuy/interior-connect-rest-app/routers"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {
	// Disable app.conf auto-loading 
	web.BConfig.RunMode = "dev" // Also sets development mode

	// Console output
	fmt.Println("Hello World from the Interior Design API (Console Test!)")

	// Simple Beego Route 
	web.Get("/", func(ctx *context.Context) { 
		ctx.Output.Body([]byte("Hello World from the Interior Design API!"))
	})
	web.Run()
}
