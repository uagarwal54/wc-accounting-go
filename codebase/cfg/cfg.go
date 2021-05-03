package cfg

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Configs struct {
		Dev DevConfigs `json:"dev"`
	}
	DevConfigs struct {
		Mysql        MysqlConfigs `json:"mysql"`
		ItemName     string       `json:"item_name"`
		ItemCategory string       `json:"item_category"`
	}
	MysqlConfigs struct {
		DBRootUser     string `json:"dbRootUser"`
		DBRootPassword string `json:"dbRootPassword"`
		DBName         string `json:"dbName"`
		DBHost         string `json:"dbHost"`
		Port           int    `json:"port"`
	}
)

var ConfigInst Configs

func GetConfigs(configFilePath string) (cfg Configs, err error) {
	var CfgData []byte
	if CfgData, err = ioutil.ReadFile(configFilePath); err != nil {
		return
	}
	if err = json.Unmarshal(CfgData, &cfg); err != nil {
		return
	}
	return
}
