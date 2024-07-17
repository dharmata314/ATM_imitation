package handlers

import (
	"ATM-service/internal/entities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var accounts = make(map[string]*entities.Account)
var mu sync.RWMutex

var logChannel = make(chan string, 100)

func StartLogger() {
	for logMessage := range logChannel {
		log.Println(logMessage)
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	id := generateID()
	account := &entities.Account{ID: id}
	mu.RLock()
	accounts[id] = account
	mu.RUnlock()

	logChannel <- fmt.Sprintf("Аккаунт id: %s создан в: в %s", id, time.Now())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var request struct {
		Amount float64 `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	mu.RLock()
	account, exists := accounts[id]
	mu.RUnlock()
	if !exists {
		http.Error(w, "Аккаунт не найден", http.StatusNotFound)
		return
	}

	account.Deposit(request.Amount)
	logChannel <- fmt.Sprintf("Баланс пополен %f на аккаунте id: %s в %s", request.Amount, id, time.Now())

	w.WriteHeader(http.StatusOK)
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var request struct {
		Amount float64 `json:"amount"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	mu.RLock()
	account, exists := accounts[id]
	mu.RUnlock()
	if !exists {
		http.Error(w, "Аккаунт не найден", http.StatusNotFound)
		return
	}

	if err := account.Withdraw(request.Amount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	logChannel <- fmt.Sprintf("Снятие средств %f с аккаунта id: %s в %s", request.Amount, id, time.Now())

	w.WriteHeader(http.StatusOK)
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	mu.RLock()
	account, exists := accounts[id]
	mu.RUnlock()
	if !exists {
		http.Error(w, "Аккаунт не найден", http.StatusNotFound)
		return
	}

	balance := account.GetBalance()
	logChannel <- fmt.Sprintf("Проверка баланса аккаунта id: %s в %s", id, time.Now())

	json.NewEncoder(w).Encode(map[string]float64{"Баланс": balance})
}

func generateID() string {
	return strconv.Itoa(len(accounts))
}
