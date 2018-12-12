package AppController

import "github.com/astaxie/beego"

type AppController struct {
	beego.Controller
}

func (c *AppController) Index() {

	c.Layout = "layout/layout.html"
	c.TplName = "app/index.html"
}
