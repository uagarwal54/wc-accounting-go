package items

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"wc-accounting-go/codebase/cfg"
	"wc-accounting-go/codebase/model"
)

type (
	addItemResp struct {
		Message    string
		ItemStatus map[string]string
		StatusCode int
	}
	fetchItemResponse struct {
		Message string
		Items   []*model.Item
	}
)

func addItem(w http.ResponseWriter, r *http.Request) {
	var err error
	var IList model.Items
	if err = json.Unmarshal(readRequestData(w, r), &IList); err != nil {
		fmt.Println(err)
	}
	var numberOfRows int
	var addItemRespInst addItemResp
	var finalItemList model.Items
	addItemRespInst.ItemStatus = make(map[string]string)
	numberOfRows, err = model.CountItemRows()
	for _, item := range IList.ItemList {
		numberOfRows = numberOfRows + 1
		item.ItemId = "iid_" + strconv.Itoa(numberOfRows)
		finalItemList.ItemList = append(finalItemList.ItemList, item)
		addItemRespInst.ItemStatus[item.ItemName] = "Added the item to DB store"
	}
	if err = finalItemList.InsertIntoItem(); err != nil {
		fmt.Println("Error while inserting the item data: ", IList.ItemList)
		fmt.Println("Error: ", err)
		addItemRespInst.Message = err.Error()
		addItemRespInst.StatusCode = http.StatusInternalServerError
	} else {
		addItemRespInst.Message = "All the above items are inserted into the database"
		addItemRespInst.StatusCode = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addItemRespInst)
}

func fetchItems(w http.ResponseWriter, r *http.Request) {
	var fetchRequestData map[string]interface{}
	json.Unmarshal(readRequestData(w, r), &fetchRequestData)
	var fetchItemResponseInst fetchItemResponse
	var itemList []*model.Item
	
	if _, found := fetchRequestData[cfg.ConfigInst.Dev.ItemName]; found {
		fetchItemResponseInst.Message = "Processed the request using the item names"
		for _, itemName := range  fetchRequestData[cfg.ConfigInst.Dev.ItemName].([]interface{}){
			ItemInst := &model.Item{}
			ItemInst.ItemName = itemName.(string)
			if err := ItemInst.ReadItemByName(); err != nil {
				fmt.Println(err)
			}
			itemList = append(itemList, ItemInst)
		}
	} else if _, found := fetchRequestData[cfg.ConfigInst.Dev.ItemCategory]; found {
		fetchItemResponseInst.Message = "Processed the request using the item category"
	} else {
		fetchItemResponseInst.Message = "Passed wrong json key in JSON request"
	}
	//remove
	fmt.Println(itemList)
	fetchItemResponseInst.Items = itemList
	json.NewEncoder(w).Encode(fetchItemResponseInst)
}

func readRequestData(w http.ResponseWriter, r *http.Request) (bodyData []byte) {
	var err error
	if bodyData, err = ioutil.ReadAll(r.Body); err != nil {
		w.Write([]byte("Error while reading the request"))
		return
	}
	return (bodyData)
}

func decideOp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fetchItems(w, r)
	} else if r.Method == http.MethodPost {
		addItem(w, r)
	} else if r.Method == http.MethodPut {
		// Call the item update method
	} else if r.Method == http.MethodDelete {
		// Call the method to delete the item
	}

}

func AddItemRoutes(router *mux.Router) {
	router.HandleFunc("/itemMgmt", decideOp)
}
