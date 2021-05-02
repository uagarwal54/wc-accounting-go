package items

import (
	"fmt"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"wc-accounting-go/codebase/model"
)

type (
	addItemResp struct {
		Message    string
		ItemStatus map[string]string
		StatusCode int
	}
)

func addItem(w http.ResponseWriter, r *http.Request) {
	var bodyData []byte
	var err error
	var IList model.Items
	if bodyData, err = ioutil.ReadAll(r.Body); err != nil {
		w.Write([]byte("Error while reading the request"))
		return
	}
	// fmt.Println(string(bodyData))
	if err = json.Unmarshal(bodyData, &IList); err != nil {
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

func decideOp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Call the view function
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
