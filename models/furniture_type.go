package models

import "github.com/astaxie/beego/orm"

type FurnitureType struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(20)"`
    Slug    string  `orm:"size(255);unique"`
    Display bool
}

func (f *FurnitureType) TableName() string {

    return "furniture_type"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(FurnitureType))
}
