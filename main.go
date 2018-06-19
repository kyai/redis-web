package main

import (
	_ "redis-web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

