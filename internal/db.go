package internal

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL
	_ "github.com/lib/pq"              // Postgres
	"xorm.io/xorm"
)

var eg *xorm.EngineGroup

// Database connects to the database.
func Database() (*xorm.EngineGroup, error) {
	pg, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		// return nil, err
		log.Fatalln(err)
	}
	// jdbc:mysql:
	// mysqlx:
	my, err := xorm.NewEngine("mysql", "jdbc:mysql//admin:admin@tcp(localhost:3306)/admin?sslmode=disable")
	if err != nil {
		// return nil, err
		log.Fatalln(err)
	}

	return xorm.NewEngineGroup(pg, my)
}

// // Database connects to the database.
// func Database() (*xorm.EngineGroup, error) {
// 	master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
// 	if err != nil {
// 		return nil, err
// 	}
// 	varejaoPG, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/test?sslmode=disable")
// 	if err != nil {
// 		return nil, err
// 	}
// 	macapaMY, err := xorm.NewEngine("mysql", "postgres://postgres:root@localhost:5432/test1?sslmode=disable")
// 	if err != nil {
// 		return nil, err
// 	}
// 	dbs := []*xorm.Engine{varejaoPG, macapaMY}

// 	return xorm.NewEngineGroup(master, dbs)
// }
