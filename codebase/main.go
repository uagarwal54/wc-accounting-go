package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"wc-accounting-go/codebase/cfg"
	"wc-accounting-go/codebase/model"
	"wc-accounting-go/codebase/modules/login"
)

const (
	configFilePath = "cfg.json"
	env            = "dev"
)

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}

func handleRequests() {
	var configs cfg.Configs
	var err error
	if configs, err = cfg.GetConfigs(configFilePath); err != nil {
		fmt.Println(err)
		return
	}
	model.DbConnect(&configs, env)
	router := mux.NewRouter()

	// The headers, methods and domains from with/from which the requests can be accepted
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Autherization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/login", login.Login)
	log.Fatal(http.ListenAndServe(":9001", handlers.CORS(headers, methods, origins)(router)))
}
