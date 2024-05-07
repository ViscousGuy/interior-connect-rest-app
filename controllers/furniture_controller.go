package controllers

import (
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/ViscousGuy/interior-connect-rest-app/models"
	"log"
	"net/http"
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
	qs := o.QueryTable(new(models.Furniture)).
		RelatedSel("Contractor", "FurnitureType", "RoomType").
		Limit(limit, (page-1)*limit)
	_, err = qs.All(&allFurniture, "Id", "Name", "Description", "Dimensions", "Price", "Slug", "Display", "Contractor", "FurnitureType", "RoomType")
	if err != nil {
		log.Printf("Database error: %s", err)
		fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
		fc.ServeJSON()
		return
	}

	// Load the related entities
	for _, f := range allFurniture {
		o.LoadRelated(&f, "FurnitureColor")
		o.LoadRelated(&f, "FurnitureMaterial")
		o.LoadRelated(&f, "FurnitureImage")
	}

	// Transform the data to include related entities
	type FurnitureResponse struct {
		models.Furniture
		Contractor        *models.Contractor   `json:"contractor,omitempty"`
		FurnitureType     *models.FurnitureType `json:"furniture_type,omitempty"`
		RoomType          *models.RoomType     `json:"room_type,omitempty"`
		FurnitureColor    []*models.FurnitureColor `json:"furniture_color,omitempty"`
		FurnitureMaterial []*models.FurnitureMaterial `json:"furniture_material,omitempty"`
		FurnitureImage    []*models.FurnitureImage `json:"furniture_image,omitempty"`
	}

	var response []FurnitureResponse
	for _, f := range allFurniture {
		fr := FurnitureResponse{
			Furniture:       f,
			Contractor:      f.Contractor,
			FurnitureType:   f.FurnitureType,
			RoomType:        f.RoomType,
			FurnitureColor:  f.FurnitureColor,
			FurnitureMaterial: f.FurnitureMaterial,
			FurnitureImage:  f.FurnitureImage,
		}

		response = append(response, fr)
	}

	// Check if result is empty
	if len(allFurniture) == 0 {
		fc.Ctx.Output.SetStatus(http.StatusNotFound)
		fc.Data["json"] = map[string]string{"message": "No furniture found"}
		fc.ServeJSON()
		return
	}

	// Success response
	fc.Data["json"] = response
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





