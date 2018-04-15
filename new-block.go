package coin

import (
	"strconv"
	"time"
)

func NextBlock(lastBlock Block) *Block {
	index := lastBlock.Index + 1
	timestamp := time.Now()
	data := "Hey! I'm block " + strconv.Itoa(index)
	previousHash := lastBlock.Hash
	return NewBlock(index, timestamp, data, previousHash)
}
