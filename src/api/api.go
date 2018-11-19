package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roomies-proyects/roomies-gastos/src/controllers"
)

//main levamta el servicio Api
func main() {
	r := gin.Default()
	r.GET("/get_resum_report", GetResumReport)
	r.POST("/new_expens", newExpense)
	r.Run(":8000")
}

func newExpense(c *gin.Context) {
	err := controllers.NewExpense(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "inserted correctly"})

}

func GetResumReport(c *gin.Context) {
	resum, err := controllers.GetResumReport()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resum)

}
