package routers

import (
	"LogManager/controllers/AppController"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/index", &AppController.AppController{}, "*:Index")
}
