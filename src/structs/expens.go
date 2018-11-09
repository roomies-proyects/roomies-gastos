package structs

import "github.com/roomies-proyects/roomies-gastos/src/db"

//Expens is a struct for expenses
type Expens struct {
	tableName   struct{} `sql:"expenses"`
	ID          int      `sql:",pk"`
	Description string
	Cost        float64
	RegistredBy int64
	TypeCost    int8
}

//Save se guarda
func (e Expens) Save() error {
	db := db.Connect()
	defer db.Close()
	return db.Insert(&e)
}
