package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roomies-proyects/roomies-gastos/src/structs"
)

//main levamta el servicio Api
func main() {
	r := gin.Default()
	//r.GET("/clima", clima)
	//r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	/*r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")*/
	r.GET("/new_expens", newExpens)
	r.Run(":8000")
}

func newExpens(c *gin.Context) {
	var json structs.Expens
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := json.Save(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error inserting": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "inserted correctly"})
}
