package structs

//TypeCost is a struct for expenses
type TypeCost struct {
	tableName   struct{} `sql:"typecosts"`
	ID          int      `sql:",pk"`
	Description string
}
