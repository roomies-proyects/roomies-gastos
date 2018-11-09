package main

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/roomies-proyects/roomies-gastos/src/db"
	"github.com/roomies-proyects/roomies-gastos/src/structs"
)

//CrearEsquema se crea el esquema Periodos
func main() {
	db := db.Connect()
	defer db.Close()
	//EliminarTablas(db)
	var t structs.TypeCost
	var e structs.Expens
	var u structs.User
	CrearEsquema(&u, db)
	CrearEsquema(&t, db)
	CrearEsquema(&e, db)
}

func CrearEsquema(model interface{}, db *pg.DB) error {
	db.Delete(model)
	err := db.CreateTable(model, &orm.CreateTableOptions{
		Temp:          false, // create temp table
		FKConstraints: true,
	})

	if err != nil {
		fmt.Printf("error en creacion de tablas %v\n", err.Error())
		return err
	}
	return nil
}
