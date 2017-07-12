package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"fmt"
	"github.com/kataras/iris/view"
	"github.com/go-resty/resty"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Gallery struct {
	item.Item

	Image []string `json:"image"`
}

type Result struct {
	Data []Gallery `json:"data"`
}

func main() {
	app := iris.New()
	app.StaticServe("views/public", "/public")
	app.StaticWeb("/js", "./views/js")
	app.StaticWeb("/fonts", "./views/fonts")
	app.StaticWeb("/css", "./views/css")
	app.AttachView(view.HTML("./views", ".html").Layout("layout.html").Reload(true))
	app.Get("/gallery/{id}", func(ctx context.Context) {
		Res := Result{}
		apiUrl := "http://yourlaptop.in.ua:8080"
		_, err := resty.R().
			SetQueryString("type=Gallery&id=" + ctx.Params().Get("id")).
			SetResult(&Res).Get(apiUrl+"/api/content")


		if err != nil {
			fmt.Println(err)
		}
		images := []string{}
		for _, d := range Res.Data {
			for _, img := range d.Image {
				images = append(images, img)
			}
		}
		ctx.ViewData("Api",apiUrl)
		ctx.ViewData("Images",images)
		ctx.View("gallery.html")
		fmt.Println(images)
	})

	err := app.Run(iris.Addr(":8081"))
	fmt.Println(err)
}
