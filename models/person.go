package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Person struct {
	Id         int64 `orm:"column(id);auto" description:"id"`
	PersonId   string
	IdcardNum  string
	Name       string
	CreateTime string
	Path       string
}

func AddPerson(person Person) () {
	o := orm.NewOrm()
	n, err := o.Insert(&person)
	if err != nil {
		fmt.Println("o.Insert err:", err)
		return
	}
	fmt.Printf("插入%d条数据", n)
}
