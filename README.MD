# GoLang Rest Api Development

A simple phonebook/contacts manager.

## API ENDPOINTS

### Create User
- Path : `/api/account`
- Method: `POST`
- Fields: `username, email, password`
- Response: `200`

### Login
- Path : `/api/login`
- Method: `POST`
- Fields: `username, password`
- Response: `200`

### All Contact By User Id
- Path : `/api/user/{id}/contacts`
- Method: `GET`
- Response: `200`

### Create Contact
- Path : `/api/contact`
- Method: `POST`
- Fields: `contact_name, phone_number, user_id`
- Response: `200`

### Details a Contact
- Path : `/api/contact/{id}`
- Method: `GET`
- Response: `200`

### Update Contact
- Path : `/api/contact/{id}`
- Method: `PUT`
- Fields: `contact_name, phone_number`
- Response: `200`

### Delete Contact
- Path : `/api/contact/{id}`
- Method: `DELETE`
- Response: `200`

## Required Packages
- Dependency management
    * [dep](https://github.com/golang/dep)
    * [migrate](https://github.com/mattes/migrate/cli)
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- ORM   
    * [gorm](github.com/jinzhu/gorm)
- Routing
    * [mux](github.com/gorilla/mux)

## Quick Run Project
First clone https://github.com/ajengmira/gophonebook or copy source then go to phonebook folder.  

```
cd phonebook

go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]

go build -tags 'mysql' -o install_path github.com/mattes/migrate/cli

migrate -path ./db/migrations -database mysql://user:password@tcp(host:port)/dbname up

go run main.go
```
