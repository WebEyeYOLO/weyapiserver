package routers

import (
	"github.com/astaxie/beego"
	"weyapiserver/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ws", &controllers.WebSocketController{})
}
