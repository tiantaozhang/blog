package models

import (
	"Mybeego/models/class"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func init() {

	orm.Debug, _ = beego.AppConfig.Bool("DB::debug")

	switch beego.AppConfig.String("DB::db") {
	case "mysql":
		dbu := beego.AppConfig.String("DB::user")
		dbp := beego.AppConfig.String("DB::pass")
		dbn := beego.AppConfig.String("DB::name")

		orm.RegisterDriver("mysql", orm.DR_MySQL)
		orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%s:%s@/%s?charset=utf8", dbu, dbp, dbn), 10, 10)

	case "sqlite":
	case "mongodb":
	default:
		log.Fatalln(beego.AppConfig.String("DB::db"))
	}

	orm.RegisterModel(new(class.User))
	_ = orm.RunSyncdb("default", false, true)

}
