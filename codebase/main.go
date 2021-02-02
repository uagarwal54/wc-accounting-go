package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"wc-accounting-go/codebase/cfg"
	"wc-accounting-go/codebase/model"
)

const (
	configFilePath = "cfg.json"
	env            = "dev"
)

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}

type ErrorResponse struct {
	errMsg string
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ovnpwnv")
	var bodyData []byte
	var err error
	var userCredMap map[string]string
	if bodyData, err = ioutil.ReadAll(r.Body); err != nil {
		w.Write([]byte("Error while reading the request"))
		return
	}
	json.Unmarshal(bodyData, &userCredMap)
	var user model.User
	if user, err = model.ReadUserDataWithUsernamePassword(userCredMap["Username"], userCredMap["Password"]); err != nil {
		var errorData ErrorResponse
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		errorData.errMsg = "User Not Found"
		json.NewEncoder(w).Encode(errorData)
		fmt.Println(err)
		return
	}
	if (user != model.User{}) {
		fmt.Println("User Found")
		w.WriteHeader(http.StatusOK)
		if user.FirstLogin {
			w.Header().Set("FirstLogin", "True")
		} else {
			w.Header().Set("FirstLogin", "False")
		}
		w.Write([]byte(fmt.Sprintf("Welcome %s", user.UserName)))
	}
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

	router.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":9001", handlers.CORS(headers, methods, origins)(router)))
}
