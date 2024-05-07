// controllers/furniture_controller.go
package controllers

import (
	"log"
	"net/http"

	"github.com/ViscousGuy/interior-connect-rest-app/models"
	"github.com/astaxie/beego/orm"
	"github.com/beego/beego/v2/server/web"
)

type ProjectController struct {
	web.Controller
}

func (fc *ProjectController) GetAllProjects() {
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
	var allProject []models.Project
	_, err = o.QueryTable(new(models.Project)).Limit(limit, (page-1)*limit).All(&allProject)
	if err != nil {
		log.Printf("Database error: %s", err)
		fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
		fc.ServeJSON()
		return
	}

	// Check if result is empty
	if len(allProject) == 0 {
		fc.Ctx.Output.SetStatus(http.StatusNotFound)
		fc.Data["json"] = map[string]string{"message": "No Project found"}
		fc.ServeJSON()
		return
	}

	// Success response
	fc.Data["json"] = allProject
	fc.ServeJSON()
}

func (fc *ProjectController) GetProjectBySlug() {
    slug := fc.Ctx.Input.Param(":slug")
    // Database operation
    o := orm.NewOrm()
    var project models.Project
    err := o.QueryTable(new(models.Project)).Filter("slug", slug).One(&project)
    if err != nil {
        log.Printf("Database error: %s", err)
        fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
        fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
        fc.ServeJSON()
        return
    }

    if project.Id == 0 {
        fc.Ctx.Output.SetStatus(http.StatusNotFound)
        fc.Data["json"] = map[string]string{"error": "Project not found"}
        fc.ServeJSON()
        return
    }

    fc.Data["json"] = project
    fc.ServeJSON()

}

