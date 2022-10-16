package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost/8000")

}

// struct to represent the datas about albums
type album struct {
	ID     string  `json:"id"` // why do we use `` ??
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"prica"`
}

// albums slice to create album data
var albums = []album{

	{ID: "0", Title: "Ah Be Abi", Artist: "Mehmet Emre Arslan", Price: 59.99},
	{ID: "1", Title: "Tavuklu Çoraplar", Artist: "Samet Kati", Price: 52.99},
	{ID: "2", Title: "Sanirim Deliriyoruz", Artist: "Mertcan Kirci", Price: 69.99},
}

// the funciton responds with the list of albums as json file
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
