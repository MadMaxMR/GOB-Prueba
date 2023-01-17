# Products-Restful
***
_This project is a exercise of a RestFull Api_

## Starting 🚀

_These instructions will allow you to obtain a copy of the project on your local machine in order to run it._

### Requirements 📋

_This project was built with Golang and the Gin framework, additionally I used Gorm ORM with a Postgresql database._
_You must have Go and Postgresql installed on your computer_

```
Golang
GO mod
Gin
Gorm
Postgresql
```

### Instalación 🔧

_**FIRST STEP** clone the project on your local machine whit the following command_

```
    git clone https://github.com/MadMaxMR/Products-Restful.git
```

_Then create a local database whit the name **productDB**._

_Then go to your Products-Restful folder, inside you will find a folder called "database" and inside a file called "connection.go", inside the file change the following variables:_
```
var host = "yourhost"
var port = "yourport"
var user = "youruser"
var dbname = "productDB" // nombre de la base de datos
var password = "yourpassword"
```

_Then in the main folder run the command go run main.go to run the project-_
```
    go run main.go
```
_The endpoint are_
```
*Create a Product* HTTP Method POST
    http://localhost:8080/products 
*Get a Product* HTTP Method GET
    http://localhost:8080/products/:sku
*Get all Products* HTTP Method GET
    http://localhost:8080/products
*Update a Product* HTTP Method PUT
    http://localhost:8080/products/:sku
*Delete a Product* HTTP Method DELETE
    http://localhost:8080/products/:sku
```
_You can see the documentation of the endpoints in the following link:_
* [Documenter Postman](https://documenter.getpostman.com/view/19456004/2s8ZDU6QRF)

## Ejecutando las pruebas ⚙️

_Explica como ejecutar las pruebas automatizadas para este sistema_
