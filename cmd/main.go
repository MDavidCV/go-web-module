package main

import (
	"os"

	"github.com/MDavidCV/go-web-module/cmd/server"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	API_KEY := os.Getenv("API_KEY")

	cfg := &server.ConfigSeverChi{
		ServerAddress:  ":" + PORT,
		LoaderFielPath: "/Users/dcastrillonv/Documents/meli-boootcamp/go/go-web/go-web-module/docs/db/products.json",
		Token:          API_KEY,
	}

	app := server.NewServerChi(cfg)

	if err := app.Run(); err != nil {
		panic(err)
	}
}
