package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Contact struct {

	ContactName string `json:"contact_name"`
	PhoneNumber string `json:"phone_number"`
	gorm.Model
}

func (con *Contact) Validate() (map[string]interface{}, bool) {

	resp := make(map[string] interface{})

	if con.ContactName == "" {
		resp["status"] = false
		resp["message"] = "Contact name is required."
		return resp, false
	}

	if con.PhoneNumber == "" {
		resp["status"] = false
		resp["message"] = "Phone number is required."
		return resp, false
	}

	resp["status"] = true
	resp["message"] = "All required fields(s) present"
	return resp, true
}

func (contact *Contact) Create() (map[string]interface{}) {

	resp := make(map[string]interface{})
	con := &Contact{}
	GetConn().Where("phone_number = ?", contact.PhoneNumber).First(con)
	if con.PhoneNumber != "" {
		resp["status"] = false
		resp["message"] = "This contact already exists"
		return resp;
	}

	GetConn().Create(contact)
	resp["status"] = true
	resp["message"] = "Contact has been created."
	resp["contact"] = GetContact(contact.ID)
	return resp
}

func (contact *Contact) Update(data *Contact) (map[string]interface{}) {
	
	resp := make(map[string]interface{})	

	GetConn().Model(&contact).Updates(data)
	resp["status"] = true
	resp["message"] = "Contact has been updated."
	resp["contact"] = GetContact(contact.ID)
	fmt.Println(data)
	return resp
}

func (contact *Contact) Delete(data *Contact) (map[string]interface{}) {

	resp := make(map[string]interface{})

	GetConn().Where("id = ?", data.ID).Delete(&contact)
	resp["status"] = true
	resp["message"] = "Contact has been deleted."
	resp["contact"] = GetContact(contact.ID)
	return resp
}
