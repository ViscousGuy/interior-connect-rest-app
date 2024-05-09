// models\contractor.go
package models

import "github.com/astaxie/beego/orm"

type Contractor struct {
    Id        int    `orm:"auto"`
    Firstname string `orm:"size(255)"`
    Lastname  string `orm:"size(255)"`
    City      string `orm:"size(50)"`
    State     string `orm:"size(50)"`
    Mobile    string `orm:"size(15)"`
    Email     string `orm:"size(255);unique"`
    Slug      string `orm:"size(255);unique"`
    Pincode   string `orm:"size(10)"`
    Verified  bool
    Active    bool
    Display   bool
    
    //Removed
    // Furniture []*Furniture `orm:"reverse(many)"`
    // Project   []*Project `orm:"reverse(many)"`
    
}

func (c *Contractor) TableName() string {
    return "contractor"
}

func init() {
    orm.RegisterModel(new(Contractor))
}
