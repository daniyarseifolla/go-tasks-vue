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
	manager := model.NewFinanceManager()

	fmt.Println("=== Personal Finance Tracker ===")
	fmt.Println()

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Add transaction")
		fmt.Println("2. Show all transactions")
		fmt.Println("3. Find transaction by ID")
		fmt.Println("4. Show balance")
		fmt.Println("0. Exit")
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			handleAddTransaction(reader, manager)
		case "2":
			handleShowAll(manager)
		case "3":
			handleFindByID(reader, manager)
		case "4":
			handleShowBalance(manager)
		case "0":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid input, try again.")
			fmt.Println()
		}
	}
}

func handleAddTransaction(reader *bufio.Reader, manager *model.FinanceManager) {
	t, err := inputTransaction(reader)
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
		return
	}

	t = manager.AddTransaction(t)
	fmt.Printf("\nTransaction #%d created:\n", t.ID)
	printTransaction(t)
	fmt.Println()
}

func handleShowAll(manager *model.FinanceManager) {
	transactions := manager.GetAllTransactions()
	if len(transactions) == 0 {
		fmt.Println("No transactions yet.")
		fmt.Println()
		return
	}

	fmt.Printf("All transactions (%d):\n", len(transactions))
	for _, t := range transactions {
		fmt.Printf("\n  #%d\n", t.ID)
		printTransaction(t)
	}
	fmt.Println()
}

func handleFindByID(reader *bufio.Reader, manager *model.FinanceManager) {
	fmt.Print("Enter transaction ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Printf("Error: invalid ID: %s\n\n", idStr)
		return
	}

	t, err := manager.GetTransactionByID(id)
	if err != nil {
		fmt.Printf("Error: %s\n\n", err)
		return
	}

	fmt.Printf("\nTransaction #%d:\n", t.ID)
	printTransaction(*t)
	fmt.Println()
}

func handleShowBalance(manager *model.FinanceManager) {
	balance := manager.CalculateBalance()
	fmt.Printf("Current balance: %.2f\n\n", balance)
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
