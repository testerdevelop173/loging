package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", GetAlbums)
	router.GET("/albums/:id", GetAlbumByID)
	router.POST("/albums", CreateAlbum)
	router.Run("localhost:8080")
}

/*
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}
*/
