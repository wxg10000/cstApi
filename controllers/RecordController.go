package controllers

import (
	"cstApi/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type RecordController struct {
	beego.Controller
}

func (this *RecordController) Post() {
	var Record models.Record
	Record.Path = this.GetString("path")

	Record.Type = this.GetString("type")

	Record.DeviceKey = this.GetString("deviceKey")

	t := this.GetString("time")
	t64, _ := strconv.ParseInt(t, 10, 64)
	tm := time.Unix(t64/1000, 0)
	Record.Time = tm.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19

	Record.Ip = this.GetString("ip")

	Record.PersonId = this.GetString("personId")

	Record.Data = this.GetString("data")

	fmt.Println(Record)
	models.AddRecord(Record)
}
