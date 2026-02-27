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
		fmt.Println("5. Filter transactions")
		fmt.Println("6. Generate report")
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
		case "5":
			handleFilter(reader, manager)
		case "6":
			handleReport(reader, manager)
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

func handleReport(reader *bufio.Reader, manager *model.FinanceManager) {
	fmt.Println("Report type:")
	fmt.Println("1. By category")
	fmt.Println("2. By month")
	fmt.Print("> ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var reporter model.Reporter

	switch choice {
	case "1":
		reporter = model.CategoryReport{}
	case "2":
		reporter = model.MonthlyReport{}
	default:
		fmt.Println("Invalid choice.")
		fmt.Println()
		return
	}

	transactions := manager.GetAllTransactions()
	report := model.GenerateReport(reporter, transactions)
	fmt.Println()
	fmt.Print(report)
	fmt.Println()
}

func handleFilter(reader *bufio.Reader, manager *model.FinanceManager) {
	fmt.Println("Filter by:")
	fmt.Println("1. Type")
	fmt.Println("2. Category")
	fmt.Println("3. Date range")
	fmt.Print("> ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var filtered []model.Transaction

	switch choice {
	case "1":
		fmt.Print("Enter type (income/expense): ")
		typeStr, _ := reader.ReadString('\n')
		typeStr = strings.TrimSpace(typeStr)

		tType, err := model.ParseTransactionType(typeStr)
		if err != nil {
			fmt.Printf("Error: %s\n\n", err)
			return
		}
		filtered = manager.GetTransactionsByType(tType)

	case "2":
		fmt.Print("Enter category: ")
		category, _ := reader.ReadString('\n')
		category = strings.TrimSpace(category)
		filtered = manager.GetTransactionsByCategory(category)

	case "3":
		fmt.Print("Enter start date (YYYY-MM-DD): ")
		from, _ := reader.ReadString('\n')
		from = strings.TrimSpace(from)

		fmt.Print("Enter end date (YYYY-MM-DD): ")
		to, _ := reader.ReadString('\n')
		to = strings.TrimSpace(to)

		filtered = manager.GetTransactionsInDateRange(from, to)

	default:
		fmt.Println("Invalid choice.")
		fmt.Println()
		return
	}

	if len(filtered) == 0 {
		fmt.Println("No transactions found.")
		fmt.Println()
		return
	}

	fmt.Printf("\nFound %d transaction(s), total: %.2f\n", len(filtered), model.SumTransactions(filtered))
	for _, t := range filtered {
		fmt.Printf("\n  #%d\n", t.ID)
		printTransaction(t)
	}
	fmt.Println()
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
