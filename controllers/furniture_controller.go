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

	o := orm.NewOrm()
	var furnitures []models.Furniture
	// Load related entities and retrieve all furniture entries
	_, err = o.QueryTable(new(models.Furniture)).
		RelatedSel("FurnitureType", "RoomType", "Contractor").
		All(&furnitures)

	if err != nil {
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture data"}
		fc.ServeJSON()
		return
	} else {
		for i := range furnitures {
			// Loading related colors
			_, err = o.QueryTable("furniture_color").
				RelatedSel("Color").
				Filter("Furniture__Id", furnitures[i].Id).
				All(&furnitures[i].FurnitureColor)
			if err != nil {
				fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture colors"}
				fc.ServeJSON()
				return
			}
			// Set the FurnitureId field for each FurnitureColor object
			for j := range furnitures[i].FurnitureColor {
				furnitures[i].FurnitureColor[j].FurnitureId = furnitures[i].Id
			}

			// Loading related materials
			_, err = o.QueryTable("furniture_material").
				RelatedSel("Material").
				Filter("Furniture__Id", furnitures[i].Id).
				All(&furnitures[i].FurnitureMaterial)
			if err != nil {
				fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
				fc.ServeJSON()
				return
			}
			// Set the FurnitureID field for each FurnitureMaterial object
			for j := range furnitures[i].FurnitureMaterial {
				furnitures[i].FurnitureMaterial[j].FurnitureID = furnitures[i].Id
			}

			// Loading related images
			_, err = o.QueryTable("furniture_image").
				Filter("Furniture__Id", furnitures[i].Id).
				All(&furnitures[i].FurnitureImage)
			if err != nil {
				fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture images"}
				fc.ServeJSON()
				return
			}
			// Set the FurnitureID field for each FurnitureImage object
			for j := range furnitures[i].FurnitureImage {
				furnitures[i].FurnitureImage[j].FurnitureID = furnitures[i].Id
			}

		}
		fc.Data["json"] = furnitures
	}
	fc.ServeJSON()
}

func (fc *FurnitureController) GetFurnitureBySlug() {
	slug := fc.Ctx.Input.Param(":slug")
	// Database operation
	o := orm.NewOrm()
	var furniture models.Furniture
	err := o.QueryTable(new(models.Furniture)).
		RelatedSel("FurnitureType", "RoomType", "Contractor"). // Load related entities
		Filter("slug", slug).
		One(&furniture)
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

	// Loading related colors
	_, err = o.QueryTable("furniture_color").
		RelatedSel("Color").
		Filter("Furniture__Id", furniture.Id).
		All(&furniture.FurnitureColor)
	if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture colors"}
		fc.ServeJSON()
		return
	}

	// Set the FurnitureID field for each FurnitureColor object
	for j := range furniture.FurnitureColor {
		furniture.FurnitureColor[j].FurnitureId = furniture.Id
	}

	// Loading related materials
	_, err = o.QueryTable("furniture_material").
		RelatedSel("Material").
		Filter("Furniture__Id", furniture.Id).
		All(&furniture.FurnitureMaterial)
	if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
		fc.ServeJSON()
		return
	}

	// Set the FurnitureID field for each FurnitureMaterial object
	for j := range furniture.FurnitureMaterial {
		furniture.FurnitureMaterial[j].FurnitureID = furniture.Id
	}

	// Loading related images
	_, err = o.QueryTable("furniture_image").
		Filter("Furniture__Id", furniture.Id).
		All(&furniture.FurnitureImage)
	if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture images"}
		fc.ServeJSON()
		return
	}

	// Set the FurnitureID field for each FurnitureImage object
	for j := range furniture.FurnitureImage {
		furniture.FurnitureImage[j].FurnitureID = furniture.Id
	}

	fc.Data["json"] = furniture
	fc.ServeJSON()
}
