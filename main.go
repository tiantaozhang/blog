package main

import (
	_ "Mybeego/models"
	_ "Mybeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
