package main

import (
	"log"

	"github.com/amrllkmn/apex/internal/activities"
	"github.com/amrllkmn/apex/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := database.InitDB()
	err := db.AutoMigrate(&activities.Activity{})
	if err != nil {
		log.Fatal("Migration failed")
	}
	r = activities.SetupRouter(r, db)
	r.Run(":8081") // Run and serve on 8080
}
