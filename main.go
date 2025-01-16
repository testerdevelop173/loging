package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Pirce  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "titleOne", Author: "authorOne", Pirce: 100.99},
	{ID: "2", Title: "titleTwo", Author: "authorTwo", Pirce: 200.99},
	{ID: "3", Title: "titleThree", Author: "authorThee", Pirce: 300.99},
}

func main() {
	path := gin.Default()
	path.GET("/albums", getAlbums)
	path.GET("/albums/:id", getAlbumByID)
	path.POST("/albums", creatAlbums)
	path.Run("localhost:8080")

}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)

}

func creatAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
