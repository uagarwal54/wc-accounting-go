package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"wc-accounting-go/codebase/cfg"
	"wc-accounting-go/codebase/model"
	"wc-accounting-go/codebase/modules/items"
	"wc-accounting-go/codebase/modules/login"
)

const (
	configFilePath = "cfg.json"
	env            = "dev"
)

func main() {
	fmt.Println("My Simple Server")
	makeDbConnection()
	model.PopulateConfigMap()
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter()
	// These are the headers, methods and domains from with/from which the requests can be accepted
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Autherization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	login.AddLoginRoute(router)
	items.AddItemRoutes(router)
	items.AddCategoryRoutes(router)

	log.Fatal(http.ListenAndServe(":9001", handlers.CORS(headers, methods, origins)(router)))
}

func makeDbConnection() {
	var err error
	if err = cfg.GetConfigsFromConfigFile(configFilePath); err != nil {
		fmt.Println(err)
		return
	}
	model.DbConnect(env)
}
