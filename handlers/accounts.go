package handlers

import (
	"encoding/json"
	"net/http"
	accountdto "test_backend/dto/account"
	dto "test_backend/dto/result"
	"test_backend/models"
	"test_backend/pkg/bcrypt"
	"test_backend/repositories"

	"github.com/go-playground/validator/v10"
)

type handlerAccount struct {
	AccountRepository repositories.AccountRepository
}

func HandlerAccount(AccountRepository repositories.AccountRepository) *handlerAccount {
	return &handlerAccount{AccountRepository}
}

func (h *handlerAccount) CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(accountdto.RequestAccount)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	account := models.Account{
		Name:     request.Name,
		Password: password,
	}

	data, err := h.AccountRepository.CreateAccount(account)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerAccount) FindAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")

	accounts, err := h.AccountRepository.FindAccounts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: accounts}
	json.NewEncoder(w).Encode(response)
}
