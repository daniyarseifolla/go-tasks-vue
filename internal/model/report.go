package model

import (
	"fmt"
	"sort"
	"strings"
)

type Reporter interface {
	GenerateReport(transactions []Transaction) string
}

func GenerateReport(r Reporter, transactions []Transaction) string {
	return r.GenerateReport(transactions)
}

type CategoryReport struct{}

func (cr CategoryReport) GenerateReport(transactions []Transaction) string {
	categories := make(map[string]float64)

	for _, t := range transactions {
		switch t.Type {
		case Income:
			categories[t.Category] += t.Amount
		case Expense:
			categories[t.Category] -= t.Amount
		}
	}

	if len(categories) == 0 {
		return "No data for report."
	}

	keys := make([]string, 0, len(categories))
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString("Category Report:\n")
	for _, cat := range keys {
		sb.WriteString(fmt.Sprintf("  %-15s %+.2f\n", cat, categories[cat]))
	}

	return sb.String()
}

type MonthlyReport struct{}

func (mr MonthlyReport) GenerateReport(transactions []Transaction) string {
	type monthData struct {
		income  float64
		expense float64
	}

	months := make(map[string]*monthData)

	for _, t := range transactions {
		month := t.Date[:7] // "YYYY-MM"

		if _, ok := months[month]; !ok {
			months[month] = &monthData{}
		}

		switch t.Type {
		case Income:
			months[month].income += t.Amount
		case Expense:
			months[month].expense += t.Amount
		}
	}

	if len(months) == 0 {
		return "No data for report."
	}

	keys := make([]string, 0, len(months))
	for k := range months {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var sb strings.Builder
	sb.WriteString("Monthly Report:\n")
	for _, month := range keys {
		d := months[month]
		net := d.income - d.expense
		sb.WriteString(fmt.Sprintf("  %s  Income: %.2f  Expense: %.2f  Net: %+.2f\n",
			month, d.income, d.expense, net))
	}

	return sb.String()
}
