package models

import "github.com/astaxie/beego/orm"

type ProjectImage struct {
    Id         int     `orm:"auto"`
    Project    *Project `orm:"rel(fk)"`
    ImagePath  string  `orm:"size(255)"`
    Display    bool
}


func (f *ProjectImage) TableName() string {

    return "project_image"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(ProjectImage))
}
