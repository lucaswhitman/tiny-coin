package coin

import (
	"crypto/sha256"
	"strconv"
	"time"
)

const Size = 32

type Block struct {
	Index        int
	TimeStamp    time.Time
	Data         string
	PreviousHash [Size]byte
	Hash         [Size]byte
}

func NewBlock(index int, timeStamp time.Time, data string, previousHash [Size]byte) *Block {
	return &Block{index, timeStamp, data, previousHash, HashBlock(index, timeStamp, data, previousHash)}
}

func HashBlock(index int, timeStamp time.Time, data string, previousHash [Size]byte) [Size]byte {
	sha := sha256.Sum256([]byte(strconv.Itoa(index) +
		timeStamp.String() +
		data +
		string(previousHash[:])))
	return sha
}
