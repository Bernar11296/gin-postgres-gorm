package api

import (
	"net/http"

	"example.com/go-gin-api/database"
	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Building RESTful API using Gin and Gorm",
	})
	return
}
func postArticle(c *gin.Context) {
	var article database.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.CreateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
	return
}
func getArticle(c *gin.Context) {
	id := c.Param("id")
	article, err := database.ReadArticle(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article": article,
	})
	return
}
func getArticles(c *gin.Context) {
	articles, err := database.ReadArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}
func putArticle(c *gin.Context) {
	var article database.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.UpdateArticle(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": res,
	})
	return
}
func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "article deleted successfully",
	})
	return
}
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home)
	r.GET("/api/v1/articles/:id", getArticle)
	r.GET("/api/v1/articles", getArticles)
	r.POST("/api/v1/articles", postArticle)
	r.PUT("/api/v1/articles/:id", putArticle)
	r.DELETE("/api/v1/articles/:id", deleteArticle)
	return r
}
