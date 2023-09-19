package main

import (
	"api-gin-gorm/configs"
	"api-gin-gorm/controllers"
	"api-gin-gorm/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dbClient := configs.ConnectPostgreSQL()

	postMigrate := migrations.PostPostgreSQLDB{DB: dbClient}
	postMigrate.MigrateTable()

	pc := controllers.PostPostgreSQLDB{DB: dbClient}

	r.GET("/api/posts", pc.GetAllPosts)
	r.GET("/api/posts/:id", pc.GetPost)
	r.POST("/api/posts", pc.AddPost)
	r.PATCH("/api/posts/:id", pc.UpdatePost)
	r.DELETE("/api/posts/:id", pc.DeletePost)

	r.Run(":8080")
}
