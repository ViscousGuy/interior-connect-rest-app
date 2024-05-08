// controllers\project_controller.go
package controllers

import (
	"fmt"
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


    // Relationa with other Tables (Contractor , ProjectImage)
	// _, err = o.QueryTable(new(models.Project)).RelatedSel("Contractor" , "ProjectImage").Limit(limit, (page-1)*limit).All(&allProject)
	
    
    // This RelatedSel works with FK 
    _, err = o.QueryTable(new(models.Project)).RelatedSel("Contractor").All(&allProject)
    
  
	if err != nil {
		log.Printf("Database error: %s", err)
		fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
		fc.ServeJSON()
		return
	} else {
        fmt.Println()
        fmt.Println(allProject)
        fmt.Println()


        // Extracting all info from Other Table where ProjectID used as FK
        // ProjectImage
        for i:= range allProject{
            _ , err = o.QueryTable("project_image").RelatedSel("Project").Filter("Project__ID",allProject[i].Id).All(&allProject[i].ProjectImage)

            if err != nil{
                fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
				fc.ServeJSON()
				return
            }

            // set the ProjectId  field for each ProjectImage object
            for j:= range allProject[i].ProjectImage{
                allProject[i].ProjectImage[j].ProjectId = allProject[i].Id
            }

            // Check if result is empty
            if len(allProject) == 0 {
                fc.Ctx.Output.SetStatus(http.StatusNotFound)
                fc.Data["json"] = map[string]string{"message": "No Project found"}
                fc.ServeJSON()
                return
            }

        // Success response
        }
        fc.Data["json"] = allProject
        
    }
    fc.ServeJSON()
}









func (fc *ProjectController) GetProjectBySlug() {
    slug := fc.Ctx.Input.Param(":slug")
    // Database operation
    o := orm.NewOrm()

    var project models.Project // to store data of single
    
    err := o.QueryTable(new(models.Project)).RelatedSel("Contractor").Filter("slug", slug).One(&project)
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

    //  Adding Related Path
    _ , err = o.QueryTable("project_image").RelatedSel("Project").Filter("Project__ID",project.Id).All(&project.ProjectImage)
    if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
		fc.ServeJSON()
		return
	}

    // ?? setting the ProjectId field for each ProjectImage object
    for j:= range project.ProjectImage{project.ProjectImage[j].ProjectId = project.Id}

    fc.Data["json"] = project
    fc.ServeJSON()

}

