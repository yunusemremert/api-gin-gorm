package controllers

import (
	"api-gin-gorm/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type PostPostgreSQLDB struct {
	DB *gorm.DB
}

func (p PostPostgreSQLDB) GetAllPosts(c *gin.Context) {
	var posts []models.Post

	if result := p.DB.Find(&posts); result.Error != nil {
		fmt.Println(result.Error)
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func (p PostPostgreSQLDB) GetPost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	if result := p.DB.First(&post, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	c.IndentedJSON(http.StatusOK, post)
}

func (p PostPostgreSQLDB) AddPost(c *gin.Context) {
	var newPost models.Post

	if err := c.BindJSON(&newPost); err != nil {
		return
	}

	if result := p.DB.Create(&newPost); result.Error != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": result.Error})
	}

	c.IndentedJSON(http.StatusCreated, newPost)
}

func (p PostPostgreSQLDB) UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	if result := p.DB.First(&post, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	var updatePost models.Post

	if err := c.BindJSON(&updatePost); err != nil {
		return
	}

	p.DB.Model(&post).Updates(&updatePost)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Updated post"})
}

func (p PostPostgreSQLDB) DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	if result := p.DB.First(&post, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	// p.DB.Delete(&post) // deleted_at update
	p.DB.Delete(&post, id) // remove

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Deleted post"})
}
