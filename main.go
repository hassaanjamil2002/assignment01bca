package main

import (
	"A01/assignment01bca"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Create the genesis block with an empty transaction pool and difficulty of 2
	genesisBlock := assignment01bca.BlockchainInstance.NewBlock("", 2)
	fmt.Println("Genesis Block Created:", genesisBlock)

	// Main program loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// Display menu options
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add a new transaction")
		fmt.Println("2. Create a new block")
		fmt.Println("3. List all blocks")
		fmt.Println("4. Change a block transaction")
		fmt.Println("5. Verify the blockchain")
		fmt.Println("6. Exit")

		// Get user input
		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := scanner.Text()

		// Perform the action based on user input
		switch choice {
		case "1":
			// Add a new transaction
			fmt.Print("Enter sender's address: ")
			scanner.Scan()
			sender := scanner.Text()

			fmt.Print("Enter recipient's address: ")
			scanner.Scan()
			recipient := scanner.Text()

			fmt.Print("Enter the value of the transaction: ")
			scanner.Scan()
			valueStr := scanner.Text()
			value, err := strconv.ParseFloat(valueStr, 32)
			if err != nil {
				fmt.Println("Invalid value. Please enter a valid float number.")
				break
			}

			// Add the transaction to the transaction pool
			assignment01bca.BlockchainInstance.AddTransaction(sender, recipient, float32(value))
			fmt.Println("Transaction added to the pool.")

		case "2":
			// Create a new block with the current transaction pool and difficulty 2
			previousBlock := assignment01bca.BlockchainInstance.Chain[len(assignment01bca.BlockchainInstance.Chain)-1]
			newBlock := assignment01bca.BlockchainInstance.NewBlock(previousBlock.Hash, 2)
			fmt.Println("New Block Created:", newBlock)

		case "3":
			// List all blocks
			fmt.Println("Listing all blocks:")
			assignment01bca.BlockchainInstance.ListBlocks()

		case "4":
			// Change a block transaction
			fmt.Print("Enter the block index you want to change: ")
			scanner.Scan()
			blockIndexInput := scanner.Text()
			blockIndex, err := strconv.Atoi(blockIndexInput)
			if err != nil || blockIndex < 0 || blockIndex >= len(assignment01bca.BlockchainInstance.Chain) {
				fmt.Println("Invalid block index")
				break
			}

			fmt.Print("Enter the transaction index you want to change: ")
			scanner.Scan()
			transactionIndexInput := scanner.Text()
			transactionIndex, err := strconv.Atoi(transactionIndexInput)
			if err != nil || transactionIndex < 0 || transactionIndex >= len(assignment01bca.BlockchainInstance.Chain[blockIndex].Transactions) {
				fmt.Println("Invalid transaction index")
				break
			}

			// Get new sender, recipient, and value from the user
			fmt.Print("Enter the new sender address: ")
			scanner.Scan()
			newSender := scanner.Text()

			fmt.Print("Enter the new recipient address: ")
			scanner.Scan()
			newRecipient := scanner.Text()

			fmt.Print("Enter the new value of the transaction: ")
			scanner.Scan()
			newValueInput := scanner.Text()
			newValue, err := strconv.ParseFloat(newValueInput, 32)
			if err != nil {
				fmt.Println("Invalid value. Please enter a valid float number.")
				break
			}

			// Call the ChangeBlock function with the new values
			assignment01bca.BlockchainInstance.ChangeBlock(blockIndex, transactionIndex, newSender, newRecipient, float32(newValue))
			assignment01bca.BlockchainInstance.ListBlocks()

		case "5":
			// Verify the blockchain
			fmt.Println("Verifying blockchain:")
			assignment01bca.BlockchainInstance.VerifyChain()

		case "6":
			// Exit the program
			fmt.Println("Exiting...")
			return

		default:
			// Handle invalid input
			fmt.Println("Invalid choice. Please choose a valid option.")
		}
	}
}
