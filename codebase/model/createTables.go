package model

var createTableQueries = []string{
	`CREATE DATABASE IF NOT EXISTS wecare;`,
	`USE wecare;`,
	`create table IF NOT EXISTS user ( 
		srNum  int(10) PRIMARY KEY AUTO_INCREMENT,
		userId varchar(30),
		userName varchar(30),
		registrationDate datetime,
		password varchar(30),
		firstLogin int(1));`,

	`create table IF NOT EXISTS config ( 
		srNum  int(10) PRIMARY KEY AUTO_INCREMENT,
		configKey varchar(300),
		configValue  varchar(300));`,

	`create table IF NOT EXISTS item ( 
		srNum  int(10) PRIMARY KEY AUTO_INCREMENT,
		itemId varchar(30),
		itemName varchar(30) unique,
		itemCategory int(10));`,

	`create table IF NOT EXISTS itemCategory ( 
		srNum  int(10) PRIMARY KEY AUTO_INCREMENT,
		categoryId varchar(30),
		categoryName varchar(30));`,

	`create table IF NOT EXISTS purchase ( 
		srNum  int(10) PRIMARY KEY AUTO_INCREMENT,
		purchaseId varchar(30),
		itemId varchar(30),
		vendorName varchar(30),
		vendorAddress varchar(30),
		vendorPhNumber varchar(30),
		vendorEmail varchar(30),
		quantity int(10),
		metricUnit varchar(10),
		rate decimal(5,2),
		actualTotalCost decimal(15,5),
		roundOff int(11));`,
}

var configMap = map[string]string{
	"purchaseRoundOffDirection":             "ceil",
	"itemName":                              "itemName",
	"itemCategory":                          "itemCategory",
	"itemInsertionErrorMsgMarker":           "Something went wrong while storing the items",
	"duplicateEntryMsg":                     "Duplicate entry",
	"itemAlreadyExistsMsg":                  "The item already exists with us",
	"itemStorageFailourMsg":                 "Item not stored",
	"successfullItemsInsertionMsg":          "All the above items are inserted into the database",
	"successfullItemInsertionMsg":           "Added the item to DB store",
	"proccessedFetchItemsRequestByName":     "Processed the request using the item names",
	"itemNotFoundErrorDuringFetchingMsg":    "Item Not Found. Some error occoured while fetching the items.",
	"itemFoundMsg":                          "Item(s) Found",
	"itemCategoryNotFoundMsg":               "Item category not found",
	"itemIdNotFoundMsg":                     "Item Id not found",
	"itemNameNotFoundMsg":                   "Item Name not found",
	"itemDataNotFoundMsg":                   "Item Data not found",
	"wrongInputToFetchItems":                "Item Not Found. Please check the info passed in the request.",
	"proccessedFetchItemsRequestByCategory": "Processed the request using the item category",
	"wrongJSONKeyInRequest":                 "Passed wrong json key in JSON request",
	"successfullItemUpdationMsg":            "Item is updated successfully",
	"itemUpdationErrorMsg":                  "Error occoured while updating item: ",
	"failedItemUpdationDueToWrongInput":     "Item failed to be updated. Please check the data provided. Error: ",
	"itemSuccessfullDeletetionMsg":          "Item with the above id has been deleted",
	"itemDeletetionFailourMsg":              "Some problem occoured while deleting the item. Error: ",
	"filterAll":                             "all",
	"iidSuffix":                             "iid_",
}
