package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nurseIT2/library/internal/db"
	"github.com/nurseIT2/library/internal/routes"
)

func main() {
	db.InitDB()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8083")
}
