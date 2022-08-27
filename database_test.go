package golang_mysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	db, err := sql.Open("mysql", "root:VasylRagnar32145@tcp(localhost:3306)/golang_mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
