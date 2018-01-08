package controllers

import (
	. "blog/common"
	"blog/models"
	"github.com/go-macaron/binding"
)

func SignUp(ctx *Context, formErr binding.Errors, user models.User) {
	if ctx.Req.Method == "POST" {
		ctx.SetFormError(formErr)
		if !ctx.HasFormError() {
			dbUser := models.User{Username: user.Username, Password: user.Password, Email: user.Email}
			if exist, err := dbUser.ExistUsername(); exist {
				PanicIf(err)
				ctx.AddFormError("username", "用户名已经存在")
			}
			if exist, err := dbUser.ExistEmail(); exist {
				PanicIf(err)
				ctx.AddFormError("email", "邮箱已经存在")
			}
			if !ctx.HasError() {
				dbUser.Password = Md5(user.Password)
				err := dbUser.Insert()
				PanicIf(err)
				ctx.AddMessage("注册成功！")
			}
		}
	}
	ctx.HTML(200, "views/signup", ctx)
}
