package controllers

import (
	"cstApi/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type HeartbeatController struct {
	beego.Controller
}

func (this *HeartbeatController)Post()  {

	var device models.Device
	device.DeviceKey = this.GetString("deviceKey")

	t := this.GetString("time")
	t64, _ := strconv.ParseInt(t, 10, 64)
	tm := time.Unix(t64/1000, 0)
	device.Time = tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19


	device.Ip = this.GetString("ip")


	device.PersonCount = this.GetString("personCount")

	device.FaceCount = this.GetString("faceCount")

	device.Version = this.GetString("version")

	fmt.Println(device)
	models.UpdateDevice(device)


}

