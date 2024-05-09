// controllers\project_controller.go
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

// Define CustomProject struct
type CustomProject struct {
    Id            int    `orm:"auto"`
    ProjectName   string `orm:"size(255)"`
    Description   string `orm:"size(255)"`
    City          string `orm:"size(50)"`
    Slug          string `orm:"size(255);unique"`
    Display       bool   `orm:"null"`
    Contractor  *models.Contractor  `orm:"rel(fk)"`
    ProjectImage []*models.ProjectImage `orm:"reverse(many)"`
    // New 
    Furniture []*models.Furniture `orm:"reverse(many)"`
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

	var allProjects []models.Project

	_, err = o.QueryTable(new(models.Project)).RelatedSel("Contractor").All(&allProjects)
	if err != nil {
		log.Printf("Database error: %s", err)
		fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
		fc.ServeJSON()
		return
	} else {

        //Project Images
        for i:= range allProjects{
            _ , err = o.QueryTable("project_image").RelatedSel("Project").Filter("Project__ID",allProjects[i].Id).All(&allProjects[i].ProjectImage)

            if err != nil{
                fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
				fc.ServeJSON()
				return
            }

            // set the ProjectId  field for each ProjectImage object
            for j:= range allProjects[i].ProjectImage{
                allProjects[i].ProjectImage[j].ProjectId = allProjects[i].Id
            }

            // Check if result is empty
            if len(allProjects) == 0 {
                fc.Ctx.Output.SetStatus(http.StatusNotFound)
                fc.Data["json"] = map[string]string{"message": "No Project found"}
                fc.ServeJSON()
                return
            }
        }
		// Prepare response
		responseProjects := make([]map[string]interface{}, len(allProjects))

		for i, project := range allProjects {
			// Create a new CustomProject and copy the fields from project
			customProject := CustomProject{
				Id:            project.Id,
				ProjectName:   project.ProjectName,
				Description:   project.Description,
				City:          project.City,
				Slug:          project.Slug,
				Display:       project.Display,
				Contractor:    project.Contractor,
				ProjectImage:  project.ProjectImage,
			}


			// Loading Furniture Data
			_, err = o.QueryTable("furniture").RelatedSel("Contractor").Filter("Contractor__ID", project.Contractor.Id).All(&customProject.Furniture)
			if err != nil {
				fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
				fc.ServeJSON()
				return
			}
            // Prepare furniture response
			responseFurniture := make([]map[string]interface{}, len(customProject.Furniture))
			for j, furniture := range customProject.Furniture {
				responseFurniture[j] = map[string]interface{}{
					"id":          furniture.Id,
					"name":        furniture.Name,
					"description": furniture.Description,
					"dimensions":  furniture.Dimensions,
					"price":       furniture.Price,
					"slug":        furniture.Slug,
					"display":     furniture.Display,
				}
			}


            // Prepare project response
			responseProjects[i] = map[string]interface{}{
				"id":          customProject.Id,
				"projectName": customProject.ProjectName,
				"description": customProject.Description,
				"city":        customProject.City,
				"slug":        customProject.Slug,
				"display":     customProject.Display,
                "projectImage": customProject.ProjectImage,
				"furniture":   responseFurniture,
				"contractor" : customProject.Contractor,
			}


		}

		// Success response
		fc.Data["json"] = responseProjects
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

    // ProjectImage
    _ , err = o.QueryTable("project_image").RelatedSel("Project").Filter("Project__ID",project.Id).All(&project.ProjectImage)
    if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
		fc.ServeJSON()
		return
	}

    //  setting the ProjectId field for each ProjectImage object
    for j:= range project.ProjectImage{project.ProjectImage[j].ProjectId = project.Id}


    // Create a new CustomProject and copy the fields from project
    customProject := CustomProject{
        Id:            project.Id,
        ProjectName:   project.ProjectName,
        Description:   project.Description,
        City:          project.City,
        Slug:          project.Slug,
        Display:       project.Display,
        Contractor:    project.Contractor,
        ProjectImage:  project.ProjectImage,
    }

    // Loading Furniture Data
    _, err = o.QueryTable("furniture").RelatedSel("Contractor").Filter("Contractor__ID", project.Contractor.Id).All(&customProject.Furniture)
    if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
		fc.ServeJSON()
		return
	}

    // Prepare furniture response
    responseFurniture := make([]map[string]interface{}, len(customProject.Furniture))
    for j, furniture := range customProject.Furniture {
        responseFurniture[j] = map[string]interface{}{
            "id":          furniture.Id,
            "name":        furniture.Name,
            "description": furniture.Description,
            "dimensions":  furniture.Dimensions,
            "price":       furniture.Price,
            "slug":        furniture.Slug,
            "display":     furniture.Display,
        }
    }

    // Prepare project image response
    responseProjectImages := make([]map[string]interface{}, len(customProject.ProjectImage))
    for j, projectImage := range customProject.ProjectImage {
        responseProjectImages[j] = map[string]interface{}{
            "id":         projectImage.Id,
            "imagePath":  projectImage.ImagePath,
            "display":    projectImage.Display,
            "projectId":  projectImage.ProjectId,
        }
    }

    // Prepare project response
    responseProject := map[string]interface{}{
        "id":           customProject.Id,
        "projectName":  customProject.ProjectName,
        "description":  customProject.Description,
        "city":         customProject.City,
        "slug":         customProject.Slug,
        "display":      customProject.Display,
        "furniture":    responseFurniture,
        "projectImage": responseProjectImages,  // Add this line
        "contractor" : customProject.Contractor,
    }

    fc.Data["json"] = responseProject
    fc.ServeJSON()
}
