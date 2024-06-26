package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Narutchai01/ProjectC-BE/data"
	"github.com/Narutchai01/ProjectC-BE/handlers"
	"github.com/Narutchai01/ProjectC-BE/util"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	loginReq := new(data.LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(loginReq); err != nil {
		return err
	}

	account, err := s.store.GetAccountByEmail(loginReq.Email)
	if err != nil {
		return err
	}

	if !account.ValidPassword(loginReq.Password) {
		return fmt.Errorf("invalid password")
	}

	token, err := handlers.CreateJWT(account)
	if err != nil {
		return err
	}

	loginResp := data.LoginResponse{
		ID:    account.ID,
		Token: token,
	}

	return util.WriteJSON(w, http.StatusOK, loginResp)
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	createAccountReq := new(data.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}

	existAccount, _ := s.store.GetAccountByEmail(createAccountReq.Email)
	if existAccount != nil {
		return fmt.Errorf("email has already been taken")
	}

	account, err := data.NewAccount(createAccountReq)
	if err != nil {
		return err
	}
	id, err := s.store.CreateAccount(account)
	if err != nil {
		return err
	}
	account.ID = id

	token, err := handlers.CreateJWT(account)
	if err != nil {
		return err
	}

	loginResp := data.LoginResponse{
		ID:    account.ID,
		Token: token,
	}

	return util.WriteJSON(w, http.StatusOK, loginResp)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetAccounts(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleAccountByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccountByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	if r.Method == "PATCH" {
		return s.handleUpdateAccount(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccounts()
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := util.GetID(r, "user")
	if err != nil {
		return err
	}

	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := util.GetID(r, "user")
	if err != nil {
		return err
	}

	if err := s.store.DeleteAccount(id); err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountReq := new(data.CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}

	account, err := data.NewAccount(createAccountReq)
	if err != nil {
		return err
	}

	account.ID, err = s.store.CreateAccount(account)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := util.GetID(r, "user")
	if err != nil {
		return err
	}

	updateAccountReq := new(data.UpdateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(updateAccountReq); err != nil {
		return err
	}
	updateAccountReq.ID = id

	account, err := s.store.GetAccountByID(id)
	if err != nil {
		return err
	}

	if updateAccountReq.Name != "" {
		account.Name = updateAccountReq.Name
	}
	if updateAccountReq.Email != "" {
		account.Email = updateAccountReq.Email
	}
	if updateAccountReq.Password != "" {
		encpw, err := data.EncrptPassword(updateAccountReq.Password)
		if err != nil {
			return err
		}
		account.EncryptedPassword = string(encpw)
	}
	err = s.store.UpdateAccount(account)
	if err != nil {
		return err
	}

	return util.WriteJSON(w, http.StatusOK, account)
}
