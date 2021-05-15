// This file reads the configs from the cfg.json file and feeds them into the struct
// There are some configs that are not stored in the DB and are only stored in the cfg file. Those and the keys of the configs
// stored in the DB can be accessed from here
package cfg

import (
	"encoding/json"
	"io/ioutil"
)

type (
	Configs struct {
		Dev                                   DevConfigs `json:"dev"`
		ItemName                              string     `json:"item_name"`
		ItemCategory                          string     `json:"item_category"`
		PurchaseRoundOffDirection             string     `json:"purchase_round_off_direction"`
		InsertionErrorMsgMarker               string     `json:"insertion_error_msg_marker"`
		DuplicateEntryMsg                     string     `json:"duplicate_entry_msg"`
		AlreadyExistsMsg                      string     `json:"already_exists_msg"`
		StorageFailourMsg                     string     `json:"storage_failour_msg"`
		SuccessfullInsertionMsg               string     `json:"successfull_insertion_msg"`
		ProccessedFetchRequestByName          string     `json:"proccessed_fetch_request_by_name"`
		NotFoundErrorDuringFetchingMsg        string     `json:"not_found_error_during_fetching_msg"`
		ObjectFoundMsg                        string     `json:"object_found_msg"`
		ItemCategoryNotFoundMsg               string     `json:"item_category_not_found_msg"`
		ItemIdNotFoundMsg                     string     `json:"item_id_not_found_msg"`
		ItemDataNotFoundMsg                   string     `json:"item_data_not_found_msg"`
		WrongInputToFetchObjects              string     `json:"wrong_input_to_fetch_objects"`
		ProccessedFetchItemsRequestByCategory string     `json:"proccessed_fetch_items_request_by_category"`
		WrongJSONKeyInRequest                 string     `json:"wrong_JSON_key_in_request"`
		SuccessfullUpdationMsg                string     `json:"successfull_updation_msg"`
		UpdationErrorMsg                      string     `json:"updation_error_msg"`
		FailedUpdationDueToWrongInput         string     `json:"failed_updation_due_to_wrong_input"`
		SuccessfullDeletetionMsg              string     `json:"successfull_deletetion_msg"`
		DeletetionFailourMsg                  string     `json:"deletetion_failour_msg"`
		ItemNameNotFoundMsg                   string     `json:"item_name_not_found_msg"`
		FilterAll                             string     `json:"filter_all"`
		IidSuffix                             string     `json:"iid_suffix"`
	}
	DevConfigs struct {
		Mysql       MysqlConfigs        `json:"mysql"`
		BackendUrls BackendUrlsFromJson `json:"backendUrls"`
	}
	MysqlConfigs struct {
		DBRootUser     string `json:"dbRootUser"`
		DBRootPassword string `json:"dbRootPassword"`
		DBName         string `json:"dbName"`
		DBHost         string `json:"dbHost"`
		Port           int    `json:"port"`
	}
	BackendUrlsFromJson struct {
		ItemMgmtBaseUrl      string `json:"item_mgmt_base_url"`
		SingleItemStoreUrl   string `json:"single_item_store_url"`
		MultipleItemStoreUrl string `json:"multiple_item_store_url"`
		CategoryMgmtBaseUrl  string `json:"category_mgmt_base_url"`
	}
)

var ConfigInst Configs

func GetConfigsFromConfigFile(configFilePath string) (err error) {
	var CfgData []byte
	if CfgData, err = ioutil.ReadFile(configFilePath); err != nil {
		return
	}
	err = json.Unmarshal(CfgData, &ConfigInst)
	return
}
