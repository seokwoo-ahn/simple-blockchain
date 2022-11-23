package types

import (
	"fmt"
	"math/big"
	"simple-blockchain/common"
	"testing"
)

func TestGenTx(t *testing.T) {
	txData := TxData{
		Nonce: 0,
		To:    &common.Address{},
		Value: big.NewInt(1),
		Data:  nil,
	}
	tx := NewTx(txData)
	fmt.Println(tx)
}
