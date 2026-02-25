package model

import "fmt"

type FinanceManager struct {
	transactions []Transaction
	nextID       int
}

func NewFinanceManager() *FinanceManager {
	return &FinanceManager{
		transactions: []Transaction{},
		nextID:       1,
	}
}

func (fm *FinanceManager) AddTransaction(t Transaction) Transaction {
	t.ID = fm.nextID
	fm.nextID++
	fm.transactions = append(fm.transactions, t)
	return t
}

func (fm *FinanceManager) GetAllTransactions() []Transaction {
	return fm.transactions
}

func (fm *FinanceManager) GetTransactionByID(id int) (*Transaction, error) {
	for i := range fm.transactions {
		if fm.transactions[i].ID == id {
			return &fm.transactions[i], nil
		}
	}
	return nil, fmt.Errorf("transaction with ID %d not found", id)
}

func (fm *FinanceManager) CalculateBalance() float64 {
	var balance float64
	for _, t := range fm.transactions {
		switch t.Type {
		case Income:
			balance += t.Amount
		case Expense:
			balance -= t.Amount
		}
	}
	return balance
}
