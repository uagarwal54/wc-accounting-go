package login

import(
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	
	"wc-accounting-go/codebase/model"
)

type (
	LoginResp struct{
		Message string
		StatusCode int
		FirstLogin string 
	}
)

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
	if user, err = model.ReadUserDataWithUsernamePassword(userCredMap["username"], userCredMap["password"]); err != nil {
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

func AddLoginRoute(router *mux.Router){
	router.HandleFunc("/login", login)
}