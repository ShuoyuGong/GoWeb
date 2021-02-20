package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64     `orm:"index"`
	TopicLastUserId int64     `orm:"index"`
}

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplayTime       time.Time `orm:"index"`
	ReplayCount      int64
	ReplayLastUserId int64
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Category), new(Topic))
}
