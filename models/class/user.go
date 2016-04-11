package class

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     string `orm:"pk"`
	Nick   string
	Info   string `orm:"null"`
	Hobby  string `orm:"null"`
	Email  string `orm:"unique"`
	Avator string `orm:"null"`
	Url    string `orm:"null"`

	Followers int
	Following int

	Regtime time.Time `orm:"auto_now_add"`

	Password string
	Private  int
}

const (
	PR_Zero = iota
	PR_live
	PR_login
	PR_post
)

func Testmodel() {

	var u User
	u.Id = "jike"
	u.Nick = "jike"
	u.Email = "123@q.com"
	u.Password = "123"

	o := orm.NewOrm()
	_, err := o.Insert(&u)
	if err != nil {
		fmt.Println(err)
	}
	u1 := User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	u2 := User{Nick: "jike"}
	err = o.Read(&u2, "nick")
	fmt.Println(u2)

	u2.Nick = "jike2"
	_, err = o.Update(&u2)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)
	fmt.Println(u1)

	_, err = o.Delete(&u)
	u1 = User{Id: "jike"}
	err = o.Read(&u1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u1)
	}

}
