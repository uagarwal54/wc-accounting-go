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
}
