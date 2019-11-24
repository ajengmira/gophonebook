package controllers

import (
	"net/http"
	"phonebook/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	contact := &models.Contact{}
	json.NewDecoder(r.Body).Decode(contact)

	if resp, ok := contact.Validate(); !ok {
		Respond(w, resp)
	}else {
		resp = contact.Create()
		Respond(w, resp)
	}
}

var GetUserContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")
	fmt.Printf("%s %d %s", r.URL.Path, user, "\n")
	params := mux.Vars(r)

	contacts := models.GetUserContact(params["id"])
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contacts" : contacts}
	Respond(w, resp)
}

var MyContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user")
	contacts := models.GetUserContact(user)
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contacts" : contacts}
	Respond(w, resp)
}

var GetContact = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	contact := models.GetContact(uint32(id))
	resp := map[string]interface{} {"status" : true, "message" : "Request success", "contact" : contact}
	Respond(w, resp)
}

var UpdateContact = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	contactName := r.FormValue("contact_name")
	phoneNumber := r.FormValue("phone_number")

	// get the param id
	mdl, err := models.GetContact(uint32(id))
	if err != nil {
		RespondError(w, "Cannot find contact", http.StatusUnprocessableEntity)
		return
	}

	contact := new(models.Contact)
	contact.Name = contactName
	contact.NameEN = phoneNumber

	resp = mdl.Update(contact)
	Respond(w, resp)
}

var DeleteContact = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 32)

	// get the param id
	contact := models.GetContact(uint32(id))

	resp = contact.Delete()
	Respond(w, resp)
}