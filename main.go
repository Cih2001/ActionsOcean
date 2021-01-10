package main

import (
	"Cih2001/ActionsOcean/controller"
	"Cih2001/ActionsOcean/model"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

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
	e.Validator = &CustomValidator{validator: validator.New()}

	controller.InitializeRoutes(e)

	// print logs on stdout
	e.Use(middleware.Logger())

	// disabling ssl certificates check. we are running this inside a base container with
	// no cert DB.
	// not disabling ssl certificates, our requests to employee api server will fail with err:
	// x509: certificate signed by unknown authority
	// CAUTION:Disabling security checks is dangerous and should be avoided
	// TODO: fix by using a modified container or adding certs manually.
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	if err := e.Start(serverAddress); err != nil {
		fmt.Printf("Error listening on %s with error:%s\n", serverAddress, err.Error())
	}
}
