package items

import(
	"fmt"

	"net/http"
	"io/ioutil"
	"encoding/json"

	"wc-accounting-go/codebase/model"
	"github.com/gorilla/mux"
)

type(
	Items struct{
		ItemList []model.Item `json: "items"`
	}
)

func addItem(w http.ResponseWriter, r *http.Request) {
	var bodyData []byte
	var err error
	var IList Items
	if bodyData, err = ioutil.ReadAll(r.Body); err != nil {
		w.Write([]byte("Error while reading the request"))
		return
	}
	fmt.Println(string(bodyData))
	err = json.Unmarshal(bodyData, &IList.ItemList)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(IList.ItemList)
	
}

func decideOp(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		// Call the view function
	}else if r.Method == http.MethodPost{
		addItem(w, r)
	}else if r.Method == http.MethodPut{
		// Call the item update method
	}else if r.Method == http.MethodDelete{
		// Call the method to delete the item
	}

}

func AddItemRoutes(router *mux.Router){
	router.HandleFunc("/itemMgmt", decideOp)
}