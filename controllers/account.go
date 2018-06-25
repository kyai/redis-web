package controllers

import (
	"redis-web/redis"
)

type AccountController struct {
	BaseController
}

// 登录
func (this *AccountController) SignIn() {
	this.TplName = "account/sign_in.html"
}

func (this *AccountController) SignInDo() {}

// 注册
func (this *AccountController) SignUp() {
	this.TplName = "account/sign_up.html"
}

func (this *AccountController) SignUpDo() {
	username := this.GetString("username")
	password := this.GetString("password")

	conn := redis.RedisPool.Get()
	defer conn.Close()
	prefix := "redisweb_"
	_, err1 := conn.Do("SET", prefix+"username", username)
	_, err2 := conn.Do("SET", prefix+"password", password)
	if err1 != nil || err2 != nil {
		this.JSON(102, "sign up error", nil)
		return
	}
	this.JSON(101, "success", nil)
}
