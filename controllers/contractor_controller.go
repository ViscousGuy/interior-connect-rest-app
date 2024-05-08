// controllers\contractor_controller.go

//  Contractors does not have FK , we have to join explicitly Furniture and Project


// controllers\contractor_controller.go
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
	} else {
		// Prepare response
		responseContractors := make([]map[string]interface{}, len(allContractors))

		for i, contractor := range allContractors {
			// Loading Furniture Data
			_, err = o.QueryTable("furniture").RelatedSel("Contractor").Filter("Contractor__ID", contractor.Id).All(&contractor.Furniture)
			if err != nil {
				fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
				fc.ServeJSON()
				return
			}

			// Prepare furniture response
			responseFurniture := make([]map[string]interface{}, len(contractor.Furniture))
			for j, furniture := range contractor.Furniture {
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

			// Loading Project Data
			_, err = o.QueryTable("project").RelatedSel("Contractor").Filter("Contractor__ID", contractor.Id).All(&contractor.Project)
			if err != nil {
				fc.Ctx.Output.SetStatus(500)
				fc.Data["json"] = map[string]interface{}{"error": "Failed to load project data"}
				fc.ServeJSON()
				return
			}

			// Prepare project response
			responseProjects := make([]map[string]interface{}, len(contractor.Project))
			for k, project := range contractor.Project {
				responseProjects[k] = map[string]interface{}{
					"id":          project.Id,
					"name":        project.ProjectName,
					"description": project.Description,
					"city":        project.City,
					"slug":        project.Slug,
					"display":     project.Display,
				}
			}

			// Prepare contractor response
			responseContractors[i] = map[string]interface{}{
				"id":         contractor.Id,
				"firstname":  contractor.Firstname,
				"lastname":   contractor.Lastname,
				"city":       contractor.City,
				"state":      contractor.State,
				"mobile":     contractor.Mobile,
				"email":      contractor.Email,
				"slug":       contractor.Slug,
				"pincode":    contractor.Pincode,
				"verified":   contractor.Verified,
				"active":     contractor.Active,
				"display":    contractor.Display,
				"furniture":  responseFurniture,
				"projects":   responseProjects,
			}
		}

		// Success response
		fc.Data["json"] = responseContractors
	}
	fc.ServeJSON()
}


// This code will create a new response structure 
// that only includes the fields you want to display. 
// It removes the nested Contractor field in 
// the Furniture and Project objects to avoid redundancy 
// and circular references. 
// It also removes the FurnitureType, RoomType, FurnitureColor, 
// FurnitureMaterial, and FurnitureImage fields 
// from the Furniture object, 
// and the ProjectImage field from the Project object





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
        fc.Data["json"] = map[string]string{"error": "Contractor not found"}
        fc.ServeJSON()
        return
    }

	// Loading Furniture Data
	_, err = o.QueryTable("furniture").RelatedSel("Contractor").Filter("Contractor__ID", contractor.Id).All(&contractor.Furniture)
	if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
		fc.ServeJSON()
		return
	}

	// Prepare furniture response
	responseFurniture := make([]map[string]interface{}, len(contractor.Furniture))
	for j, furniture := range contractor.Furniture {
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

	// Loading Project Data
	_, err = o.QueryTable("project").RelatedSel("Contractor").Filter("Contractor__ID", contractor.Id).All(&contractor.Project)
	if err != nil {
		fc.Ctx.Output.SetStatus(500)
		fc.Data["json"] = map[string]interface{}{"error": "Failed to load project data"}
		fc.ServeJSON()
		return
	}

	// Prepare project response
	responseProjects := make([]map[string]interface{}, len(contractor.Project))
	for k, project := range contractor.Project {
		responseProjects[k] = map[string]interface{}{
			"id":          project.Id,
			"name":        project.ProjectName,
			"description": project.Description,
			"city":        project.City,
			"slug":        project.Slug,
			"display":     project.Display,
		}
	}

	// Prepare contractor response
	responseContractor := map[string]interface{}{
		"id":         contractor.Id,
		"firstname":  contractor.Firstname,
		"lastname":   contractor.Lastname,
		"city":       contractor.City,
		"state":      contractor.State,
		"mobile":     contractor.Mobile,
		"email":      contractor.Email,
		"slug":       contractor.Slug,
		"pincode":    contractor.Pincode,
		"verified":   contractor.Verified,
		"active":     contractor.Active,
		"display":    contractor.Display,
		"furniture":  responseFurniture,
		"projects":   responseProjects,
	}

	fc.Data["json"] = responseContractor
    fc.ServeJSON()
}



/* 
for result Refer : https://pastebin.com/cAqqaYyS

*/
// func (fc *ContractorController) GetContractorBySlug() {
// 	slug := fc.Ctx.Input.Param(":slug")
// 	o := orm.NewOrm()
//     var contractor models.Contractor
// 	err := o.QueryTable(new(models.Contractor)).Filter("slug", slug).One(&contractor)
// 	if err != nil {
//         log.Printf("Database error: %s", err)
//         fc.Ctx.Output.SetStatus(http.StatusInternalServerError)
//         fc.Data["json"] = map[string]string{"error": "Database error: " + err.Error()}
//         fc.ServeJSON()
//         return
//     }
// 	if contractor.Id == 0 {
//         fc.Ctx.Output.SetStatus(http.StatusNotFound)
//         fc.Data["json"] = map[string]string{"error": "Contractor not found"}
//         fc.ServeJSON()
//         return
//     }

// 	// Loading Furniture Data
// 	_, err = o.QueryTable("furniture").RelatedSel("Contractor").Filter("Contractor__ID", contractor.Id).All(&contractor.Furniture)
// 	if err != nil {
// 		fc.Ctx.Output.SetStatus(500)
// 		fc.Data["json"] = map[string]interface{}{"error": "Failed to load furniture materials"}
// 		fc.ServeJSON()
// 		return
// 	}

// 	// Loading Project Data
// 	_, err = o.QueryTable("project").RelatedSel("Contractor").Filter("Contractor__ID", contractor.Id).All(&contractor.Project)
// 	if err != nil {
// 		fc.Ctx.Output.SetStatus(500)
// 		fc.Data["json"] = map[string]interface{}{"error": "Failed to load project data"}
// 		fc.ServeJSON()
// 		return
// 	}

// 	fc.Data["json"] = contractor
//     fc.ServeJSON()
// }


