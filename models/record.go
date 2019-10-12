package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Record struct {
	Id        int64 `orm:"column(id);auto" description:"id"`
	Ip        string
	DeviceKey string
	PersonId  string
	Time      string
	Type      string
	Path      string
	Data      string
}

func AddRecord(record Record) () {
	o := orm.NewOrm()
	n, err := o.Insert(&record)
	if err != nil {
		fmt.Println("o.Insert err:", err)
		return
	}
	fmt.Printf("插入%d条数据", n)
}
