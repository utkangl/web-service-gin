package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("127.0.0.1:8080")

}

// struct to represent the datas about albums
type album struct {
	ID     string  `json:"id"` // why do we use `` ??
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to create album data
var albums = []album{

	{ID: "0", Title: "Ah Be Abi", Artist: "Mehmet Emre Arslan", Price: 59.99},
	{ID: "1", Title: "Tavuklu Çoraplar", Artist: "Samet Kati", Price: 52.99},
	{ID: "2", Title: "Sanirim Deliriyoruz", Artist: "Mertcan Kirci", Price: 69.99},
}

// the function responds with the list of albums as json file
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// the func to add a new album from json to albums slice
func postAlbums(c *gin.Context) {

	// create an object of album struct
	var newAlbum album

	// if new object returns a error within BindJSON function, return nothing
	if err := c.BindJSON(&newAlbum); err != nil {
		return

		// else if error doesn't occur, add the new album to albums slice of album struct
	} else if err := c.BindJSON(&newAlbum); err == nil {

		albums = append(albums, newAlbum)

		// IndentedJSON function takes an integer parameter(httpstatus as integer)
		// and an object, serializes the given object as JSON to convert to stream bytes and store it
		c.IndentedJSON(http.StatusCreated, newAlbum)
	}

}

// the func to get a specific album by its id
func getAlbumByID(c *gin.Context) {

	// ??????
	id := c.Param("id")

	// if parameter id is equal to an album's id, return the album
	for _, a := range albums {

		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// if requested album does not exist, display a message
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
