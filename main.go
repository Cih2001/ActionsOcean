package main

import (
	"Cih2001/ActionsOcean/controller"
	"Cih2001/ActionsOcean/model"
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// check arguments
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <server_listen_addr> <db_username> <db_password>\n", os.Args[0])
		fmt.Printf("e.g.: %s :1323 \n", os.Args[0])
		fmt.Println("provided arguments:", os.Args)
		return
	}

	serverAddress := os.Args[1]
	dbUsername := os.Args[2]
	dbPassword := os.Args[3]

	// initialize the db
	model.InitializeDB(dbUsername, dbPassword)

	// start the server
	// we use labstack echo framework.
	e := echo.New()

	controller.InitializeRoutes(e)

	// print logs on stdout
	e.Use(middleware.Logger())

	if err := e.Start(serverAddress); err != nil {
		fmt.Printf("Error listening on %s with error:%s\n", serverAddress, err.Error())
	}
}
