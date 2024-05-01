package db

import (
	"database/sql"

	"github.com/Narutchai01/ProjectC-BE/data"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*data.Account) (int, error)
	DeleteAccount(int) error
	UpdateAccount(*data.Account) error
	GetAccountByID(int) (*data.Account, error)
	GetAccounts() ([]*data.Account, error)
	GetAccountByEmail(string) (*data.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostGresStore() (*PostgresStore, error) {
	connStr := "user=admin password=admin123 dbname=kmutt sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS Account (
		user_id SERIAL PRIMARY KEY,
		name varchar(100),
		email varchar(100),
		encrypted_password bpchar,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}
