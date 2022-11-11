package consensus

import (
	"fmt"
	"math/big"
	"testing"

	"simple-blockchain/common"
	"simple-blockchain/core/types"
)

func SampleHeader() *types.Header {
	header := types.Header{
		ParentHash: common.BytesToHash([]byte("Parent")),
		TxHash:     common.BytesToHash([]byte("txhash")),
		Difficulty: big.NewInt(20),
		Number:     big.NewInt(1),
		Time:       uint64(4),
		Nonce:      types.EncodeNonce(uint64(5)),
	}
	return &header
}

func TestSealHash(t *testing.T) {
	header := SampleHeader()
	fmt.Println(SealHash(header))
}
