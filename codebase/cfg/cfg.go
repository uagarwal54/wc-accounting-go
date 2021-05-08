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
		ItemInsertionErrorMsgMarker           string     `json:"item_insertion_error_msg_marker"`
		DuplicateEntryMsg                     string     `json:"duplicate_entry_msg"`
		ItemAlreadyExistsMsg                  string     `json:"item_already_existsMsg"`
		ItemStorageFailourMsg                 string     `json:"item_storage_failour_msg"`
		SuccessfullItemsInsertionMsg          string     `json:"successfull_items_insertion_msg"`
		SuccessfullItemInsertionMsg           string     `json:"successfull_item_insertion_msg"`
		ProccessedFetchItemsRequestByName     string     `json:"proccessed_fetch_items_request_by_name"`
		ItemNotFoundErrorDuringFetchingMsg    string     `json:"item_not_found_error_during_fetching_msg"`
		ItemFoundMsg                          string     `json:"item_found_msg"`
		ItemCategoryNotFoundMsg               string     `json:"item_category_not_found_msg"`
		ItemIdNotFoundMsg                     string     `json:"item_id_not_found_msg"`
		ItemDataNotFoundMsg                   string     `json:"item_data_not_found_msg"`
		WrongInputToFetchItems                string     `json:"wrong_input_to_fetch_items"`
		ProccessedFetchItemsRequestByCategory string     `json:"proccessed_fetch_items_request_by_category"`
		WrongJSONKeyInRequest                 string     `json:"wrong_JSON_key_in_request"`
		SuccessfullItemUpdationMsg            string     `json:"successfull_item_updation_msg"`
		ItemUpdationErrorMsg                  string     `json:"item_updation_error_msg"`
		FailedItemUpdationDueToWrongInput     string     `json:"failed_item_updation_due_to_wrong_input"`
		ItemSuccessfullDeletetionMsg          string     `json:"item_successfull_deletetion_msg"`
		ItemDeletetionFailourMsg              string     `json:"item_deletetion_failour_msg"`
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
