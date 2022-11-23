package encode

import (
	"fmt"
	"math/big"
	"simple-blockchain/common"
	"simple-blockchain/core/types"
	"testing"
)

func TestTxEncode(t *testing.T) {
	txData := types.TxData{
		Nonce: 1,
		To:    &common.Address{},
		Value: big.NewInt(999),
		Data:  []byte("cranberry"),
	}
	tx := types.NewTx(txData)
	encoded := TxEncode(tx)
	fmt.Println(encoded)
}
