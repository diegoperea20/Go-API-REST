package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK , albums) 
	//Serializando el JSON
}

func postAlbums(c *gin.Context){
	var newAlbum album
	//binJson siver para pasar un json a una structura(clase)
	errr:=c.BindJSON(&newAlbum)
	if errr!=nil{
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated , newAlbum) 
	//Serializando el JSON
}

func getAlbumByID(c *gin.Context){
	id:=c.Param("id")
	for _,a:=range albums{
		if a.ID==id{
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func editAlbum(c *gin.Context) {
	id := c.Param("id")
	for i, a := range albums {
		if a.ID == id {
			var editionAlbum album
			// BindJSON is used to parse the JSON request and update the album
			err := c.BindJSON(&editionAlbum)
			if err != nil {
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Update the existing album with the new values
			albums[i] = editionAlbum

			// Return the updated album in the response
			c.IndentedJSON(http.StatusOK, editionAlbum)
			return
		}
	}
	// If the album with the specified ID is not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}


func deleteAlbum(c *gin.Context){

	id:=c.Param("id")
	for i, a:=range albums{
		if a.ID==id{
			albums=append(albums[:i],albums[i+1:]...)
			c.IndentedJSON(http.StatusOK,albums)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound,gin.H{"message":"album not found"})
}

func main() {
	router :=gin.Default()

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", editAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	router.Run("localhost:8080")//se puede definir el puerto en el que se va a ejecutar
}