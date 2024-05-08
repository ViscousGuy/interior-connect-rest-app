// models\project.go
package models

import "github.com/astaxie/beego/orm"

type Project struct {
    Id            int    `orm:"auto"`
    ProjectName   string `orm:"size(255)"`
    Description   string `orm:"size(255)"`
    City          string `orm:"size(50)"`
    Slug          string `orm:"size(255);unique"`
    Display       bool   `orm:"null"`
    Contractor  *Contractor  `orm:"rel(fk)"`
    ProjectImage []*ProjectImage `orm:"reverse(many)"`
    ContractorId  int     `orm:"-" json:"contractor_id"`
}

func (p *Project) TableName() string {
    return "project"
}

func init() {
    orm.RegisterModel(new(Project))
}
