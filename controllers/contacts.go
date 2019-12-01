package controllers

import (
	"net/http"
	"gophonebook/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	contact := &models.Contact{}
	json.NewDecoder(r.Body).Decode(contact)

	if resp, ok := contact.Validate(); !ok {
		var jsonData, err = json.Marshal(resp)
		if err != nil {
		    fmt.Println(err.Error())
		    return
		}

		w.Header().Set("Content-Type", "application/json") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}else {
		resp = contact.Create()
		var jsonData, err = json.Marshal(resp)
		if err != nil {
		    fmt.Println(err.Error())
		    return
		}

		w.Header().Set("Content-Type", "application/json") // normal header
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData) 
	}
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {

	contact := models.GetContacts()
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contact" : contact}
	var jsonData, err = json.Marshal(resp)
	if err != nil {
	    fmt.Println(err.Error())
	    return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

var GetContact = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	contact := models.GetContact(uint(id))
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contact" : contact}
	 
	var jsonData, err = json.Marshal(resp)
	if err != nil {
	    fmt.Println(err.Error())
	    return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	contactName := r.FormValue("contact_name")
	phoneNumber := r.FormValue("phone_number")

	// get the param id
	mdl := models.GetContact(uint(id))
	
	contact := new(models.Contact)
	contact.ContactName = contactName
	contact.PhoneNumber = phoneNumber
	contacts := mdl.Update(contact)
	
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contacts" : contacts}
	var jsonData, err = json.Marshal(resp)
	if err != nil {
	    fmt.Println(err.Error())
	    return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	// get the param id
	mdl := models.GetContact(uint(id))

	contact := new(models.Contact)
	contact.ID = uint(id)
	contacts := mdl.Delete(contact)
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contacts" : contacts}
	var jsonData, err = json.Marshal(resp)
	if err != nil {
	    fmt.Println(err.Error())
	    return
	}

	w.Header().Set("Content-Type", "application/json") // normal header
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}