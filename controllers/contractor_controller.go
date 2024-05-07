// controllers/material_controller.go
package controllers

import (
	"log"
	"net/http"

	"github.com/ViscousGuy/interior-connect-rest-app/models"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
)

type ContractorController struct {
	web.Controller
}

func (fc *ContractorController) GetAllContractors() {
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
	var allContractors []models.Contractor
	_, err = o.QueryTable(new(models.Contractor)).Limit(limit, (page-1)*limit).All(&allContractors)
	if err != nil {
		log.Printf("Database error: %s", err)
		fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
		fc.ServeJSON()
		return
	}

	// Check if result is empty
	if len(allContractors) == 0 {
		fc.Ctx.Output.SetStatus(http.StatusNotFound)
		fc.Data["json"] = map[string]string{"message": "No furniture found"}
		fc.ServeJSON()
		return
	}

	// Success response
	fc.Data["json"] = allContractors
	fc.ServeJSON()
}

func (fc *ContractorController) GetContractorBySlug() {
	slug := fc.Ctx.Input.Param(":slug")
	o := orm.NewOrm()
    var contractor models.Contractor
	err := o.QueryTable(new(models.Contractor)).Filter("slug", slug).One(&contractor)
	if err != nil {
        log.Printf("Database error: %s", err)
        fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
        fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
        fc.ServeJSON()
        return
    }
	if contractor.Id == 0 {
        fc.Ctx.Output.SetStatus(http.StatusNotFound)
        fc.Data["json"] = map[string]string{"error": "Furniture not found"}
        fc.ServeJSON()
        return
    }
	fc.Data["json"] = contractor
    fc.ServeJSON()

}

