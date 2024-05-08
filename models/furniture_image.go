package models

import "github.com/astaxie/beego/orm"

type FurnitureImage struct {
    Id          int     `orm:"auto"`
    ImagePath   string  `orm:"size(255)"`
    Furniture   *Furniture `orm:"rel(fk)" json:"-"`
    FurnitureID int        `orm:"-" json:"furniture_id"`
}


func (f *FurnitureImage) TableName() string {

    return "furniture_image"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(FurnitureImage))
}
