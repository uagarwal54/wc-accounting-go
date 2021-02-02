package model

import (
	"time"

	"github.com/beego/beego/client/orm"
)

type (
	User struct {
		SrNum            int       `orm:"column(srNum);pk"`
		UserId           string    `orm:"column(userId)"`
		UserName         string    `orm:"column(userName)"`
		RegistrationDate time.Time `orm:"column(registrationDate)"`
		Password         string    `orm:"column(password)"`
		FirstLogin       bool      `orm:"column(firstLogin)"`
	}
)

func init() {
	orm.RegisterModel(new(User))
}

func ReadAllUserData() (users []User, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("user").All(&users)
	return
}

func ReadUserDataWithUsernamePassword(username, password string) (user User, err error) {
	// orm.Debug = true
	o := orm.NewOrm()
	err = o.QueryTable("user").Filter("userName", username).Filter("password", password).One(&user) // For every combination of username and password a single record must be there.
	return
}
