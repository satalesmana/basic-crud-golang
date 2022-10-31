package main

import (
	"app-basic-crud/app/config"
	"app-basic-crud/app/database"
	"app-basic-crud/app/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	//cek argumen
	argLength := len(os.Args[1:])
	if argLength > 0 {
		if os.Args[1] == "migrate" {
			database.Migrate()
		}

	}

	//jalankan server
	if argLength == 0 {
		cfg := config.GetConfig()
		r := gin.Default()
		routes := routes.Routes(r)
		routes.Run(cfg.App.Host + ":" + cfg.App.Port)
	}
}
