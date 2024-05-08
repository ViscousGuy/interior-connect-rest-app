package models

import "github.com/astaxie/beego/orm"

type FurnitureMaterial struct {
    Id           int `orm:"auto"`
    Furniture    *Furniture `orm:"rel(fk)" json:"-"`
    Material     *Material `orm:"rel(fk)"`
    FurnitureID  int        `orm:"-" json:"furniture_id"`
}


func (f *FurnitureMaterial) TableName() string {

    return "furniture_material"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(FurnitureMaterial))
}
