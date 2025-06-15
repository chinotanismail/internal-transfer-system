package main

import (
	"github.com/chinotanismail/internal-transfer-system/config"
	"github.com/chinotanismail/internal-transfer-system/routes"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}
