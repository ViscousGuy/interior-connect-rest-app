package models

import "github.com/astaxie/beego/orm"

type Color struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(50)"`
    Display bool
}

func (f *Color) TableName() string {

    return "color"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(Color))
}
