package items

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"wc-accounting-go/codebase/cfg"
	"wc-accounting-go/codebase/model"
)

type (
	// itemResponseToBeSent is the struct used in all the add item requests' response
	itemResponseToBeSent struct {
		Message    string
		ItemStatus []itemResponse
		StatusCode int
	}

	// itemResponse is the struct that has the fields that are returned to the user
	//so that only the desired fields are exposed to the user
	itemResponse struct {
		ItemName     string `json: "itemname,omitempty"`
		ItemCategory int    `json: "itemcategory,omitempty"`
		ItemId       string `json: "itemid,omitempty"`
		Message      string `json: "message,omitempty"`
	}
)

// This function uses multi insert of beego to insert multiple records inside the item table.
// We wrote this function and then decided not to use it for the time being as at that time we felt that having single
// inserts into the table would suffice
func addMultipleItemsInItemAtOnce(w http.ResponseWriter, r *http.Request) {
	var err error
	var IList model.Items
	if err = json.Unmarshal(readRequestData(w, r), &IList); err != nil {
		fmt.Println(err)
	}
	var numberOfRows int
	var addItemRespInst itemResponseToBeSent
	var finalItemList model.Items

	numberOfRows, err = model.CountItemRows()
	for _, item := range IList.ItemList {
		var itemRespInst itemResponse
		numberOfRows = numberOfRows + 1
		item.ItemId = "iid_" + strconv.Itoa(numberOfRows)
		finalItemList.ItemList = append(finalItemList.ItemList, item)
		itemRespInst.ItemName = item.ItemName
		itemRespInst.ItemCategory = item.ItemCategory
		itemRespInst.ItemId = item.ItemId
		itemRespInst.Message = "Added the item to DB store"
		addItemRespInst.ItemStatus = append(addItemRespInst.ItemStatus, itemRespInst)
	}
	if err = finalItemList.InsertRecordsIntoItem(); err != nil {
		fmt.Println("Error while inserting the item data: ", finalItemList.ItemList)
		fmt.Println("Error: ", err)
		addItemRespInst.Message = "Something went wrong while storing the items"
		addItemRespInst.StatusCode = http.StatusBadRequest
		for _, itemRsponseObj := range addItemRespInst.ItemStatus {
			if strings.Contains(err.Error(), itemRsponseObj.ItemName) {
				if strings.Contains(err.Error(), "Duplicate entry") {
					itemRsponseObj.Message = "The item already exists with us"
				} else {
					itemRsponseObj.Message = err.Error()
				}
			} else {
				itemRsponseObj.Message = "Item not stored"
			}
		}
	} else {
		addItemRespInst.Message = "All the above items are inserted into the database"
		addItemRespInst.StatusCode = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addItemRespInst)
}

func storeItem(inputItem *model.Item) (itemId string, err error) {
	var numberOfRows int
	numberOfRows, err = model.CountItemRows()
	numberOfRows = numberOfRows + 1
	itemId = "iid_" + strconv.Itoa(numberOfRows)
	inputItem.ItemId = itemId
	if err = inputItem.InsertRecordIntoItem(); err != nil {
		return
	}
	return
}

func storeMultipleItemsWithSingleInsertForEachItem(w http.ResponseWriter, r *http.Request) {
	var IList model.Items
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &IList); err != nil {
		fmt.Println(err)
	}
	var addItemRespInst itemResponseToBeSent
	for _, item := range IList.ItemList {
		var itemRespInst itemResponse
		itemRespInst.ItemName = item.ItemName
		itemRespInst.ItemCategory = item.ItemCategory
		if itemRespInst.ItemId, err = storeItem(&item); err != nil {
			fmt.Println("Error while inserting the item data: ", item)
			fmt.Println("Error: ", err)
			addItemRespInst.Message = "Something went wrong while storing the items"
			addItemRespInst.StatusCode = http.StatusBadRequest
			if strings.Contains(err.Error(), "Duplicate entry") {
				itemRespInst.Message = "The item already exists with us"
			} else {
				itemRespInst.Message = err.Error()
			}
		} else {
			itemRespInst.Message = "Added the item to DB store"
			addItemRespInst.Message = "The above item is stored with us now"
			addItemRespInst.StatusCode = http.StatusOK
		}
		addItemRespInst.ItemStatus = append(addItemRespInst.ItemStatus, itemRespInst)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addItemRespInst)
}

func storeSingleItem(w http.ResponseWriter, r *http.Request) {
	var inputItem model.Item
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &inputItem); err != nil {
		fmt.Println(err)
	}
	var addItemRespInst itemResponseToBeSent
	var itemRespInst itemResponse
	itemRespInst.ItemName = inputItem.ItemName
	itemRespInst.ItemCategory = inputItem.ItemCategory
	addItemRespInst.ItemStatus = append(addItemRespInst.ItemStatus, itemRespInst)
	if addItemRespInst.ItemStatus[0].ItemId, err = storeItem(&inputItem); err != nil {
		fmt.Println("Error while inserting the item data: ", inputItem)
		fmt.Println("Error: ", err)
		addItemRespInst.Message = "Something went wrong while storing the items"
		addItemRespInst.StatusCode = http.StatusBadRequest
		if strings.Contains(err.Error(), "Duplicate entry") {
			addItemRespInst.ItemStatus[0].Message = "The item already exists with us"
		} else {
			addItemRespInst.ItemStatus[0].Message = err.Error()
		}
	} else {
		addItemRespInst.ItemStatus[0].Message = "Added the item to DB store"
		addItemRespInst.Message = "The above item is stored with us now"
		addItemRespInst.StatusCode = http.StatusOK
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addItemRespInst)
}

func fetchItems(w http.ResponseWriter, r *http.Request) {
	var fetchRequestData map[string]interface{}
	json.Unmarshal(readRequestData(w, r), &fetchRequestData)
	var fetchItemResponseInst itemResponseToBeSent
	var itemResponseList []itemResponse

	if _, found := fetchRequestData[cfg.ConfigInst.Dev.ItemName]; found {
		fetchItemResponseInst.Message = "Processed the request using the item names"
		for _, itemName := range fetchRequestData[cfg.ConfigInst.Dev.ItemName].([]interface{}) {
			if strings.ToLower(itemName.(string)) == "all" {
				itemListModel := &model.Items{}
				message := "Items Found"
				if err := itemListModel.ReadAllItemData(); err != nil {
					fmt.Println(err)
					message = "Item Not Found. Some error occoured while fetching the items."
				}
				for _, itemInst := range itemListModel.ItemList {
					if itemInst.ItemName == "" {
						message = "Item Name not found"
					} else if itemInst.ItemCategory == 0 {
						message = "Item category not found"
					} else if itemInst.ItemId == "" {
						message = "Item Id not found"
					}
					// Not including the item name in this cond because it is passed from the user so mostly it would be there
					if itemInst.ItemCategory == 0 && itemInst.ItemId == "" {
						message = "Item Data not found"
					}
					populateResponseItemList(&itemResponseList, &itemInst, message)
				}
				break // Come out of the outer most for loop
			} else {
				message := "Item Found"
				itemInst := &model.Item{}
				itemInst.ItemName = itemName.(string)
				if err := itemInst.ReadItemByName(); err != nil {
					fmt.Println(err)
					message = "Item Not Found. Please check the info passed in the request."
				}
				populateResponseItemList(&itemResponseList, itemInst, message)
			}
		}
	} else if _, found := fetchRequestData[cfg.ConfigInst.Dev.ItemCategory]; found {
		fetchItemResponseInst.Message = "Processed the request using the item category"
		desiredCategory := fetchRequestData[cfg.ConfigInst.Dev.ItemCategory]
		fetchedItemList := &model.Items{}
		if err := fetchedItemList.ReadAllItemsInACategory(int(desiredCategory.(float64))); err != nil {
			fmt.Println(err)
			fetchItemResponseInst.Message = "Something went wrong: " + err.Error()
			fetchItemResponseInst.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(fetchItemResponseInst)
			return
		}
		for _, fetchedItem := range fetchedItemList.ItemList {
			populateResponseItemList(&itemResponseList, &fetchedItem, "Item Found")
		}
	} else {
		fetchItemResponseInst.Message = "Passed wrong json key in JSON request"
		fetchItemResponseInst.StatusCode = http.StatusBadRequest
	}
	fetchItemResponseInst.StatusCode = http.StatusOK
	fetchItemResponseInst.ItemStatus = itemResponseList
	json.NewEncoder(w).Encode(fetchItemResponseInst)
}

func updateItems(w http.ResponseWriter, r *http.Request) {
	var IList model.Items
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &IList); err != nil {
		fmt.Println(err)
	}
	var addItemRespInst itemResponseToBeSent
	for _, item := range IList.ItemList {
		itemPtr := &item
		var itemResponseInst itemResponse
		itemResponseInst.ItemName = item.ItemName
		itemResponseInst.ItemId = item.ItemId
		itemResponseInst.ItemCategory = item.ItemCategory
		itemResponseInst.Message = "Item is updated successfully"

		if err = itemPtr.UpdateRecord(); err != nil {
			fmt.Println("Error occoured while updating item: ", item)
			fmt.Println(err)
			itemResponseInst.Message = "Item failed to be updated. Please check the data provided. Error: " + err.Error()
		}
		addItemRespInst.ItemStatus = append(addItemRespInst.ItemStatus, itemResponseInst)
	}
	addItemRespInst.StatusCode = http.StatusOK
	json.NewEncoder(w).Encode(addItemRespInst)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	var inputItem model.Item
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &inputItem); err != nil {
		fmt.Println(err)
	}
	var itemResponseInst itemResponse
	itemResponseInst.ItemId = inputItem.ItemId
	itemResponseInst.Message = "Item with the above id has been deleted"
	if err = inputItem.DeleteRecords(); err != nil {
		fmt.Println(err)
		itemResponseInst.Message = "Some problem occoured while deleting the item. Error: " + err.Error()
	}
	json.NewEncoder(w).Encode(itemResponseInst)
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
	} else if r.Method == http.MethodPost && r.URL.Path == "/itemMgmt" {
		// Unused function
		addMultipleItemsInItemAtOnce(w, r)
	} else if r.Method == http.MethodPut {
		updateItems(w, r)
	} else if r.Method == http.MethodDelete {
		deleteItem(w, r)
	} else if r.Method == http.MethodPost && r.URL.Path == "/itemMgmt/storeItem" {
		// store single item
		storeSingleItem(w, r)
	} else if r.Method == http.MethodPost && r.URL.Path == "/itemMgmt/storeItems" {
		// store multiple items
		storeMultipleItemsWithSingleInsertForEachItem(w, r)
	}

}

func populateResponseItemList(itemResponseList *[]itemResponse, ItemInst *model.Item, message string) {
	var tempItemResponseInst itemResponse
	tempItemResponseInst.ItemName = ItemInst.ItemName
	tempItemResponseInst.ItemCategory = ItemInst.ItemCategory
	tempItemResponseInst.ItemId = ItemInst.ItemId
	tempItemResponseInst.Message = message
	*itemResponseList = append(*itemResponseList, tempItemResponseInst)
}

func AddItemRoutes(router *mux.Router) {
	router.HandleFunc("/itemMgmt", decideOp)
	router.HandleFunc("/itemMgmt/storeItem", decideOp)
	router.HandleFunc("/itemMgmt/storeItems", decideOp)
}
