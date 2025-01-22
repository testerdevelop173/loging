package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// структура Album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// хранение альбома
var Albums = map[string]*Album{
	"1": {ID: "1", Title: "titleOne", Author: "authorOne", Price: 100.99},
	"2": {ID: "2", Title: "titleTwo", Author: "authorTwo", Price: 200.99},
	"3": {ID: "3", Title: "titleThree", Author: "authorThree", Price: 300.99},
}

// добавляет новый альбом
func CreateAlbum(c *gin.Context) {
	var newAlbum Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		Logger().Printf("Произошла ошибка при создании альбома: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Albums[newAlbum.ID] = &newAlbum
	Logger().Printf("Создан новый альбом: %+v\n", newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// поиск альбома по ID
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Проверяем, существует ли альбом с таким ID
	album, ok := Albums[id]
	if !ok {
		Logger().Printf("Альбом с ID=%s не найден\n", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}

	Logger().Printf("Запрошен альбом с ID=%s\n", id)
	c.IndentedJSON(http.StatusOK, album)
}

// вывести все альбомы
func GetAlbums(c *gin.Context) {
	var allAlbums []*Album
	for _, album := range Albums {
		allAlbums = append(allAlbums, album)
	}

	Logger().Printf("Получены все альбомы\n")
	c.IndentedJSON(http.StatusOK, allAlbums)
}
