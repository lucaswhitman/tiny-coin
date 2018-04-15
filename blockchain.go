package coin

import "fmt"

func main() {
	// Create the blockchain and add the genesis block
	blockchain := make([]Block, 3)
	blockchain = append(blockchain, CreateGenesisBlock())
	previousBlock = blockchain[0]

	// How many blocks should we add to the chain
	// after the genesis block
	blocksToAdd = 20

	// Add blocks to the chain
	for i := 0; i < blocksToAdd; i++ {
		blockToAdd = NextBlock(previousBlock)
		blockchain = append(blockchain, blockToAdd)
		previousBlock = blockToAdd
		// Tell everyone about it!
		fmt.Println("Block #{} has been added to the blockchain!".format(blockToAdd.index))
		fmt.Println("Hash: {}\n".format(blockToAdd.hash))
	}
}
