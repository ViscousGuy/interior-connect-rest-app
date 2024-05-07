package routers

import (
	"github.com/ViscousGuy/interior-connect-rest-app/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/furnitures", &controllers.FurnitureController{}, "get:GetAllFurniture") 
	web.Router("/furnitures/:slug", &controllers.FurnitureController{}, "get:GetFurnitureBySlug")
	web.Router("/materials", &controllers.MaterialController{},"get:GetAllMaterial")
	web.Router("/contractors", &controllers.ContractorController{},"get:GetAllContractors")
	web.Router("/contractors/:slug", &controllers.ContractorController{},"get:GetContractorBySlug")
}
