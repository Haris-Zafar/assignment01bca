// assignment01bca/blockchain.go

package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
	NextBlock    *Block
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		NextBlock:    nil,
	}
	block.CurrentHash = CalculateHash(fmt.Sprintf("%v%v%v", transaction, nonce, previousHash))
	return block
}

func DisplayBlocks(genesisBlock *Block) {
	block := genesisBlock
	for block != nil {
		fmt.Printf("Transaction: %s\n", block.Transaction)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("Current Hash: %s\n", block.CurrentHash)
		fmt.Println("-------------------------------")
		block = block.NextBlock
	}
}

func ChangeBlock(blockToChange *Block, newTransaction string) {
	if blockToChange != nil {
		blockToChange.Transaction = newTransaction
		blockToChange.CurrentHash = CalculateHash(fmt.Sprintf("%v%v%v", blockToChange.Transaction, blockToChange.Nonce, blockToChange.PreviousHash))
	}
}

func VerifyChain(genesisBlock *Block) bool {
	block := genesisBlock
	prevHash := ""
	for block != nil {
		if block.PreviousHash != prevHash {
			return false
		}
		prevHash = block.CurrentHash
		block = block.NextBlock
	}
	return true
}

func CalculateHash(stringToHash string) string {
	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
}
