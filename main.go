package main

import (
	"fmt"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {
	// Console output
	fmt.Println("Hello World from the Interior Design API (Console Test!)")

	// Simple Beego Route 
	web.Get("/", func(ctx *context.Context) { 
		ctx.Output.Body([]byte("Hello World from the Interior Design API!"))
	})
	
	web.Run()
}
