package db

import (
	"database/sql"
	"os"

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
	CreateImage(*data.Image) (int, error)
	GetImages() ([]*data.Image, error)
	GetImageByID(int) (*data.Image, error)
	DeleteImage(int) error
	CreateResult(*data.Result) (int, error)
	GetResults() ([]*data.Result, error)
	GetResultByID(int) (*data.Result, error)
	GetResultsByUserID(int) ([]*data.Result, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostGresStore() (*PostgresStore, error) {
	connStr := "host=" + os.Getenv("DBHOST") + " port=" + os.Getenv("DBPORT") + " user=" + os.Getenv("DBUSER") + " password=" + os.Getenv("DBPASSWORD") + " dbname=" + os.Getenv("DBNAME") + " sslmode=disable"

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
	if err := s.createPredictResultTable(); err != nil {
		return err
	}
	if err := s.createImageTable(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS Account (
		user_id SERIAL PRIMARY KEY,
		name varchar(100) NOT NULL,
		email varchar(100) UNIQUE NOT NULL,
		encrypted_password bpchar NOT NULL,
		created_at timestamp NOT NULL
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createDiseaseTable() error {
	query := `CREATE TABLE IF NOT EXISTS Disease (
		disease_id SERIAL PRIMARY KEY,
		disease_name varchar(100) UNIQUE NOT NULL,
		plant_name varchar(100) NOT NULL,
		description JSONB NOT NULL,
		created_at timestamp NOT NULL
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createPredictResultTable() error {
	query := `CREATE TABLE IF NOT EXISTS PredictResult (
		result_id SERIAL PRIMARY KEY,
		user_id INTEGER REFERENCES Account(user_id) NOT NULL,
		disease_id INTEGER REFERENCES Disease(disease_id),
		predict_result BOOLEAN NOT NULL,
		created_at timestamp NOT NULL
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) createImageTable() error {
	query := `CREATE TABLE IF NOT EXISTS Image (
		image_id SERIAL PRIMARY KEY,
		result_id INTEGER REFERENCES PredictResult(result_id) NOT NULL,
		image_url TEXT NOT NULL,
		created_at timestamp NOT NULL
	)`

	_, err := s.db.Exec(query)
	return err
}
