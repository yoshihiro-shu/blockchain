package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	timestamp    int64
	nonce        int
	prevHash     [32]byte
	transactions []*Transaction
}

func NewBlock(nonce int, prevHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		prevHash:     prevHash,
		timestamp:    time.Now().Unix(),
		transactions: transactions,
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		TimeStamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PrevHash     [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		TimeStamp:    b.timestamp,
		Nonce:        b.nonce,
		PrevHash:     b.prevHash,
		Transactions: b.transactions,
	})
}

func (b *Block) Print() {
	fmt.Printf("timestamp        %d\n", b.timestamp)
	fmt.Printf("nonce            %d\n", b.nonce)
	fmt.Printf("previous_hash    %s\n", b.prevHash)
	for _, t := range b.transactions {
		t.Print()
	}
}
