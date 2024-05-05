package db

import (
	"database/sql"
	"fmt"

	"github.com/Narutchai01/ProjectC-BE/data"
)

func (s *PostgresStore) CreateAccount(acc *data.Account) (int, error) {
	query := `INSERT INTO Account 
		(name, email, encrypted_password, created_at)
		VALUES ($1, $2, $3, $4) RETURNING user_id`

	var id int
	err := s.db.QueryRow(
		query,
		acc.Name, acc.Email, acc.EncryptedPassword, acc.CreatedAt).Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *PostgresStore) UpdateAccount(acc *data.Account) error {
	query := `UPDATE Account 
		SET name = $2, email = $3, encrypted_password = $4 
		WHERE user_id = $1`

	_, err := s.db.Query(query, acc.ID, acc.Name, acc.Email, acc.EncryptedPassword)
	return err
}

func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("DELETE FROM Account WHERE user_id = $1", id)
	return err
}

func (s *PostgresStore) GetAccountByID(id int) (*data.Account, error) {
	rows, err := s.db.Query("SELECT * FROM Account WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccountByEmail(email string) (*data.Account, error) {
	rows, err := s.db.Query("SELECT * FROM Account WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("account with email [%s] not found", email)
}

func (s *PostgresStore) GetAccounts() ([]*data.Account, error) {
	rows, err := s.db.Query("SELECT * FROM Account")
	if err != nil {
		return nil, err
	}

	accounts := []*data.Account{}
	for rows.Next() {
		account, err := scanIntoAccount(rows)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanIntoAccount(rows *sql.Rows) (*data.Account, error) {
	account := new(data.Account)
	err := rows.Scan(
		&account.ID,
		&account.Name,
		&account.Email,
		&account.EncryptedPassword,
		&account.CreatedAt)

	return account, err
}
