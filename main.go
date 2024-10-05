package main

import (
	"A01/assignment01bca"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create the genesis block
	genesisBlock := assignment01bca.NewBlock("Genesis Block", 0, "")
	fmt.Println("Genesis Block Created:", genesisBlock)

	// Main program loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Display menu options
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a new block")
		fmt.Println("2. List all blocks")
		fmt.Println("3. Change a block transaction")
		fmt.Println("4. Verify the blockchain")
		fmt.Println("5. Exit")

		// Get user input
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		// Perform the action based on user input
		switch choice {
		case "1":
			// Add a new block
			fmt.Print("Enter transaction details (e.g., 'Alice to Bob'): ")
			scanner.Scan()
			transaction := scanner.Text()
			nonce := len(assignment01bca.Blockchain) // You can use the block count as the nonce
			previousBlock := assignment01bca.Blockchain[len(assignment01bca.Blockchain)-1]
			newBlock := assignment01bca.NewBlock(transaction, nonce, previousBlock.Hash)
			fmt.Println("New Block Added:", newBlock)

		case "2":
			// List all blocks
			fmt.Println("Listing all blocks:")
			assignment01bca.ListBlocks()

		case "3":
			// Change a block transaction
			fmt.Print("Enter the block index you want to change: ")
			scanner.Scan()
			indexInput := scanner.Text()
			index, err := strconv.Atoi(indexInput)
			if err != nil || index < 0 || index >= len(assignment01bca.Blockchain) {
				fmt.Println("Invalid block index")
			} else {
				fmt.Print("Enter new transaction details (e.g., 'Eve to Charlie'): ")
				scanner.Scan()
				newTransaction := scanner.Text()
				assignment01bca.ChangeBlock(index, newTransaction)
				fmt.Println("Block transaction updated.")
				assignment01bca.ListBlocks()
			}

		case "4":
			// Verify the blockchain
			fmt.Println("Verifying blockchain:")
			assignment01bca.VerifyChain()

		case "5":
			// Exit the program
			fmt.Println("Exiting...")
			return

		default:
			// Handle invalid input
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
