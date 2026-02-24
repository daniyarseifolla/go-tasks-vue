package cli

import (
	"bufio"
	"finance-tracker/internal/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Personal Finance Tracker ===")
	fmt.Println()

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Add transaction")
		fmt.Println("0. Exit")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			t, err := inputTransaction(reader)
			if err != nil {
				fmt.Printf("Error: %s\n\n", err)
				continue
			}
			fmt.Println()
			fmt.Println("Transaction created:")
			printTransaction(t)
			fmt.Println()
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid input, try again.")
			fmt.Println()
		}
	}
}

func inputTransaction(reader *bufio.Reader) (model.Transaction, error) {
	fmt.Print("Enter type (income/expense): ")
	typeStr, _ := reader.ReadString('\n')
	typeStr = strings.TrimSpace(typeStr)

	tType, err := model.ParseTransactionType(typeStr)
	if err != nil {
		return model.Transaction{}, err
	}

	fmt.Print("Enter amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("invalid amount: %s", amountStr)
	}

	fmt.Print("Enter category: ")
	category, _ := reader.ReadString('\n')
	category = strings.TrimSpace(category)

	return model.NewTransaction(amount, category, tType)
}

func printTransaction(t model.Transaction) {
	typeLabel := "Income"
	if t.Type == model.Expense {
		typeLabel = "Expense"
	}
	fmt.Printf("  Type:     %s\n", typeLabel)
	fmt.Printf("  Amount:   %.2f\n", t.Amount)
	fmt.Printf("  Category: %s\n", t.Category)
	fmt.Printf("  Date:     %s\n", t.Date)
}
