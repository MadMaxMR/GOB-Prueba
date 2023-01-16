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

_Finaliza con un ejemplo de cómo obtener datos del sistema o como usarlos para una pequeña demo_

