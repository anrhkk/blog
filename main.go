package main

import (
	"blog/static/assets"
	"blog/static/views"
	"github.com/go-macaron/bindata"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/session"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/binding"
	"gopkg.in/macaron.v1"
	"blog/controllers/admin"
	"blog/controllers"
	"blog/models"
	"blog/common"
)

func init() {
	common.InitConfig()
}

func main() {
	m := macaron.Classic()
	m.Use(common.InitContext())
	m.Use(pongo2.Pongoer(pongo2.Options{
		// 模板文件目录，默认为 "templates"
		Directory: "views",
		TemplateFileSystem: bindata.Templates(bindata.Options{
			Asset:      views.Asset,
			AssetDir:   views.AssetDir,
			AssetNames: views.AssetNames,
		}),
	}))
	m.Use(macaron.Static("assets",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      assets.Asset,
				AssetDir:   assets.AssetDir,
				AssetNames: assets.AssetNames,
			}),
		},
	))
	m.Use(cache.Cacher())
	m.Use(captcha.Captchaer())
	m.Use(session.Sessioner(session.Options{
		CookieName: "session_id",
	}))
	m.Use(csrf.Csrfer())
	m.Map(models.SetEngine())
	//路由
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "views/index")
	})
	m.Route("/signup", "GET,POST", binding.Form(models.UserSignUpForm{}), controllers.SignUp)
	m.Route("/signin", "GET,POST", binding.Form(models.UserSignInForm{}), controllers.SignIn)
	//admin
	m.Group("/admin", func() {
		m.Route("/login", "GET,POST", admin.Get)
	})
	m.Run("127.0.0.1", 8888)
}
