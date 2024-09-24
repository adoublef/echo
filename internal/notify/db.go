package notify

import "github.com/jmoiron/sqlx"

type DB struct {
	RWC *sqlx.DB
}
