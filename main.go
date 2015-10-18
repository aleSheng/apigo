package main

import (
	"time"
	_ "github.com/anlint/apigo/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

func demo() {
	t1 := time.NewTicker(time.Second * 5)
	t2 := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-t1.C:
			println("5s timer")

		case <-t2.C:
			println("10s timer")
		}
	}
}

func main() {
//	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	mode := beego.AppConfig.String("runmode")
	if mode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
//	go demo()
	beego.Run()
}
