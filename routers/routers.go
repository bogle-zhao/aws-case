package routers

import (
	"aws-case/controller"
	"github.com/beego/beego/v2/server/web"
)

func init() {

	nsUser := web.NewNamespace(
		"/user",
		web.NSCtrlGet("/register", (*controller.UserController).Register),
		web.NSCtrlPost("/register", (*controller.UserController).Register),

		web.NSCtrlGet("/login", (*controller.UserController).Login),
		web.NSCtrlPost("/login", (*controller.UserController).Login),
		web.NSCtrlPost("/update", (*controller.UserController).UpdateAvatar),
		web.NSCtrlGet("/logout", (*controller.UserController).Logout),
	)
	web.CtrlGet("/*", (*controller.UserController).Index)
	web.AddNamespace(nsUser)
}
