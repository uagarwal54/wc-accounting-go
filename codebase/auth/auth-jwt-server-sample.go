package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mySuperSecretPhrase")

func main() {
	fmt.Println("My Simple Server")
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Info")
}

func handleRequests() {
	http.Handle("/login", isAuthorized(homePage))
	log.Fatal(http.ListenAndServe(":9001", nil))
}

// This function will check if all headers are in place and the values of the headers are correct
func isAuthorized(endPoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There is an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			if token.Valid {
				// If token is valid we are calling the actual function that will handle the  http request
				// i.e the function that was passed as parameter to the isAuthorized function
				endPoint(w, r)
			}
		} else {
			fmt.Println("Hit Recieved")
			bodyData, err := ioutil.ReadAll(r.Body)
			fmt.Println(string(bodyData))
			fmt.Println(err)
			w.Write([]byte("Not Authorized"))
		}
	})
}
