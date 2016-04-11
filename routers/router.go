package routers

import (
	"Mybeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/profile", &controllers.UserController{}, "get:Profile")

	beego.Router("/api/user/profile", &controllers.UserController{}, "get:API_Profile")

	beego.Router("/api/user/testmodel", &controllers.UserController{}, "get:TestModel")

}
