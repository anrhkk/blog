package main

import (
	"blog/static/assets"
	"blog/static/views"
	"github.com/go-macaron/bindata"
	"github.com/go-macaron/pongo2"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"gopkg.in/macaron.v1"
	"net/http"
	"time"
)

func app(state overseer.State) {
	m := macaron.Classic()
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
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Data["Name"] = "jeremy"
		ctx.HTML(200, "views/index") // 200 为响应码
	})
	http.Serve(state.Listener, m)
}

func main() {
	overseer.Run(overseer.Config{
		Program:   app,
		Address:   "127.0.0.1:8888",
		NoRestart: true,
		Fetcher: &fetcher.File{
			Path:     "blog",
			Interval: 1 * time.Second,
		},
	})
}
