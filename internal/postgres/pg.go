package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PgDb struct {
	Database *sql.DB
}

func NewPg() *PgDb {
	connStr := "user=naborly dbname=naborly password=naborly sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return &PgDb{Database: db}
}

func (pgDb *PgDb) Init() {
	err := pgDb.createUserTable()
	if err != nil {
		panic(err)
	}
}

func manyToManyTable(name string, relation int64) {

}
