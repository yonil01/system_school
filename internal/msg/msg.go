package msg

import (


	"github.com/jmoiron/sqlx"
)
type Model struct {
	db *sqlx.DB
}


func GetByCode(code int, typ string, msg string) (int, string, string) {
	return code, typ, msg
}
