package controllers

import (
	"github.com/go-macaron/binding"
	"blog/models"
)

func SignIn(ctx *Context, formErr binding.Errors, user models.User) {
	if ctx.Req.Method == "POST" {
		ctx.SetFormError(formErr)
		password := Md5(user.Password)
		user := &models.User{Username: user.Username, Password: password}
		if !ctx.HasError() {
			if has, err := user.Exist(); has {
				PanicIf(err)
				if user.Active == 0 {
					ctx.AddError("用户已被禁用")
					ctx.HTML(200, "views/signin", ctx)
					return
				}
				ctx.Redirect("/admin/dashboard")
			} else {
				ctx.AddError("用户名或密码不正确")
			}
		}
	}
	ctx.HTML(200, "views/signin", ctx)
}
