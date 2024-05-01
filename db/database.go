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
	CreateDisease(*data.Disease) (int, error)
	GetDiseases() ([]*data.Disease, error)
	GetDiseaseByID(int) (*data.Disease, error)
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
	if err := s.createAccountTable(); err != nil {
		return err
	}
	if err := s.createDiseaseTable(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createAccountTable() error {
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

func (s *PostgresStore) createDiseaseTable() error {
	query := `CREATE TABLE IF NOT EXISTS Disease (
		disease_id SERIAL PRIMARY KEY,
		disease_name varchar(100),
		plant_name varchar(100),
		description JSONB,
		created_at timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}
