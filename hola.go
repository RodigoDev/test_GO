/*package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID       string  `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Telefono float64 `json:"telefono"`
}

var albums = []album{
	{ID: "1", Nombre: "Miguel", Apellido: "Angel", Telefono: 1111},
	{ID: "2", Nombre: "Nolberto", Apellido: "Beto", Telefono: 2222},
	{ID: "3", Nombre: "Rene", Apellido: "Reno", Telefono: 3333},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
} */
