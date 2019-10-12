package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

var (
	ImgRegs map[string]*ImgReg
)

type ImgReg struct {
	Id         int64 `orm:"column(id);auto" description:"id"`
	DeviceKey  string
	PersonId   string
	Name       string
	Time       string
	ImgPath    string
	NewImgPath string
	FaceId     string
	Ip         string
	Feature    string
	FeatureKey string
}

func AddImgReg(imgReg ImgReg) (PersonId string) {
	o := orm.NewOrm()
	_, err := o.Insert(&imgReg)
	if err != nil {
		fmt.Println("o.Insert err:", err)
		return
	}

	return imgReg.PersonId
}
