# GoLang Rest Api Development

A simple phonebook/contacts manager.

## API ENDPOINTS

### All Contact
- Path : `/api/contacts`
- Method: `GET`
- Response: `200`

### Contact By Id
- Path : `/api/contact/{id}`
- Method: `GET`
- Response: `200`

### Create Contact
- Path : `/api/contact`
- Method: `POST`
- Fields: `contact_name, phone_number`
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
    * [migrate](https://github.com/golang-migrate/migrate)
- Database
    * [MySql](https://github.com/go-sql-driver/mysql)
- ORM   
    * [gorm](https://github.com/jinzhu/gorm)
- Routing
    * [mux](https://github.com/gorilla/mux)

## Quick Run Project
First clone https://github.com/ajengmira/gophonebook or copy source then go to phonebook folder.  

```
##go to your db editor
create database "phonebook"

##go to your console
cd phonebook

##go get [-d] [-t] [-u] [-v] [-insecure] [build flags] [packages]

go get -u github.com/golang/dep/cmd/dep
go get -u github.com/golang-migrate/migrate
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
go get -u github.com/gorilla/mux
go get -u github.com/mattes/migrate
go get -u github.com/joho/godotenv

go build -tags 'mysql' -o install_path github.com/golang-migrate/migrate

go build -tags 'mysql' -o c://go_project/bin/migrate github.com/mattes/migrate/cli

migrate -path ./db/migrations -database mysql://user:password@tcp(host:port)/dbname up

go run main.go
```
