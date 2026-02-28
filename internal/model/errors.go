package model

import (
	"errors"
	"fmt"
)

var (
	ErrNegativeAmount         = errors.New("amount cannot be negative")
	ErrEmptyCategory          = errors.New("category cannot be empty")
	ErrInvalidTransactionType = errors.New("type must be 'income' or 'expense'")
)

type ErrTransactionNotFound struct {
	ID int
}

func (e *ErrTransactionNotFound) Error() string {
	return fmt.Sprintf("transaction with ID %d not found", e.ID)
}
