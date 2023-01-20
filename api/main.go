package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Album struct {
	id     string  `json:"id"`
	title  string  `json:"title"`
	artist string  `json:"artist"`
	price  float64 `json:"price"`
}

var Database = []Album{
	{"1234", "Gospel 2022", "Unknown", 12.95},
	{"2345", "Hip Hop 2022", "Unknown", 13.95},
	{"3456", "Lofi 2022", "Unknown", 17.95},
	{"4567", "Classical 2022", "Unknown", 14.95},
}

func pageNotFound(c *gin.Context) {
	c.String(404, "Opps! The page you're looking for does not exist")
}

func home(c *gin.Context) {
	c.String(200, "Welcome to the album api")
}

// todo
func getAllAlbums(c *gin.Context) {
	fmt.Println(c)
	c.JSON(200, Database)
}

// todo
func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, item := range Database {
		if id == item.id {
			c.JSON(200, item)
			return
		}
	}

	c.String(400, "Album could not be found")
}

// todo
func postAlbum(c *gin.Context) {
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.String(400, "Invalid post parameters")
		return
	}

	Database = append(Database, newAlbum)

	c.JSON(200, gin.H{
		"message": "successfully created",
		"album":   newAlbum,
	})
}

// todo
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.NoRoute(pageNotFound)
	router.GET("/", home)
	router.GET("/albums", getAllAlbums)
	router.GET("/album/:id", getAlbum)
	router.POST("/createAlbum", postAlbum)

	return router
}

func main() {
	SetupRouter().Run(":8000")
}
