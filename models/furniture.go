package models

import "github.com/astaxie/beego/orm"

type Furniture struct {
    FurnitureId     int     `orm:"auto;column(furniture_id)"` // Note the field name change
    Name            string  `orm:"size(100)"` 
    Type            string  `orm:"size(50)"`
    Style           string  `orm:"size(50)"`
    Description     string  `orm:"type(text)"`   
    Dimensions      string  `orm:"size(50)"` 
    Price           float64 `orm:"digits(10);decimals(2)"`
    PrimaryImagePath string `orm:"column(primary_image_path);size(255)"` 
    ContractorId    int     `orm:"column(contractor_id)"` 
}

func (f *Furniture) TableName() string {
    // Explicitly specifying the table name as "Furniture"
    return "Furniture" 
}

func init() {
    orm.RegisterModel(new(Furniture))
}
