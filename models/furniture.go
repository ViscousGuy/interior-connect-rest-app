package models

import "github.com/astaxie/beego/orm"

type Furniture struct {
    Id              int     `orm:"auto"`
    FurnitureType   *FurnitureType `orm:"rel(fk)"`
    RoomType        *RoomType `orm:"rel(fk)"`
    Name            string  `orm:"size(100)"`
    Description     string  `orm:"type(text)"`
    Dimensions      string  `orm:"size(50)"`
    Price           float64 `orm:"digits(10);decimals(2)"`
    Contractor      *Contractor `orm:"rel(fk)"`
    Slug            string  `orm:"size(255);unique"`
    Display         bool
    FurnitureColor  []*FurnitureColor `orm:"reverse(many)"` // Add this line
    FurnitureMaterial []*FurnitureMaterial `orm:"reverse(many)"` // Add this line
    FurnitureImage []*FurnitureImage `orm:"reverse(many)"` // Add this line
}


func (f *Furniture) TableName() string {
    // Explicitly specifying the table name as "Furniture"
    // Our databse has "furniture" (with lower case)
    return "furniture"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(Furniture))
}
