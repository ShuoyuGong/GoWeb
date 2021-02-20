package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
	_ "web/models"
	_ "web/routers"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:1763034gong@/default?charset=utf8")
}
func main() {
	orm.Debug = true
	orm.DefaultTimeLoc = time.UTC

	// 如果是开发模式， 则显示命令信息
	// 非强制模式下自动建表
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Informational("[orm] Create table err : ", err)
	}
	beego.Run()
}
