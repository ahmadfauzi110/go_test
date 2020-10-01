package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Handle("/get_user",  loggingMiddleware(http.HandlerFunc(returnAllUsers))).Methods("GET")
	router.HandleFunc("/login_user", loginUser).Methods("POST")
	router.HandleFunc("/register_user", registerUser).Methods("POST")
	router.Handle("/update_user", loggingMiddleware(http.HandlerFunc(updateUser))).Methods("PUT")
	router.Handle("/update_foto_user", loggingMiddleware(http.HandlerFunc(updateFotoUser))).Methods("PUT")
	router.Handle("/delete_user", loggingMiddleware(http.HandlerFunc(deleteUser))).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}