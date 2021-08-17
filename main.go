package main

import (
	"log"
	"os"

	env "github.com/joho/godotenv"
	"github.com/oswaldom-code/web-cert-checker/routers"
)

//main function run a gin-gonic web server
func main() {
	//load .env file
	err := env.Load()
	if err != nil {
		log.Println("Error loading .env file -> ", err)
	}
	//listen on port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	//create routers
	engine := routers.SetupRouter()
	//start server
	if err := engine.Run(":" + port); err != nil {
		log.Panic("Error: ", err.Error())
	}
}
