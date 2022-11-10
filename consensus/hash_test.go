package consensus

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SampleHeader() *types.Header {
	header := types.Header{
		ParentHash:  common.BytesToHash([]byte("Parent")),
		UncleHash:   common.BytesToHash([]byte("Uncle")),
		Coinbase:    common.BytesToAddress([]byte("Address")),
		Root:        common.BytesToHash([]byte("Root")),
		TxHash:      common.BytesToHash([]byte("txhash")),
		ReceiptHash: common.BytesToHash([]byte("receipt")),
		Bloom:       types.BytesToBloom([]byte("bloom")),
		Difficulty:  big.NewInt(1),
		Number:      big.NewInt(1),
		GasLimit:    uint64(2),
		GasUsed:     uint64(3),
		Time:        uint64(4),
		Extra:       []byte("extra"),
		MixDigest:   common.BytesToHash([]byte("mixdigest")),
		Nonce:       types.EncodeNonce(uint64(5)),
	}
	return &header
}

func TestSealHash(t *testing.T) {
	header := SampleHeader()
	fmt.Println(SealHash(header))
}
