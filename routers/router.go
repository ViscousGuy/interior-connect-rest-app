package routers

import (
	"github.com/ViscousGuy/interior-connect-rest-app/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/furnitures", &controllers.FurnitureController{}, "get:GetAllFurnitures") 
	web.Router("/furnitures/:slug", &controllers.FurnitureController{}, "get:GetFurnitureBySlug")
	web.Router("/contractors", &controllers.ContractorController{},"get:GetAllContractors")
	web.Router("/contractors/:slug", &controllers.ContractorController{},"get:GetContractorBySlug")
	web.Router("/projects", &controllers.ProjectController{},"get:GetAllProjects")
	web.Router("/projects/:slug", &controllers.ProjectController{},"get:GetProjectBySlug")
	web.Router("/materials", &controllers.MaterialController{},"get:GetAllMaterials")
}
