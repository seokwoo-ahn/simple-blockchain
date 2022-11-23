package types

import (
	"fmt"
	"math/big"
	"simple-blockchain/common"
	"testing"
)

func TestGenTx(t *testing.T) {
	txData1 := TxData{
		Nonce: 0,
		To:    &common.Address{},
		Value: big.NewInt(1),
		Data:  nil,
	}
	tx1 := NewTx(txData1)
	tx1.Hash()
	fmt.Println(tx1)

	txData2 := TxData{
		Nonce: 1,
		To:    &common.Address{},
		Value: big.NewInt(999),
		Data:  []byte("cranberry"),
	}
	tx2 := NewTx(txData2)
	tx2.Hash()
	fmt.Println(tx2)
}
