package main

import (
	_ "aws-case/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run(":80")
}
