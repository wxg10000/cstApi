package controllers

import (
	"cstApi/models"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type ImgRegController struct {
	beego.Controller
}

func (this *ImgRegController) Post() {
	var imgReg models.ImgReg
	imgReg.DeviceKey = this.GetString("deviceKey")
	imgReg.ImgPath = this.GetString("imgPath")
	imgReg.PersonId = this.GetString("personId")
	imgReg.FaceId = this.GetString("faceId")

	t := time.Now()
	fmt.Println(t)
	imgReg.Time = t.Format("2006-01-02 15:04:05") //2018-07-11 15:10:19

	imgReg.NewImgPath = this.GetString("newImgPath")
	imgReg.Feature = this.GetString("feature")
	imgReg.FeatureKey = this.GetString("featureKey")
	imgReg.Ip = this.GetString("ip")
	models.AddImgReg(imgReg)
}
