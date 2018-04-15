package coin

import "time"

func CreateGenesisBlock() *Block {
	// Manually construct a block with
	// index zero and arbitrary previous hash
	var zero [32]byte
	return NewBlock(0, time.Now(), "Genesis Block", zero)
}
