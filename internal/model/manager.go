package model

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	return nil, &ErrTransactionNotFound{ID: id}
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

func (fm *FinanceManager) GetTransactionsByType(t TransactionType) []Transaction {
	var result []Transaction
	for _, tr := range fm.transactions {
		if tr.Type == t {
			result = append(result, tr)
		}
	}
	return result
}

func (fm *FinanceManager) GetTransactionsByCategory(category string) []Transaction {
	var result []Transaction
	for _, t := range fm.transactions {
		if t.Category == category {
			result = append(result, t)
		}
	}
	return result
}

func (fm *FinanceManager) GetTransactionsInDateRange(from, to string) []Transaction {
	var result []Transaction
	for _, t := range fm.transactions {
		if t.Date >= from && t.Date <= to {
			result = append(result, t)
		}
	}
	return result
}

func (fm *FinanceManager) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(fm.transactions, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal transactions: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (fm *FinanceManager) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var transactions []Transaction
	if err := json.Unmarshal(data, &transactions); err != nil {
		return fmt.Errorf("failed to parse transactions: %w", err)
	}

	fm.transactions = transactions

	maxID := 0
	for _, t := range fm.transactions {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	fm.nextID = maxID + 1

	return nil
}

func SumTransactions(transactions []Transaction) float64 {
	var total float64
	for _, t := range transactions {
		total += t.Amount
	}
	return total
}
