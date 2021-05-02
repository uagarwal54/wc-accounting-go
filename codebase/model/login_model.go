package model

import (
	"time"

	"github.com/beego/beego/client/orm"
)

type (
	User struct {
		SrNum            int       `orm:"column(srNum);pk" json:"srnum,omitempty"`
		UserId           string    `orm:"column(userId)" json:"userid,omitempty"`
		UserName         string    `orm:"column(userName)" json:"username"`
		RegistrationDate time.Time `orm:"column(registrationDate)" json:"registrationdate,omitempty"`
		Password         string    `orm:"column(password)" json:"password"`
		FirstLogin       int      `orm:"column(firstLogin)" json:"firstlogin,omitempty"`
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
