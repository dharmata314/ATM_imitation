package entities

import "fmt"

func (a *Account) Deposit(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.Balance < amount {
		return fmt.Errorf("Недостаточно средств")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.Balance
}
