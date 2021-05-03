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
	// addItemResp is the struct used in all the add item requests' response
	addItemResp struct {
		Message    string
		ItemStatus map[string]string
		StatusCode int
	}
	// fetchItemResponse is the struct used in all the fetch item requests' response
	fetchItemResponse struct {
		Message    string
		Items      []itemResponse
		StatusCode int
	}
	// itemResponse is the struct that has the fields that are returned to the user
	//so that only the desired fields are exposed to the user
	itemResponse struct {
		ItemName     string
		ItemCategory int
		ItemId       string
		Message      string
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
	var itemList []itemResponse

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
					populateResponseItemList(&itemList, &itemInst, message)
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
				populateResponseItemList(&itemList, itemInst, message)
			}
		}
	} else if _, found := fetchRequestData[cfg.ConfigInst.Dev.ItemCategory]; found {
		fetchItemResponseInst.Message = "Processed the request using the item category"
	} else {
		fetchItemResponseInst.Message = "Passed wrong json key in JSON request"
		fetchItemResponseInst.StatusCode = http.StatusBadRequest
	}
	fetchItemResponseInst.StatusCode = http.StatusOK
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

func populateResponseItemList(itemList *[]itemResponse, ItemInst *model.Item, message string) {
	var tempItemResponseInst itemResponse
	tempItemResponseInst.ItemName = ItemInst.ItemName
	tempItemResponseInst.ItemCategory = ItemInst.ItemCategory
	tempItemResponseInst.ItemId = ItemInst.ItemId
	tempItemResponseInst.Message = message
	*itemList = append(*itemList, tempItemResponseInst)
}

func AddItemRoutes(router *mux.Router) {
	router.HandleFunc("/itemMgmt", decideOp)
}
