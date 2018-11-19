package structs

import "github.com/roomies-proyects/roomies-gastos/src/db"

//Expens is a struct for expenses
type Expens struct {
	tableName   struct{} `sql:"expenses"`
	ID          int      `sql:",pk"`
	Description string   `json:"description"`
	Cost        float64  `json:"cost"`
	RegistredBy string   `json:"registred_by"`
	State       string   `json:"state"`
	TypeCost    string   `json:"type_cost"`
}

//Save se guarda
func (e Expens) Save() error {
	db := db.Connect()
	defer db.Close()
	return db.Insert(&e)
}
