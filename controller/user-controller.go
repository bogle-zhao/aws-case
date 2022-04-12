package controller

import (
	"aws-case/entity"
	"aws-case/log"
	"aws-case/service"
	"github.com/beego/beego/v2/server/web"
)

var accountService service.AccountService

func init() {
	accountService = service.NewAccountService()
}

type UserController struct {
	web.Controller
}

func (this *UserController) Index() {
	account := this.GetSession("account")
	if account == nil {
		this.Data["title"] = "登陆"
		this.TplName = "user/login.html"
		return
	}
	this.pageResult(account.(*entity.Account), nil)
}

func (this *UserController) Register() {
	account := this.GetSession("account")
	if account != nil {
		this.pageResult(account.(*entity.Account), nil)
		return
	}
	method := this.Ctx.Request.Method
	this.Data["title"] = "注册"
	if method == "GET" {
		this.TplName = "user/register.html"
	} else if method == "POST" {
		file, header, err := this.GetFile("avatar")
		if err == nil {
			defer file.Close()
		}
		username := this.GetString("username")
		password := this.GetString("password")
		if username == "" || password == "" || err != nil {
			log.Error("register avatar error", err)
			this.Data["error"] = "Incomplete information submitted"
			this.TplName = "user/register.html"
			return
		}
		account, err := accountService.Register(username, password, header.Filename, &file)
		this.pageResult(account, err)
	}
}

func (this *UserController) pageResult(account *entity.Account, err error) {
	if err != nil {
		this.Data["title"] = "登陆"
		this.Data["error"] = err
		this.TplName = "user/login.html"
		return
	}
	this.Data["title"] = "用户信息"
	this.SetSession("account", account)
	this.Data["account"] = *account
	this.TplName = "index.html"
}

func (this *UserController) Login() {
	account := this.GetSession("account")
	if account != nil {
		this.pageResult(account.(*entity.Account), nil)
		return
	}

	method := this.Ctx.Request.Method
	this.Data["title"] = "登陆"

	if method == "GET" {
		this.TplName = "user/login.html"
	} else if method == "POST" {
		username := this.GetString("username")
		password := this.GetString("password")
		account, err := accountService.Login(username, password)
		this.pageResult(account, err)
	}
}

func (this *UserController) UpdateAvatar() {
	method := this.Ctx.Request.Method
	this.Data["title"] = "用户信息"
	if method == "GET" {
		account := this.GetSession("account")
		if account == nil {
			this.Data["title"] = "登陆"
			this.TplName = "user/login.html"
			return
		}
		this.pageResult(account.(*entity.Account), nil)
	} else if method == "POST" {
		username := this.GetString("username")
		file, header, err := this.GetFile("avatar")
		if file == nil {
			account := this.GetSession("account")
			if account == nil {
				this.Data["title"] = "登陆"
				this.TplName = "user/login.html"
				return
			}
			this.pageResult(account.(*entity.Account), nil)
			this.Data["error"] = "Incomplete information submitted"
			return
		}
		if username == "" {
			this.Data["error"] = "Incomplete information submitted"
			this.TplName = "user/login.html"
			return
		}
		account, err := accountService.UpdateAvatar(username, header.Filename, &file)
		this.pageResult(account, err)
	}
}

func (this *UserController) Logout() {
	this.Data["title"] = "登陆"
	this.DelSession("account")
	this.TplName = "user/login.html"
}
