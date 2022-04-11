package main

import (
	"fmt"
	"log"

	"github.com/UfiairENE/bantu_solution/internal/config"
	"github.com/UfiairENE/bantu_solution/pkg/router"
	"github.com/UfiairENE/bantu_solution/pkg/router/connection"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.Setup()
	connection.ConnectToDB()

}

func main() {
	// loadconfig := config.LoadConfig()
	getConfig := config.GetConfig()
	validatorRef := validator.New()
	r := router.Setup(validatorRef)

	log.Printf("Server is starting at 127.0.0.1:%s", 8080)
	log.Fatal(r.Run(":8080"))

	fmt.Println(getConfig.Server.Port)

}
