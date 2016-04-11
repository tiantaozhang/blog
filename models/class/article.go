package class

import (
	"time"
)

type Article struct {
	Id      int
	Title   string
	Content string `orm:"type(text)`
	Author  *User
	Replys  int
	Views   int

	Tag string

	Time time.Time `orm:"auto_now_add"`
}
