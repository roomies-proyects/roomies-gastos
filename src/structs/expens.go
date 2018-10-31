package structs

//Expens is a struct for expenses
type Expens struct {
	ID          int
	Description string
	Cost        float64
	RegistredBy int64
	TypeCost    int8
}
