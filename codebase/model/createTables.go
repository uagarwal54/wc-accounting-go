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
		configKey varchar(30),
		configValue  varchar(30));`,

	`create table IF NOT EXISTS item ( 
		srNum  int(10) PRIMARY KEY AUTO_INCREMENT,
		itemId varchar(30),
		itemName varchar(30),
		itemCategory varchar(30));`,
	
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
