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

func (i Inventory) ViewItems() {
	if len(i.items) == 0 {
		fmt.Println("Inventory is empty!!")
		return
	}

	const COL_WIDTH int = 10

	fmt.Println("Inventory Items")
	fmt.Println("============================================")
	fmt.Println("| Item Name          | Quantity | Price    |")
	fmt.Println("--------------------------------------------")

	for _, item := range i.items {
		quantity_str := fmt.Sprintf("%d", item.quantity)
		price_str := fmt.Sprintf("%.2f", item.price)

		name := fmt.Sprintf(" %v%*s ", item.name, COL_WIDTH*2-2-len(item.name), "")
		quantity := fmt.Sprintf(" %v%*s ", quantity_str, COL_WIDTH-2-len(quantity_str), "")
		price := fmt.Sprintf(" %v%*s ", price_str, COL_WIDTH-2-len(price_str), "")
		fmt.Printf("|%v|%v|%v|\n", name, quantity, price)
	}

	fmt.Println("|                    |          |          |")
	fmt.Println("============================================")
}

func (i Inventory) UpdateItem() {
	if len(i.items) == 0 {
		fmt.Println("Inventory is empty!!")
		return
	}

	var itemName string
	fmt.Printf("Enter Item name:")
	fmt.Scan(&itemName)

	if _, ok := i.items[itemName]; ok {
		item := InventoryItem{}
		fmt.Printf("Enter item's new name:")
		fmt.Scan(&item.name)

		fmt.Printf("Enter item's new price:")
		fmt.Scan(&item.price)

		fmt.Printf("Enter item' new quantity:")
		fmt.Scan(&item.quantity)

		delete(i.items, itemName)
		i.AddItem(item)
		fmt.Println("Item updated!!")
	} else {
		fmt.Println("Item not found!!")
	}
}

func (i Inventory) DeleteItem() {
	if len(i.items) == 0 {
		fmt.Println("Inventory is empty!!")
		return
	}

	var itemName string
	fmt.Printf("Enter Item name:")
	fmt.Scan(&itemName)

	if _, ok := i.items[itemName]; ok {
		delete(i.items, itemName)
		fmt.Println("Item deleted!!")
	} else {
		fmt.Println("Item not found!!")
	}
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
		// Add item case
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

		// Update item case
		case 2:
			inventory.UpdateItem()

		// Delete item case
		case 3:
			inventory.DeleteItem()

		// View items case
		case 4:
			inventory.ViewItems()

		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println()
	}
}
