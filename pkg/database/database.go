package database

import (
	"database/sql"

	_ "github.com/lib/pq"

	"typicalypetprojects/pkg/logging"
)

func ConnectPg() (*sql.DB, error) {
	// Строка подключения к базе данных PostgreSQL
	connStr := "user=root dbname=test_db password=root host=172.18.0.2 port=5432 sslmode=disable"

	// Открываем соединение с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logging.ErrorMessage("ConnectPg", err.Error(), 1)
	}
	defer db.Close()

	// Проверяем, что соединение с базой данных успешно установлено
	err = db.Ping()
	if err != nil {
		logging.ErrorMessage("ConnectPg", err.Error(), 2)
	}

	logging.SuccessMessage("ConnectPg")
	return db, nil
}
