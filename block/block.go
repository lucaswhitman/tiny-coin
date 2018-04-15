package block

import (
	"crypto/sha256"
	"strconv"
	"time"

	"github.com/lucaswhitman/tiny-coin/transaction"
)

const Size = 32

type Data struct {
	ProofOfWork  int
	Transactions []transaction.Tranaction
}

type Block struct {
	Index        int
	TimeStamp    time.Time
	Data         Data
	PreviousHash [Size]byte
	Hash         [Size]byte
}

func NewBlock(index int, timeStamp time.Time, data Data, previousHash [Size]byte) Block {
	return Block{index, timeStamp, data, previousHash, HashBlock(index, timeStamp, data, previousHash)}
}

func HashBlock(index int, timeStamp time.Time, data Data, previousHash [Size]byte) [Size]byte {
	sha := sha256.Sum256([]byte(strconv.Itoa(index) +
		timeStamp.String() +
		strconv.Itoa(data.ProofOfWork) +
		string(previousHash[:])))
	return sha
}

func CreateGenesisBlock() Block {
	// Manually construct a block with
	// index zero and arbitrary previous hash
	var zero [32]byte
	var transactions []transaction.Tranaction
	data := Data{0, transactions}
	return NewBlock(0, time.Now(), data, zero)
}
