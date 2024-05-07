package routers

import (
	"github.com/ViscousGuy/interior-connect-rest-app/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/furniture", &controllers.FurnitureController{}, "get:GetAllFurniture") 
	web.Router("/furniture/:id", &controllers.FurnitureController{}, "get:GetFurniture") // New route
	web.Router("/material" ,&controllers.MaterialController{},"get:GetAllMaterial")
}
