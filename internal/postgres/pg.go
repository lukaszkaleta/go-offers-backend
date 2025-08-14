package postgres

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type PgDb struct {
	Database *sql.DB
}

func (db *PgDb) ExecuteSqls(sqls []string) error {
	for _, sql := range sqls {
		err := db.ExecuteSql(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *PgDb) ExecuteSql(sql string) error {
	_, err := db.Database.Exec(sql)
	return err
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

func (db *PgDb) tableEntity(name string, id int) TableEntity {
	return TableEntity{Name: name, Id: id}
}

func ExecuteFromFile(path string) {
	sqlStatements, err := os.ReadFile(path)
	ifPanic(err)
	sqlArray := strings.Split(string(sqlStatements), ";")
	ifPanic(NewPg().ExecuteSqls(sqlArray))
}

func ifPanic(e error) {
	if e != nil {
		panic(e)
	}
}
