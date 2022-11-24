package types

import (
	"fmt"
	"math/big"
	"simple-blockchain/common"
	"testing"
)

func sampleHeader() *Header {
	header := Header{
		ParentHash: common.BytesToHash([]byte("Parent")),
		TxHash:     common.BytesToHash([]byte("txhash")),
		Difficulty: big.NewInt(20),
		Number:     big.NewInt(1),
		Time:       uint64(4),
		Nonce:      EncodeNonce(uint64(5)),
	}
	return &header
}

func sampleTxs() []*Transaction {
	var txs []*Transaction
	txData1 := TxData{
		Nonce: 0,
		To:    &common.Address{},
		Value: big.NewInt(1),
		Data:  nil,
	}
	tx1 := NewTx(txData1)
	tx1.Hash()
	txs = append(txs, tx1)

	txData2 := TxData{
		Nonce: 1,
		To:    &common.Address{},
		Value: big.NewInt(999),
		Data:  []byte("cranberry"),
	}
	tx2 := NewTx(txData2)
	tx2.Hash()
	txs = append(txs, tx2)
	return txs
}

func TestGenBlock(t *testing.T) {
	header := sampleHeader()
	txs := sampleTxs()

	block := NewBlock(header, txs)
	fmt.Println(block)
}
