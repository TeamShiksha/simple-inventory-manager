package main

import (
	"fmt"
)

type InventoryItem struct {
	name     string
	quantity int
	price    float32
}

func main() {
	OPERATIONS := []string{
		"Add item",
		"Update Item",
		"Delete Item",
		"Show all items",
		"Exit",
	}

	fmt.Println("INVENTORY MANAGEMENT SYSTEM")
	for {
		for i, step := range OPERATIONS {
			fmt.Printf("%d. %v\n", i+1, step)
		}

		fmt.Printf("Enter your choice:")

		// Read user input
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		fmt.Println("Your choice:", choice)

		switch choice {
		case 1:
			item := InventoryItem{}
			fmt.Println("item", item)
			fmt.Printf("Enter Item name:")
			fmt.Scan(&item.name)

			fmt.Printf("Enter Item price:")
			fmt.Scan(&item.price)

			fmt.Printf("Enter Item quantity:")
			fmt.Scan(&item.quantity)

			fmt.Println("item final", item)
		case 2:
			fmt.Println("You chose Operation 2")
		case 3:
			fmt.Println("You chose Operation 3")
		case 4:
			fmt.Println("You chose Operation 4")
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println()
	}
}
