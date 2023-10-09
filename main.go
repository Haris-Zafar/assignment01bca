//M Haris Zafar
//20i-1885
//Assignment # 1
package main

import (
	"fmt"
	a1 "github.com/Haris-Zafar/assignment01bca"
)

func main() {
	genesisBlock := a1.NewBlock("Genesis Block", 0, "")

	block1 := a1.NewBlock("Alice to Bob", 12345, genesisBlock.CurrentHash)
	block2 := a1.NewBlock("Bob to Carol", 67890, block1.CurrentHash)

	genesisBlock.NextBlock = block1
	block1.NextBlock = block2

	a1.DisplayBlocks(genesisBlock)

	a1.ChangeBlock(block1, "Modified Transaction in Block 1")

	isValid := a1.VerifyChain(genesisBlock)
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}
}
