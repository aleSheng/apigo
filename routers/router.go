package routers

import (
	"github.com/anlint/apigo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/serv",
			beego.NSRouter("/getone/:id",
				&controllers.ServController{},
				"get:Getone"),
			beego.NSRouter("/getall",
				&controllers.ServController{},
				"get:Getall"),
		),
		beego.NSNamespace("/lint",
			beego.NSRouter("/getone/:id",
				&controllers.LintController{},
				"get:Getone"),
			beego.NSRouter("/getall",
				&controllers.LintController{},
				"get:Getall"),
			beego.NSRouter("/getone/:id/devote",
				&controllers.LintController{},
				"get:Devote"),
		),
	)
	beego.AddNamespace(ns)
}
