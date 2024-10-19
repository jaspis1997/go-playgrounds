package line

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

type Database = *bun.DB

func NewDatabase(filename string) Database {
	sqlite, err := sql.Open(sqliteshim.ShimName, filename)
	if err != nil {
		panic(err)
	}
	return bun.NewDB(sqlite, sqlitedialect.New())
}
