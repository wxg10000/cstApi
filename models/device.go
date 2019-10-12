package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Device struct {
	Id          int64 `orm:"column(id);auto" description:"id"`
	DeviceKey   string
	Enterprise  string
	Time        string
	Ip          string
	PersonCount string
	FaceCount   string
	Version     string
}

func AddDevice(device Device) (Id int64) {
	o := orm.NewOrm()
	Id, err := o.Insert(&device)

	if err != nil {
		fmt.Println("o.Insert err:", err)
		return
	}

	fmt.Printf("插入了数据，id是%d\n", Id)
	return Id
}

func QueryDevice(deviceKey string) (device Device) {
	o := orm.NewOrm()
	device.DeviceKey = deviceKey
	err := o.Read(&device, "device_key")
	if err != nil {
		fmt.Println("o.Read err:", err)
		return
	}

	return device
}

func UpdateDevice(device Device) () {
	d := QueryDevice(device.DeviceKey)
	if d.Id == 0 {
		fmt.Println("o.Read err:查不到数据")
		AddDevice(device)
		return
	}
	DeleteDevice(device.DeviceKey)
	AddDevice(device)
}

func DeleteDevice(deviceKey string) () {
	o := orm.NewOrm()
	device := QueryDevice(deviceKey)
	n, err := o.Delete(&device)
	if err != nil {
		fmt.Println("o.Read err:", err)
	}
	fmt.Printf("删除了%d 条数据", n)
}
