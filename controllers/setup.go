package controllers

import (
	"redis-web/redis"
)

type SetupController struct {
	BaseController
}

func (this *SetupController) Index() {
	this.TplName = "setup/index.html"
}

func (this *SetupController) CheckConn() {
	host := this.GetString("host")
	port := this.GetString("port")

	redis.InitRedisPool(host+":"+port, "", 3, 5, 240)
	conn := redis.RedisPool.Get()
	defer conn.Close()
	if _, err := conn.Do("PING"); err != nil {
		this.JSON(102, err.Error(), nil)
		return
	}

	this.JSON(101, "success", nil)
}
