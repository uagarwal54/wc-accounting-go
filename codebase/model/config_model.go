package model

import (
	"github.com/beego/beego/client/orm"
)

type (
	Config struct {
		SrNum       int    `orm:"column(srNum);pk" json:"srnum"`
		ConfigKey   string `orm:"column(configKey)" json:"configKey"`
		ConfigValue string `orm:"column(configValue)" json:"configValue"`
	}
)

var ConfigMap orm.Params

func init() {
	orm.RegisterModel(new(Config))
}

func PopulateConfigMap() (err error) {
	ConfigMap = make(orm.Params)
	o := orm.NewOrm()
	_, err = o.Raw("SELECT * FROM config").RowsToMap(&ConfigMap, "configKey", "configValue")
	return
}
