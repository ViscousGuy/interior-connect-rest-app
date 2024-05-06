package models

import "github.com/astaxie/beego/orm"

type Furniture struct {
    Id               int     `orm:"auto"`
    FurnitureTypeId  int
    RoomTypeId       int
    Name             string  `orm:"size(100)"`
    Description      string  `orm:"type(text)"`
    Dimensions       string  `orm:"size(50)"`
    Price            float64 `orm:"digits(10);decimals(2)"`
    ContractorId     int
    Slug             string  `orm:"size(255)"`
    Display          bool
}


func (f *Furniture) TableName() string {
    // Explicitly specifying the table name as "Furniture"
 
    
    // Our databse has "furniture" (with lower case)

    return "furniture"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(Furniture))
}
