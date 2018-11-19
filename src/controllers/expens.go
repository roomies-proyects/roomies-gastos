package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/roomies-proyects/roomies-gastos/src/structs"
)

//NewExpense se guarda
func NewExpense(c *gin.Context) error {
	cost, err := strconv.ParseFloat(c.PostForm("cost"), 64)
	if err != nil {
		return err
	}
	expense := structs.Expens{
		Description: c.PostForm("description"),
		Cost:        cost,
		RegistredBy: c.PostForm("registred_by"),
		TypeCost:    c.PostForm("type_cost"),
	}

	if err = expense.Save(); err != nil {
		return err
	}
	return nil
}

func GetResumReport() (structs.ResumReport, error) {
	resum := structs.ResumReport{Debt: 1, InFavor: 1}
	return resum, nil
}
