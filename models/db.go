package models

func GetContact(Id uint) *Contact {

	contact := &Contact{}
	GetConn().First(contact, Id)
	return contact
}

func GetContacts() *Contact {

	contact := &Contact{}
	GetConn().Order("id desc").Find(&contact)
	return contact
}
