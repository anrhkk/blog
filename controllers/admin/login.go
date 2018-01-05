package admin

import (
	"gopkg.in/macaron.v1"
)

func Get(ctx *macaron.Context) {
	ctx.HTML(200, "views/admin/login")
}

func Post(ctx *macaron.Context) {
	ctx.HTML(200, "views/admin/login")
}
