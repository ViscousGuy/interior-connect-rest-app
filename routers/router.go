package routers

import (
	"github.com/ViscousGuy/interior-connect-rest-app/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/furnitures", &controllers.FurnitureController{}, "get:GetAllFurniture") 
	web.Router("/materials" ,&controllers.MaterialController{},"get:GetAllMaterial")
}
