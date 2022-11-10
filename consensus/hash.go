package consensus

import (
	"encoding/binary"
	"math/big"
	"simple-blockchain/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := crypto.NewKeccakState()

	enc := []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
	}
	if header.BaseFee != nil {
		enc = append(enc, header.BaseFee)
	}

	val := EncodeBlock(enc)

	hasher.Write(val)
	hasher.Sum(hash[:0])
	return hash
}

func EncodeBlock(input []interface{}) []byte {
	var val []byte
	for i, v := range input {
		var b []byte
		if i < 2 {
			b = v.(common.Hash).Bytes()
		} else if i < 3 {
			b = v.(common.Address).Bytes()
		} else if i < 6 {
			b = v.(common.Hash).Bytes()
		} else if i < 7 {
			b = v.(types.Bloom).Bytes()
		} else if i < 9 {
			b = v.(*big.Int).Bytes()
		} else if i < 12 {
			b = make([]byte, 8)
			binary.LittleEndian.PutUint64(b, v.(uint64))
		} else if i < 13 {
			b = v.([]byte)
		} else if i < 14 {
			b = v.(common.Hash).Bytes()
		} else if i < 15 {
			b = make([]byte, 8)
			binary.LittleEndian.PutUint64(b, v.(types.BlockNonce).Uint64())
		}
		val = append(val, b...)
	}
	return val
}
