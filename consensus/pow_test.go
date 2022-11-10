package consensus

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func TestSealHash(t *testing.T) {
	var header types.Header

	header.ParentHash = common.BytesToHash([]byte("Parent"))
	header.UncleHash = common.BytesToHash([]byte("Uncle"))
	header.Coinbase = common.BytesToAddress([]byte("Address"))
	header.Root = common.BytesToHash([]byte("Root"))
	header.TxHash = common.BytesToHash([]byte("txhash"))
	header.ReceiptHash = common.BytesToHash([]byte("receipt"))
	header.Bloom = types.BytesToBloom([]byte("bloom"))
	header.Difficulty = big.NewInt(0)
	header.Number = big.NewInt(1)
	header.GasLimit = uint64(2)
	header.GasUsed = uint64(3)
	header.Time = uint64(4)
	header.Extra = []byte("extra")
	header.MixDigest = common.BytesToHash([]byte("mixdigest"))
	header.Nonce = types.EncodeNonce(uint64(5))

	fmt.Println(SealHash(&header))

}
