package main

import (
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
	e := structs.Expens{Description: "500", Cost: 1, TypeCost: 1, RegistredBy: 1}
	err := e.Save()
	if err != nil {
		c.String(400, "400-Bad-Request"+err.Error())
		return
	}
	c.JSON(200, gin.H{
		"res": "Se a guardado el gasto",
	})
	return
}
