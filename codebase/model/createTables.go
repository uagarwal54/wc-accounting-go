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

	// General Configs
	"roundOffDirection":              "ceil",
	"insertionErrorMsgMarker":        "Something went wrong while storing the object in data store",
	"duplicateEntryMsg":              "Duplicate entry",
	"alreadyExistsMsg":               "The object already exists with us",
	"storageFailourMsg":              "Object not stored",
	"proccessedFetchRequestByName":   "Processed the request using the object names",
	"notFoundErrorDuringFetchingMsg": "Object Not Found. Some error occoured while fetching the objects.",
	"objectFoundMsg":                 "Object(s) Found",
	"wrongInputToFetchObjects":       "Object Not Found. Please check the info passed in the request.",
	"wrongJSONKeyInRequest":          "Passed wrong json key in JSON request",
	"filterAll":                      "all",

	// CRUD Ops Configs
	"successfullInsertionMsg":       "All the above object(s) are inserted into the database",
	"successfullUpdationMsg":        "Object is updated successfully",
	"updationErrorMsg":              "Error occoured while updating object: ",
	"failedUpdationDueToWrongInput": "Object failed to be updated. Please check the data provided. Error: ",
	"successfullDeletetionMsg":      "Object with the above id has been deleted",
	"deletetionFailourMsg":          "Some problem occoured while deleting the object. Error: ",

	// Some Item specific configs
	"iidSuffix":                             "iid_",
	"itemName":                              "itemName",
	"itemCategory":                          "itemCategory",
	"itemCategoryNotFoundMsg":               "Item category not found",
	"itemIdNotFoundMsg":                     "Item Id not found",
	"itemNameNotFoundMsg":                   "Item Name not found",
	"itemDataNotFoundMsg":                   "Item Data not found",
	"proccessedFetchItemsRequestByCategory": "Processed the request using the item category",
}
