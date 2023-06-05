package repository

import "database/sql"

type OrderRepository struct {
	db *sql.DB
}
