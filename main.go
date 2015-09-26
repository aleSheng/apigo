package main

import (
	_ "github.com/anlint/apigo/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)


func main() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	mode := beego.AppConfig.String("runmode")
	if mode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}

	beego.Run()
}
