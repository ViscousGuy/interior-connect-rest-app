package models

import "github.com/astaxie/beego/orm"

type FurnitureMaterial struct {
    Id           int `orm:"auto"`
    Furniture    *Furniture `orm:"rel(fk)"`
    Material     *Material `orm:"rel(fk)"`
}


func (f *FurnitureMaterial) TableName() string {

    return "furniture_material"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(FurnitureMaterial))
}
