package main

import (
	_ "gophonebook/models"
	"gophonebook/controllers"
	"github.com/gorilla/mux"
	"net/http"

    //"github.com/go-sql-driver/mysql"
)
func main()  {

	router := mux.NewRouter()
	router.HandleFunc("/api/contacts", controllers.GetContacts).Methods("GET")
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.GetContact).Methods("GET")
	router.HandleFunc("/api/contact", controllers.CreateContact).Methods("POST")
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.UpdateContact).Methods("PUT")
	router.HandleFunc("/api/contact/{id:[0-9]+}", controllers.DeleteContact).Methods("DELETE")

	err := http.ListenAndServe("127.0.0.1:9000", router)
	if err != nil {
		panic(err)
	}
}