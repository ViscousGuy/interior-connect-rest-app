package models

import "github.com/astaxie/beego/orm"

type Material struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(50)"`
    Display bool
}

func (f *Material) TableName() string {

    return "material"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(Material))
}
