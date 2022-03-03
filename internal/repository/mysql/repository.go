package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLRepository struct {
	db *sql.DB
}

func (r *MySQLRepository) ChangeRole(id int64, role int) error {
	_, err := r.db.Exec("UPDATE Users SET role = ? WHERE id = ?", role, id)
	return err
}

func (r *MySQLRepository) Create(id int64) error {
	_, err := r.db.Exec("INSERT INTO Users (id, role) VALUES (?, -1)", id)
	return err
}

func NewMySQLRepository(dataSourceName string) (*MySQLRepository, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &MySQLRepository{
		db: db,
	}, nil
}
