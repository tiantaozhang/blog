package controllers

import (
	"Mybeego/models/class"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Profile() {

	u.Data["userid"] = "blog1号"
	u.Data["hobby"] = []string{"chess", "code"}

	u.TplNames = "user/profile.html"
}

func (u *UserController) TestModel() {

	class.Testmodel()

	u.Data["json"] = map[string]interface{}{
		"testmodle": true,
	}

	u.ServeJson()
}
