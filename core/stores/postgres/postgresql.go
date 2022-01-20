package postgres

import (
	// imports the driver, don't remove this comment, golint requires.
	"github.com/MockyBang/go-zero/core/stores/sqlx"
	_ "github.com/lib/pq"
)

const postgresDriverName = "postgres"

// New returns a postgres connection.
func New(datasource string, opts ...sqlx.SqlOption) sqlx.SqlConn {
	return sqlx.NewSqlConn(postgresDriverName, datasource, opts...)
}
