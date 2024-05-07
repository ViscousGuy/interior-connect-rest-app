package models

import "github.com/astaxie/beego/orm"

type RoomType struct {
    Id      int     `orm:"auto"`
    Name    string  `orm:"size(50)"`
    Slug    string  `orm:"size(255);unique"`
    Display bool
}


func (f *RoomType) TableName() string {

    return "room_type"   // changing it into lowercase
}



func init() {
    orm.RegisterModel(new(RoomType))
}
