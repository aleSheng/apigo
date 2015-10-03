package routers

import (
	"github.com/anlint/apigo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/users",
			beego.NSRouter("/getone",
				&controllers.UserController{},
				"get:Getone"),
			beego.NSRouter("/getall",
				&controllers.UserController{},
				"get:Getall"),
		),
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
		),
	)
	beego.AddNamespace(ns)
}
