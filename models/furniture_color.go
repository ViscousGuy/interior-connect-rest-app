package models

import "github.com/astaxie/beego/orm"

type FurnitureColor struct {
    Id           int `orm:"auto"`
    Furniture    *Furniture `orm:"rel(fk)"`
    Color        *Color `orm:"rel(fk)"`
}


func (f *FurnitureColor) TableName() string {

    return "furniture_color"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(FurnitureColor))
}
