package main

import (
	"fmt"
)

type InventoryItem struct {
	name     string
	quantity int
	price    float32
}

type Inventory struct {
	items map[string]InventoryItem
}

func (i Inventory) AddItem(item InventoryItem) {
	i.items[item.name] = item
}

func main() {
	inventory := Inventory{make(map[string]InventoryItem)}

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

			item.name = "Toothpaste"
			item.quantity = 20
			item.price = 55
			// fmt.Printf("Enter Item name:")
			// fmt.Scan(&item.name)

			// fmt.Printf("Enter Item price:")
			// fmt.Scan(&item.price)

			// fmt.Printf("Enter Item quantity:")
			// fmt.Scan(&item.quantity)

			// fmt.Println("item final", item)

			inventory.AddItem(item)

		case 2:
			fmt.Println("You chose Operation 2")

		case 3:
			fmt.Println("You chose Operation 3")

		case 4:
			fmt.Println("Inventory Items")
			fmt.Println("--------------------------------------------")
			fmt.Println("|     Item Name |     Quantity |     Price |")

			for _, item := range inventory.items {
				fmt.Printf("|     %v |     %v |     %v |", item.name, item.quantity, item.price)
			}

			fmt.Println()

		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println()
	}
}
