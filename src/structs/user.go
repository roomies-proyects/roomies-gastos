package structs

//TypeCost is a struct for expenses
type User struct {
	tableName struct{} `sql:"users"`
	ID        int      `sql:",pk"`
	UserName  string
	Email     string
}
