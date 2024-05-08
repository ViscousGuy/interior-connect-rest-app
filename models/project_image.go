// models\project_image.go
package models

import "github.com/astaxie/beego/orm"

type ProjectImage struct {
    Id         int     `orm:"auto"`
    Project    *Project `orm:"rel(fk)" json:"-"`
    ImagePath  string  `orm:"size(255)"`
    Display    bool
    ProjectId  int     `orm:"-" json:"project_id"`
}


func (f *ProjectImage) TableName() string {

    return "project_image"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(ProjectImage))
}
