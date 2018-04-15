package main

import "fmt"
import "github.com/lucaswhitman/tiny-coin/coin"

func main() {
	// Create the blockchain and add the genesis block
	var blockchain = make([]coin.Block, 3)
	blockchain = append(blockchain, coin.CreateGenesisBlock())
	var previousBlock = blockchain[0]

	// How many blocks should we add to the chain
	// after the genesis block
	var blocksToAdd = 20

	// Add blocks to the chain
	for i := 0; i < blocksToAdd; i++ {
		var blockToAdd = coin.NextBlock(previousBlock)
		blockchain = append(blockchain, blockToAdd)
		previousBlock = blockToAdd
		// Tell everyone about it!
		fmt.Printf("Block %d has been added to the blockchain!\n", blockToAdd.Index)
		fmt.Printf("Hash: %d\n", blockToAdd.Hash)
	}
}
