package internal

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL
	_ "github.com/lib/pq"              // Postgres
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// var eg *xorm.EngineGroup

// Database connects to the database.
func Database() (*xorm.EngineGroup, error) {
	master, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/master?sslmode=disable")
	varejaoPG, err := xorm.NewEngine("postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	macapaMY, err := xorm.NewEngine("mysql", "jdbc:mysql//admin:admin@tcp(localhost:3306)/admin?sslmode=disable")

	dbs := []*xorm.Engine{varejaoPG, macapaMY}
	eg, err := xorm.NewEngineGroup(master, dbs)
	eg.SetMapper(names.SameMapper{})

	tbs, err := eg.DBMetas()
	fmt.Println(tbs)

	if err != nil {
		return nil, err
	}

	return eg, nil
}
