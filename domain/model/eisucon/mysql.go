package eisucon

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Singleton field
var dsn string

func Init(user string, password string, host string, port uint, db string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true", user, password, host, port, db)
}

// MySQLサーバーに接続
func OpenMysql() (*sqlx.DB, error) {
	//if dsn == "" {
	//	return nil, errors.New("dsn does not set")
	//}
	return sqlx.Open("sqlite3", "./prc_hub.db")
}
