package user

import (
	"log"

	"github.com/Ocelani/mercafacil/internal"
)

type (
	// Macapa defines the User at the provided database.
	Macapa struct {
		ID        int
		Name      string `xorm:"VARCHAR ( 200 ) NOT NULL"`
		Cellphone string `xorm:"VARCHAR ( 20 ) NOT NULL)"`
	}

	// Varejao defines the User at the provided database.
	Varejao struct {
		ID        int
		Name      string `xorm:"VARCHAR ( 100 ) NOT NULL"`
		Cellphone string `xorm:"VARCHAR ( 13 ) NOT NULL)"`
	}
)

func database() {
	db, err := internal.Database()
	if err != nil {
		log.Fatalln(err)
	}
	slaves := db.Slaves()
	sslaves[1]
	db.CreateTables()
}
