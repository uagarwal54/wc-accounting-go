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

type (
		LoginResp struct{
			Message string
			StatusCode int
			FirstLogin string 
		}
)

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login Attempt")
	var bodyData []byte
	var err error
	var userCredMap map[string]string
	if bodyData, err = ioutil.ReadAll(r.Body); err != nil {
		w.Write([]byte("Error while reading the request"))
		return
	}
	json.Unmarshal(bodyData, &userCredMap)
	var user model.User
	w.Header().Set("Content-Type", "application/json")
	var loginResp LoginResp
	if user, err = model.ReadUserDataWithUsernamePassword(userCredMap["Username"], userCredMap["Password"]); err != nil {
		loginResp.StatusCode = http.StatusNotFound
		loginResp.Message = "User not found"
		loginResp.FirstLogin = ""
		fmt.Println(err)
	}
	if (user != model.User{}) {
		fmt.Println("User Found")
		loginResp.StatusCode = http.StatusOK
		loginResp.Message = fmt.Sprintf("Welcome %s",user.UserName)
		if user.FirstLogin == 1{
			loginResp.FirstLogin = "True"
		} else {
			loginResp.FirstLogin = "False"
		}
	}
	json.NewEncoder(w).Encode(loginResp)
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
