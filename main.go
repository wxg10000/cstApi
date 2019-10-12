package main

import (
	"cstApi/models"
	_ "cstApi/models"
	_ "cstApi/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


var O orm.Ormer

func init() {
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", "root:mysql@tcp(127.0.0.1:3306)/cst?charset=utf8")
	//注册model
	orm.RegisterModel(new(models.Device), new(models.ImgReg),new(models.Record),new(models.Person))
	//自动建表
	orm.RunSyncdb("default", false, true)
	//初始化ormer
	O = orm.NewOrm()
	//初始化数据
	//datainit()

}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
