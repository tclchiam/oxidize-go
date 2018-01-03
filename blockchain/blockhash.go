package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"math/big"

	"github.com/tclchiam/block_n_go/chainhash"
	"github.com/tclchiam/block_n_go/tx"
)

const (
	targetBits = 16
	hashLength = 256
)

var (
	target = big.NewInt(1).Lsh(big.NewInt(1), uint(hashLength-targetBits))
)

func CalculateBlockHash(header *BlockHeader, nonce int) chainhash.Hash {
	rawBlockContents := [][]byte{
		header.PreviousHash[:],
		hashTransactions(header.Transactions),
		intToHex(header.Timestamp),
		intToHex(int64(nonce)),
	}
	rawBlockData := bytes.Join(rawBlockContents, []byte(nil))
	return chainhash.CalculateHash(rawBlockData)
}

func BlockValid(block *Block) bool {
	hash := CalculateBlockHash(block.Header(), block.Nonce)

	return new(big.Int).SetBytes(hash.Slice()).Cmp(target) == -1
}

func BlockHashValid(hash chainhash.Hash) bool {
	return new(big.Int).SetBytes(hash.Slice()).Cmp(target) == -1
}

func intToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func hashTransactions(transactions []*tx.Transaction) []byte {
	var transactionHashes [][]byte

	for _, transaction := range transactions {
		transactionHashes = append(transactionHashes, transaction.ID[:])
	}

	transactionHash := sha256.Sum256(bytes.Join(transactionHashes, []byte{}))

	return transactionHash[:]
}
