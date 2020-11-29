package internal

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL
	_ "github.com/lib/pq"              // Postgres
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// Databases connects to the resources databases.
func Databases() map[string]*xorm.Engine {
	return map[string]*xorm.Engine{
		"macapa":  configDB("macapa"),
		"varejao": configDB("varejao"),
	}
}

func configDB(s string) *xorm.Engine {
	var (
		db  *xorm.Engine
		err error
	)

	switch s {
	case "macapa":
		db, err = xorm.NewEngine("mysql", "admin:admin@/admin")
		db.SetMapper(names.GonicMapper{})
		if err = db.Sync2(new(Macapa)); err != nil {
			log.Fatalln(err)
		}

	case "varejao":
		db, err = xorm.NewEngine("postgres", "postgres://admin:admin@localhost:5432/postgres?sslmode=disable")
		db.SetMapper(names.GonicMapper{})
		if err = db.Sync2(new(Varejao)); err != nil {
			log.Fatalln(err)
		}
	}

	return db
}

// func tables(s string, db *xorm.Engine) *xorm.Engine {
// 	var u *xorm.Session

// 	db.SetMapper(names.GonicMapper{})

// 	switch s {
// 	case "macapa":
// 		u = db.Table(new(Macapa))
// 	case "varejao":
// 		u = db.Table(new(Varejao))
// 	}
// 	// u.Engine().TableName("users")

// 	if err := db.Sync2(u); err != nil {
// 		log.Fatalln(err)
// 	}

// 	return db
// }

// func setupDB(business string) *xorm.Engine {
// 	// master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/master?sslmode=disable")
// 	db, err := configDB(business)
// 	// eg, err := xorm.NewEngineGroup(master, db)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return db
// }
