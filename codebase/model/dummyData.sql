use wecare;

insert into `user` (userId, userName, registrationDate, password, firstLogin) values("wc_admin_1","wc_admin","2021-04-27 06:44:00","admin123",1);

insert into `config` (configKey, configValue) values ("purchaseRoundOffDirection","ceil");
insert into `config` (configKey, configValue) values ("itemName","itemName");
insert into `config` (configKey, configValue) values ("itemCategory","itemCategory");
insert into `config` (configKey, configValue) values ("itemInsertionErrorMsgMarker","Something went wrong while storing the items");
insert into `config` (configKey, configValue) values ("duplicateEntryMsg","Duplicate entry");
insert into `config` (configKey, configValue) values ("itemAlreadyExistsMsg","The item already exists with us");
insert into `config` (configKey, configValue) values ("itemStorageFailourMsg","Item not stored");
insert into `config` (configKey, configValue) values ("successfullItemsInsertionMsg","All the above items are inserted into the database");
insert into `config` (configKey, configValue) values ("successfullItemInsertionMsg","Added the item to DB store");
insert into `config` (configKey, configValue) values ("proccessedFetchItemsRequestByName","Processed the request using the item names");
insert into `config` (configKey, configValue) values ("itemNotFoundErrorDuringFetchingMsg","Item Not Found. Some error occoured while fetching the items.");
insert into `config` (configKey, configValue) values ("itemFoundMsg","Item(s) Found");
insert into `config` (configKey, configValue) values ("itemCategoryNotFoundMsg","Item category not found");
insert into `config` (configKey, configValue) values ("itemIdNotFoundMsg","Item Id not found");
insert into `config` (configKey, configValue) values ("itemNameNotFoundMsg","Item Name not found");
insert into `config` (configKey, configValue) values ("itemDataNotFoundMsg","Item Data not found");
insert into `config` (configKey, configValue) values ("wrongInputToFetchItems","Item Not Found. Please check the info passed in the request.");
insert into `config` (configKey, configValue) values ("proccessedFetchItemsRequestByCategory","Processed the request using the item category");
insert into `config` (configKey, configValue) values ("wrongJSONKeyInRequest","Passed wrong json key in JSON request");
insert into `config` (configKey, configValue) values ("successfullItemUpdationMsg","Item is updated successfully");
insert into `config` (configKey, configValue) values ("itemUpdationErrorMsg","Error occoured while updating item: ");
insert into `config` (configKey, configValue) values ("failedItemUpdationDueToWrongInput","Item failed to be updated. Please check the data provided. Error: ");
insert into `config` (configKey, configValue) values ("itemSuccessfullDeletetionMsg","Item with the above id has been deleted");
insert into `config` (configKey, configValue) values ("itemDeletetionFailourMsg","Some problem occoured while deleting the item. Error: ");
insert into `config` (configKey, configValue) values ("filterAll","all");
insert into `config` (configKey, configValue) values ("iidSuffix","iid_");

