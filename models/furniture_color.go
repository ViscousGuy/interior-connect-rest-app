package models

import "github.com/astaxie/beego/orm"

type FurnitureColor struct {
    Id           int `orm:"auto"`
    Furniture    *Furniture `orm:"rel(fk)" json:"-"`
    Color        *Color `orm:"rel(fk)"`
    FurnitureId  int     `orm:"-" json:"furniture_id"`
}


func (f *FurnitureColor) TableName() string {

    return "furniture_color"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(FurnitureColor))
}
