package items

import (
	"fmt"

	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"wc-accounting-go/codebase/cfg"
	"wc-accounting-go/codebase/model"
)

type (
	// categoryResponseToBeSent is the struct used in all the add item requests' response
	categoryResponseToBeSent struct {
		Message        string
		CategoryStatus []categoryResponse
		StatusCode     int
	}

	// itemResponse is the struct that has the fields that are returned to the user
	//so that only the desired fields are exposed to the user
	categoryResponse struct {
		CategoryName string `json: "categoryName,omitempty"`
		CategoryId   string `json: "categoryId,omitempty"`
		Message      string `json: "message,omitempty"`
	}
	categoryForUpdateInput struct {
		CategoryName    string `json: "categoryName,omitempty"`
		NewCategoryName string `json: "newCategoryName,omitempty"`
	}
)

func storeCategory(inputItemCategory *model.Itemcategory) (categoryId string, err error) {
	var numberOfRows int
	numberOfRows, err = model.CountCategoryRows()
	numberOfRows = numberOfRows + 1
	categoryId = model.ConfigMap[cfg.ConfigInst.CidPrefix].(string) + strconv.Itoa(numberOfRows)
	fmt.Println(categoryId)
	inputItemCategory.CategoryId = categoryId
	if err = inputItemCategory.InsertRecordIntoItemCategory(); err != nil {
		return
	}
	return
}

func storeCategoriesInDB(w http.ResponseWriter, r *http.Request) {
	var CList model.Categories
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &CList); err != nil {
		fmt.Println(err)
	}
	var addCategoryRespInst categoryResponseToBeSent
	for _, categoryInst := range CList.CategoryList {
		var categoryRespInst categoryResponse
		categoryRespInst.CategoryName = categoryInst.CategoryName
		if categoryRespInst.CategoryId, err = storeCategory(&categoryInst); err != nil {
			fmt.Println("Error while inserting the category data: ", categoryInst)
			fmt.Println("Error: ", err)
			addCategoryRespInst.Message = model.ConfigMap[cfg.ConfigInst.InsertionErrorMsgMarker].(string)
			addCategoryRespInst.StatusCode = http.StatusBadRequest
			if strings.Contains(err.Error(), model.ConfigMap[cfg.ConfigInst.DuplicateEntryMsg].(string)) {
				categoryRespInst.Message = model.ConfigMap[cfg.ConfigInst.AlreadyExistsMsg].(string)
			} else {
				categoryRespInst.Message = err.Error()
			}
		} else {
			categoryRespInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullInsertionMsg].(string)
			addCategoryRespInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullInsertionMsg].(string)
			addCategoryRespInst.StatusCode = http.StatusOK
		}
		addCategoryRespInst.CategoryStatus = append(addCategoryRespInst.CategoryStatus, categoryRespInst)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(addCategoryRespInst)
}

func fetchCategories(w http.ResponseWriter, r *http.Request) {
	var fetchRequestData map[string]interface{}
	json.Unmarshal(readRequestData(w, r), &fetchRequestData)
	var fetchCategoryResponseInst categoryResponseToBeSent
	var categoryResponseList []categoryResponse

	if _, found := fetchRequestData[cfg.ConfigInst.CategoryName]; found {
		fetchCategoryResponseInst.Message = model.ConfigMap[cfg.ConfigInst.ProccessedFetchRequestByName].(string)
		for _, categoryName := range fetchRequestData[cfg.ConfigInst.CategoryName].([]interface{}) {
			if strings.ToLower(categoryName.(string)) == model.ConfigMap[cfg.ConfigInst.FilterAll].(string) {
				categoryListModel := &model.Categories{}
				message := model.ConfigMap[cfg.ConfigInst.ObjectFoundMsg].(string)
				if err := categoryListModel.ReadAllCategoryData(); err != nil {
					fmt.Println(err)
					message = model.ConfigMap[cfg.ConfigInst.NotFoundErrorDuringFetchingMsg].(string)
				}
				for _, categoryInst := range categoryListModel.CategoryList {
					if categoryInst.CategoryName == "" {
						message = model.ConfigMap[cfg.ConfigInst.CategoryNameNotFoundMsg].(string)
					} else if categoryInst.CategoryId == "" {
						message = model.ConfigMap[cfg.ConfigInst.CategoryIdNotFoundMsg].(string)
					}
					// Not including the item name in this cond because it is passed from the user so mostly it would be there
					if categoryInst.CategoryId == "" && categoryInst.CategoryName == "" {
						message = model.ConfigMap[cfg.ConfigInst.CategoryDataNotFoundMsg].(string)
					}
					populateResponseCategoryList(&categoryResponseList, &categoryInst, message)
				}
			} else {
				message := model.ConfigMap[cfg.ConfigInst.ObjectFoundMsg].(string)
				categoryInst := &model.Itemcategory{}
				categoryInst.CategoryName = categoryName.(string)
				if err := categoryInst.ReadCategoryByName(); err != nil {
					fmt.Println(err)
					message = model.ConfigMap[cfg.ConfigInst.WrongInputToFetchObjects].(string)
				}
				populateResponseCategoryList(&categoryResponseList, categoryInst, message)
			}
		}
	} else {
		fetchCategoryResponseInst.Message = model.ConfigMap[cfg.ConfigInst.WrongJSONKeyInRequest].(string)
		fetchCategoryResponseInst.StatusCode = http.StatusBadRequest
	}
	fetchCategoryResponseInst.StatusCode = http.StatusOK
	fetchCategoryResponseInst.CategoryStatus = categoryResponseList
	json.NewEncoder(w).Encode(fetchCategoryResponseInst)
}

func updateCategories(w http.ResponseWriter, r *http.Request) {
	var CList []categoryForUpdateInput
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &CList); err != nil {
		fmt.Println(err)
	}
	var addCategoryRespInst categoryResponseToBeSent
	for _, category := range CList {
		categoryPtr := &model.Itemcategory{}
		categoryPtr.CategoryName = category.CategoryName
		var categoryResponseInst categoryResponse
		categoryResponseInst.CategoryName = category.CategoryName
		categoryResponseInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullUpdationMsg].(string)

		if err = categoryPtr.UpdateRecord(category.NewCategoryName); err != nil {
			fmt.Println("Error occoured while updating item: ", category)
			fmt.Println(err)
			categoryResponseInst.Message = model.ConfigMap[cfg.ConfigInst.FailedUpdationDueToWrongInput].(string) + err.Error()
		}
		categoryResponseInst.CategoryName = category.NewCategoryName
		addCategoryRespInst.CategoryStatus = append(addCategoryRespInst.CategoryStatus, categoryResponseInst)
	}
	addCategoryRespInst.StatusCode = http.StatusOK
	json.NewEncoder(w).Encode(addCategoryRespInst)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	var inputItemCategory model.Itemcategory
	var err error
	if err = json.Unmarshal(readRequestData(w, r), &inputItemCategory); err != nil {
		fmt.Println(err)
	}
	var categoryResponseInst categoryResponse
	categoryResponseInst.CategoryName = inputItemCategory.CategoryName
	categoryResponseInst.Message = model.ConfigMap[cfg.ConfigInst.SuccessfullDeletetionMsg].(string)
	if err = inputItemCategory.DeleteRecord(); err != nil {
		fmt.Println(err)
		categoryResponseInst.Message = model.ConfigMap[cfg.ConfigInst.DeletetionFailourMsg].(string) + err.Error()
	}
	json.NewEncoder(w).Encode(categoryResponseInst)
}

func decideOpForCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fetchCategories(w, r)
	} else if r.Method == http.MethodPatch {
		updateCategories(w, r)
	} else if r.Method == http.MethodDelete {
		deleteCategory(w, r)
	} else if r.Method == http.MethodPut {
		// store multiple items
		storeCategoriesInDB(w, r)
	}

}

func populateResponseCategoryList(categoryResponseList *[]categoryResponse, CategoryInst *model.Itemcategory, message string) {
	var tempCategoryResponseInst categoryResponse
	tempCategoryResponseInst.CategoryName = CategoryInst.CategoryName
	tempCategoryResponseInst.CategoryId = CategoryInst.CategoryId
	tempCategoryResponseInst.Message = message
	*categoryResponseList = append(*categoryResponseList, tempCategoryResponseInst)
}

func AddCategoryRoutes(router *mux.Router) {
	router.HandleFunc(cfg.ConfigInst.Dev.BackendUrls.CategoryMgmtBaseUrl, decideOpForCategories)
}
