package routers

import (
	"MYProject/QandA/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/auth"
)

func init() {

	beego.Router("/", &controllers.MainController{})
	ns :=
		beego.NewNamespace("/v1",

			//beego.NSCond(auth.Basic_Cond("xiaoyi", "caicaikan")),
			beego.NSRouter("/list:id([0-9])", &controllers.QAController{}),
		)

	ns1 :=
		beego.NewNamespace("/manager",
			beego.NSBefore(auth.Basic("xiaoyi", "caicaikan")),
			beego.NSRouter("/updown", &controllers.UPController{}),
		)
	beego.AddNamespace(ns, ns1)
	//beego.Router("/list", &controllers.QAController{})
}
