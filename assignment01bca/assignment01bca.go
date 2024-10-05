package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// Block structure
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

// Blockchain slice to hold blocks
var Blockchain []Block

// NewBlock creates a new block with a transaction, nonce, and previous hash
func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{Transaction: transaction, Nonce: nonce, PreviousHash: previousHash}
	block.Hash = CalculateHash(block.Transaction + strconv.Itoa(block.Nonce) + block.PreviousHash)
	Blockchain = append(Blockchain, *block)
	return block
}

// ListBlocks prints all blocks in the blockchain
func ListBlocks() {
	fmt.Println("\n--- Blockchain ---")
	for i, block := range Blockchain {
		fmt.Printf("\nBlock %d\n", i)
		fmt.Println(strings.Repeat("=", 40))
		fmt.Printf("Transaction  : %s\n", block.Transaction)
		fmt.Printf("Nonce        : %d\n", block.Nonce)
		fmt.Printf("PreviousHash : %s\n", block.PreviousHash)
		fmt.Printf("Hash         : %s\n", block.Hash)
		fmt.Println(strings.Repeat("=", 40))
	}
	fmt.Println("\n--- End of Blockchain ---\n")
}

// ChangeBlock modifies the transaction of a block at a given index
func ChangeBlock(index int, newTransaction string) {
	if index >= 0 && index < len(Blockchain) {
		Blockchain[index].Transaction = newTransaction
		Blockchain[index].Hash = CalculateHash(Blockchain[index].Transaction + strconv.Itoa(Blockchain[index].Nonce) + Blockchain[index].PreviousHash)
	} else {
		fmt.Println("Invalid block index")
	}
}

// VerifyChain checks the integrity of the blockchain
func VerifyChain() bool {
	for i := 1; i < len(Blockchain); i++ {
		if Blockchain[i].PreviousHash != Blockchain[i-1].Hash {
			fmt.Println("Blockchain is invalid!")
			return false
		}
	}
	fmt.Println("Blockchain is valid")
	return true
}

// CalculateHash calculates the hash for a given string
func CalculateHash(stringToHash string) string {
	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	return hex.EncodeToString(hash.Sum(nil))
}
