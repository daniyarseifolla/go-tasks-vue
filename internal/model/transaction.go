package model

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

func ParseTransactionType(s string) (TransactionType, error) {
	switch s {
	case string(Income):
		return Income, nil
	case string(Expense):
		return Expense, nil
	default:
		return "", errors.New("type must be 'income' or 'expense'")
	}
}

type Transaction struct {
	ID       int
	Amount   float64
	Category string
	Date     string
	Type     TransactionType
}

func NewTransaction(amount float64, category string, t TransactionType) (Transaction, error) {
	if amount < 0 {
		return Transaction{}, errors.New("amount cannot be negative")
	}

	if category == "" {
		return Transaction{}, errors.New("category cannot be empty")
	}

	return Transaction{
		Amount:   amount,
		Category: category,
		Date:     time.Now().Format("2006-01-02"),
		Type:     t,
	}, nil
}
