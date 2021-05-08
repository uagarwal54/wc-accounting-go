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

func storeItem(inputItem *model.Item) (itemId string, err error) {
	var numberOfRows int
	numberOfRows, err = model.CountItemRows()
	numberOfRows = numberOfRows + 1
	itemId = model.ConfigMap[cfg.ConfigInst.IidSuffix].(string) + strconv.Itoa(numberOfRows)
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
			addItemRespInst.Message = model.ConfigMap[cfg.ConfigInst.ItemInsertionErrorMsgMarker].(string)
			addItemRespInst.StatusCode = http.StatusBadRequest
			if strings.Contains(err.Error(), model.ConfigMap[cfg.ConfigInst.DuplicateEntryMsg].(string)) {
				itemRespInst.Message = model.ConfigMap[cfg.ConfigInst.ItemAlreadyExistsMsg].(string)
			} else {
				itemRespInst.Message = err.Error()
			}
		} else {
			itemRespInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullItemInsertionMsg].(string)
			addItemRespInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullItemsInsertionMsg].(string)
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
		addItemRespInst.Message = model.ConfigMap[cfg.ConfigInst.ItemInsertionErrorMsgMarker].(string)
		addItemRespInst.StatusCode = http.StatusBadRequest
		if strings.Contains(err.Error(), model.ConfigMap[cfg.ConfigInst.DuplicateEntryMsg].(string)) {
			addItemRespInst.ItemStatus[0].Message = model.ConfigMap[cfg.ConfigInst.ItemAlreadyExistsMsg].(string)
		} else {
			addItemRespInst.ItemStatus[0].Message = err.Error()
		}
	} else {
		addItemRespInst.ItemStatus[0].Message = model.ConfigMap[cfg.ConfigInst.SuccessfullItemInsertionMsg].(string)
		addItemRespInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullItemInsertionMsg].(string)
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

	if _, found := fetchRequestData[cfg.ConfigInst.ItemName]; found {
		fetchItemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.ProccessedFetchItemsRequestByName].(string)
		for _, itemName := range fetchRequestData[cfg.ConfigInst.ItemName].([]interface{}) {
			if strings.ToLower(itemName.(string)) == model.ConfigMap[cfg.ConfigInst.FilterAll].(string) {
				itemListModel := &model.Items{}
				message := model.ConfigMap[cfg.ConfigInst.ItemFoundMsg].(string)
				if err := itemListModel.ReadAllItemData(); err != nil {
					fmt.Println(err)
					message = model.ConfigMap[cfg.ConfigInst.ItemNotFoundErrorDuringFetchingMsg].(string)
				}
				for _, itemInst := range itemListModel.ItemList {
					if itemInst.ItemName == "" {
						message = model.ConfigMap[cfg.ConfigInst.ItemNameNotFoundMsg].(string)
					} else if itemInst.ItemCategory == 0 {
						message = model.ConfigMap[cfg.ConfigInst.ItemCategoryNotFoundMsg].(string)
					} else if itemInst.ItemId == "" {
						message = model.ConfigMap[cfg.ConfigInst.ItemIdNotFoundMsg].(string)
					}
					// Not including the item name in this cond because it is passed from the user so mostly it would be there
					if itemInst.ItemCategory == 0 && itemInst.ItemId == "" {
						message = model.ConfigMap[cfg.ConfigInst.ItemDataNotFoundMsg].(string)
					}
					populateResponseItemList(&itemResponseList, &itemInst, message)
				}
				break // Come out of the outer most for loop
			} else {
				message := model.ConfigMap[cfg.ConfigInst.ItemFoundMsg].(string)
				itemInst := &model.Item{}
				itemInst.ItemName = itemName.(string)
				if err := itemInst.ReadItemByName(); err != nil {
					fmt.Println(err)
					message = model.ConfigMap[cfg.ConfigInst.WrongInputToFetchItems].(string)
				}
				populateResponseItemList(&itemResponseList, itemInst, message)
			}
		}
	} else if _, found := fetchRequestData[cfg.ConfigInst.ItemCategory]; found {
		fetchItemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.ProccessedFetchItemsRequestByCategory].(string)
		desiredCategory := fetchRequestData[cfg.ConfigInst.ItemCategory]
		fetchedItemList := &model.Items{}
		if err := fetchedItemList.ReadAllItemsInACategory(int(desiredCategory.(float64))); err != nil {
			fmt.Println(err)
			fetchItemResponseInst.Message = "Something went wrong: " + err.Error()
			fetchItemResponseInst.StatusCode = http.StatusInternalServerError
			json.NewEncoder(w).Encode(fetchItemResponseInst)
			return
		}
		for _, fetchedItem := range fetchedItemList.ItemList {
			populateResponseItemList(&itemResponseList, &fetchedItem, model.ConfigMap[cfg.ConfigInst.ItemFoundMsg].(string))
		}
	} else {
		fetchItemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.WrongJSONKeyInRequest].(string)
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
		itemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullItemUpdationMsg].(string)

		if err = itemPtr.UpdateRecord(); err != nil {
			fmt.Println("Error occoured while updating item: ", item)
			fmt.Println(err)
			itemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.FailedItemUpdationDueToWrongInput].(string) + err.Error()
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
	itemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.ItemSuccessfullDeletetionMsg].(string)
	if err = inputItem.DeleteRecords(); err != nil {
		fmt.Println(err)
		itemResponseInst.Message = model.ConfigMap[cfg.ConfigInst.ItemDeletetionFailourMsg].(string) + err.Error()
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
	} else if r.Method == http.MethodPatch {
		updateItems(w, r)
	} else if r.Method == http.MethodDelete {
		deleteItem(w, r)
	} else if r.Method == http.MethodPut && r.URL.Path == cfg.ConfigInst.Dev.BackendUrls.SingleItemStoreUrl {
		// store single item
		storeSingleItem(w, r)
	} else if r.Method == http.MethodPut && r.URL.Path == cfg.ConfigInst.Dev.BackendUrls.MultipleItemStoreUrl {
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
	router.HandleFunc(cfg.ConfigInst.Dev.BackendUrls.ItemMgmtBaseUrl, decideOp)
	router.HandleFunc(cfg.ConfigInst.Dev.BackendUrls.SingleItemStoreUrl, decideOp)
	router.HandleFunc(cfg.ConfigInst.Dev.BackendUrls.MultipleItemStoreUrl, decideOp)
}
