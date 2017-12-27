package main

import (
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/bindata"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer(macaron.RenderOptions{
		// 模板文件目录，默认为 "templates"
		Directory: "views",
	}))
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Use(macaron.Static("public",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      public.Asset,
				AssetDir:   public.AssetDir,
				AssetNames: public.AssetNames,
				Prefix:     "",
			}),
		},
	))

	m.Use(macaron.Renderer(macaron.RenderOptions{
		TemplateFileSystem: bindata.Templates(bindata.Options{
			Asset:      templates.Asset,
			AssetDir:   templates.AssetDir,
			AssetNames: templates.AssetNames,
			Prefix:     "",
		}),
	}))
	m.Run()
}
