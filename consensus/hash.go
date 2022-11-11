package consensus

import (
	"encoding/binary"
	"math/big"
	"simple-blockchain/core/types"
	"simple-blockchain/crypto"

	"simple-blockchain/common"
)

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := crypto.NewKeccakState()

	enc := []interface{}{
		header.ParentHash,
		header.TxHash,
		header.Difficulty,
		header.Number,
		header.Time,
	}

	val := EncodeBlockHeader(enc)

	hasher.Write(val)
	hasher.Sum(hash[:0])
	return hash
}

func EncodeBlockHeader(input []interface{}) []byte {
	var val []byte
	for i, v := range input {
		var b []byte
		if i < 2 {
			b = v.(common.Hash).Bytes()
		} else if i < 4 {
			b = v.(*big.Int).Bytes()
		} else if i < 5 {
			b = make([]byte, 8)
			binary.LittleEndian.PutUint64(b, v.(uint64))
		} else if i < 6 {
			b = make([]byte, 8)
			binary.LittleEndian.PutUint64(b, v.(types.BlockNonce).Uint64())
		}
		val = append(val, b...)
	}
	return val
}
