package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Transaction struct to store transaction details
type Transaction struct {
	TransactionID              string
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

// Block structure with transactions and nonce
type Block struct {
	Transactions []*Transaction
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain struct with chain and transaction pool
type Blockchain struct {
	Chain           []*Block
	TransactionPool []*Transaction
}

// Global blockchain instance
var BlockchainInstance = &Blockchain{Chain: []*Block{}, TransactionPool: []*Transaction{}}

// NewTransaction creates a new transaction
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	transaction := &Transaction{
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
	// Calculate transaction ID as hash
	transaction.TransactionID = CalculateHash(sender + recipient + fmt.Sprintf("%f", value))
	return transaction
}

// AddTransaction adds a transaction to the transaction pool
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	transaction := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, transaction)
}

// NewBlock creates a new block with transactions from the transaction pool and a derived nonce
func (bc *Blockchain) NewBlock(previousHash string, difficulty int) *Block {
	// Derive nonce with Proof-of-Work
	nonce := ProofOfWork(bc.TransactionPool, previousHash, difficulty)

	// Create new block
	block := &Block{
		Transactions: bc.TransactionPool,
		Nonce:        nonce,
		PreviousHash: previousHash,
	}
	block.Hash = CalculateHash(strconv.Itoa(nonce) + previousHash)

	// Add block to chain
	bc.Chain = append(bc.Chain, block)

	// Reset the transaction pool after creating a block
	bc.TransactionPool = []*Transaction{}

	return block
}

// ListBlocks prints all blocks in the blockchain along with their transactions in JSON format
func (bc *Blockchain) ListBlocks() {
	fmt.Println("\n--- Blockchain ---")
	for i, block := range bc.Chain {
		fmt.Printf("\nBlock %d\n", i)
		fmt.Println(strings.Repeat("=", 40))
		fmt.Printf("Nonce        : %d\n", block.Nonce)
		fmt.Printf("PreviousHash : %s\n", block.PreviousHash)
		fmt.Printf("Hash         : %s\n", block.Hash)

		// Display transactions in JSON format
		transactionsJSON, err := json.MarshalIndent(block.Transactions, "", "  ")
		if err != nil {
			fmt.Println("Error displaying transactions:", err)
		} else {
			fmt.Println("Transactions : ", string(transactionsJSON))
		}
		fmt.Println(strings.Repeat("=", 40))
	}
	fmt.Println("\n--- End of Blockchain ---\n")
}

// ChangeBlock modifies the transactions of a block at a given index
// ChangeBlock allows the user to modify a specific transaction in a given block by index
func (bc *Blockchain) ChangeBlock(blockIndex int, transactionIndex int, newSender string, newRecipient string, newValue float32) {
	// Check if the block index is valid
	if blockIndex >= 0 && blockIndex < len(bc.Chain) {
		block := bc.Chain[blockIndex]

		// Check if the transaction index is valid
		if transactionIndex >= 0 && transactionIndex < len(block.Transactions) {
			// Modify the transaction
			block.Transactions[transactionIndex].SenderBlockchainAddress = newSender
			block.Transactions[transactionIndex].RecipientBlockchainAddress = newRecipient
			block.Transactions[transactionIndex].Value = newValue

			// Recalculate the transaction ID and block hash after modification
			block.Transactions[transactionIndex].TransactionID = CalculateHash(newSender + newRecipient + fmt.Sprintf("%f", newValue))
			block.Hash = CalculateHash(strconv.Itoa(block.Nonce) + block.PreviousHash + bc.transactionData(block.Transactions))

			fmt.Println("Transaction successfully updated!")
		} else {
			fmt.Println("Invalid transaction index")
		}
	} else {
		fmt.Println("Invalid block index")
	}
}

// Helper function to convert transaction data to a string for hash calculation
func (bc *Blockchain) transactionData(transactions []*Transaction) string {
	var transactionData strings.Builder
	for _, tx := range transactions {
		transactionData.WriteString(tx.SenderBlockchainAddress + tx.RecipientBlockchainAddress + fmt.Sprintf("%f", tx.Value))
	}
	return transactionData.String()
}

// VerifyChain checks the integrity of the blockchain
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Chain); i++ {
		if bc.Chain[i].PreviousHash != bc.Chain[i-1].Hash {
			fmt.Println("Blockchain is invalid!")
			return false
		}
	}
	fmt.Println("Blockchain is valid")
	return true
}

// ProofOfWork derives a nonce by finding a hash with a specific difficulty level
func ProofOfWork(transactions []*Transaction, previousHash string, difficulty int) int {
	nonce := 0
	for {
		hash := CalculateHash(strconv.Itoa(nonce) + previousHash)
		// Check if hash satisfies the difficulty (e.g., starts with "00" for difficulty 2)
		if strings.HasPrefix(hash, strings.Repeat("0", difficulty)) {
			break
		}
		nonce++
	}
	return nonce
}

// CalculateHash calculates the hash for a given string
func CalculateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
