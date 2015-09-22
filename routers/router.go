package routers

import (
	"github.com/anlint/apigo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSRouter("/register",
				&controllers.UserController{},
				"post:Register"),
			beego.NSRouter("/login",
				&controllers.UserController{},
				"post:Login"),
			beego.NSRouter("/logout",
				&controllers.UserController{},
				"post:Logout"),
			beego.NSRouter("/passwd",
				&controllers.UserController{},
				"post:Passwd"),
			beego.NSRouter("/uploads",
				&controllers.UserController{},
				"post:Uploads"),
			beego.NSRouter("/downloads",
				&controllers.UserController{},
				"get:Downloads"),
			beego.NSRouter("/getall",
				&controllers.UserController{},
				"get:Getall"),
		),
		// beego.NSNamespace("/roles",
		// 	beego.NSRouter("/:id",
		// 		&controllers.RoleController{},
		// 		"get:GetOne;put:Put;delete:Delete"),
		// 	beego.NSRouter("/",
		// 		&controllers.RoleController{},
		// 		"get:GetAll;post:Post"),
		// ),
	)
	beego.AddNamespace(ns)
}
