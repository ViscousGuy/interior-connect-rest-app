// controllers/furniture_controller.go
package controllers

import (
	"log"
	"net/http"

	"github.com/ViscousGuy/interior-connect-rest-app/models"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
)

type FurnitureController struct {
	web.Controller
}

func (fc *FurnitureController) GetAllFurnitures() {
	// Parse query parameters for pagination
	page, err := fc.GetInt("page", 1)
	if err != nil || page < 1 {
		fc.Ctx.Output.SetStatus(http.StatusBadRequest)
		fc.Data["json"] = map[string]string{"error": "invalid page number"}
		fc.ServeJSON()
		return
	}
	limit, err := fc.GetInt("limit", 10)
	if err != nil || limit < 1 {
		fc.Ctx.Output.SetStatus(http.StatusBadRequest)
		fc.Data["json"] = map[string]string{"error": "invalid limit value"}
		fc.ServeJSON()
		return
	}

	// Database operation
	o := orm.NewOrm()
	var allFurniture []models.Furniture
	_, err = o.QueryTable(new(models.Furniture)).Limit(limit, (page-1)*limit).All(&allFurniture)
	if err != nil {
		log.Printf("Database error: %s", err)
		fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
		fc.ServeJSON()
		return
	}

	// Check if result is empty
	if len(allFurniture) == 0 {
		fc.Ctx.Output.SetStatus(http.StatusNotFound)
		fc.Data["json"] = map[string]string{"message": "No furniture found"}
		fc.ServeJSON()
		return
	}

	// Success response
	fc.Data["json"] = allFurniture
	fc.ServeJSON()
}

func (fc *FurnitureController) GetFurnitureBySlug() {
    slug := fc.Ctx.Input.Param(":slug")
    // Database operation
    o := orm.NewOrm()
    var furniture models.Furniture
    err := o.QueryTable(new(models.Furniture)).Filter("slug", slug).One(&furniture)
    if err != nil {
        log.Printf("Database error: %s", err)
        fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
        fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
        fc.ServeJSON()
        return
    }

    if furniture.Id == 0 {
        fc.Ctx.Output.SetStatus(http.StatusNotFound)
        fc.Data["json"] = map[string]string{"error": "Furniture not found"}
        fc.ServeJSON()
        return
    }

    fc.Data["json"] = furniture
    fc.ServeJSON()

}





