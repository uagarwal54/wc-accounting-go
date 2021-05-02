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
		Mysql MysqlConfigs `json:"mysql"`
	}
	MysqlConfigs struct {
		DBRootUser     string `json:"dbRootUser"`
		DBRootPassword string `json:"dbRootPassword"`
		DBName         string `json:"dbName"`
		DBHost         string `json:"dbHost"`
		Port           int    `json:"port"`
	}
)

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
