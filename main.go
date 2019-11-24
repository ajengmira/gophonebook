package main

import (
	_ "phonebook/models"
	"phonebook/controllers"
	"github.com/gorilla/mux"
	"net/http"

	"phonebook/app"
)
func main()  {

	router := mux.NewRouter()
	router.HandleFunc("/api/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/account", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/auth", controllers.Login).Methods("POST")
	router.HandleFunc("/api/contact", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/api/user/{id:[0-9]+}/contacts", controllers.GetUserContact).Methods("GET")
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.GetContact).Methods("GET")
	router.HandleFunc("/api/me/contacts", controllers.MyContact).Methods("GET")
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.DeleteContact).Methods("DELETE")

	router.Use(app.JwtMiddleWare)

	err := http.ListenAndServe("127.0.0.1:9000", router)
	if err != nil {
		panic(err)
	}
}