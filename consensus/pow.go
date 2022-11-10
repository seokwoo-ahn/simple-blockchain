package consensus

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"golang.org/x/crypto/sha3"
)

// // two256 is a big integer representing 2^256
// var two256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

// // mine is the actual proof-of-work miner that searches for a nonce starting from
// // seed that results in correct final block difficulty.
// func mine(block *types.Block, id int, seed uint64, abort chan struct{}, found chan *types.Block) {
// 	// Extract some data from the header
// 	var (
// 		header  = block.Header()
// 		hash    = ethash.SealHash(header).Bytes()
// 		target  = new(big.Int).Div(two256, header.Difficulty)
// 		number  = header.Number.Uint64()
// 		dataset = ethash.dataset(number, false)
// 	)
// 	// Start generating random nonces until we abort or find a good one
// 	var (
// 		attempts  = int64(0)
// 		nonce     = seed
// 		powBuffer = new(big.Int)
// 	)
// 	logger := ethash.config.Log.New("miner", id)
// 	logger.Trace("Started ethash search for new nonces", "seed", seed)
// search:
// 	for {
// 		// We don't have to update hash rate on every nonce, so update after after 2^X nonces
// 		attempts++
// 		// Compute the PoW value of this nonce
// 		digest, result := hashimotoFull(dataset.dataset, hash, nonce)
// 		if powBuffer.SetBytes(result).Cmp(target) <= 0 {
// 			// Correct nonce found, create a new header with it
// 			header = types.CopyHeader(header)
// 			header.Nonce = types.EncodeNonce(nonce)
// 			header.MixDigest = common.BytesToHash(digest)

// 			// Seal and return a block (if still needed)
// 			select {
// 			case found <- block.WithSeal(header):
// 				logger.Trace("Ethash nonce found and reported", "attempts", nonce-seed, "nonce", nonce)
// 			case <-abort:
// 				logger.Trace("Ethash nonce found but discarded", "attempts", nonce-seed, "nonce", nonce)
// 			}
// 			break search
// 		}
// 		nonce++

// 	}
// 	// Datasets are unmapped in a finalizer. Ensure that the dataset stays live
// 	// during sealing so it's not unmapped while being read.
// 	runtime.KeepAlive(dataset)
// }

// func hashimotoFull(dataset []uint32, hash []byte, nonce uint64) ([]byte, []byte) {
// 	lookup := func(index uint32) []uint32 {
// 		offset := index * hashWords
// 		return dataset[offset : offset+hashWords]
// 	}
// 	return hashimoto(hash, nonce, uint64(len(dataset))*4, lookup)
// }

// // hashimoto aggregates data from the full dataset in order to produce our final
// // value for a particular header hash and nonce.
// func hashimoto(hash []byte, nonce uint64, size uint64, lookup func(index uint32) []uint32) ([]byte, []byte) {
// 	// Calculate the number of theoretical rows (we use one buffer nonetheless)
// 	rows := uint32(size / mixBytes)

// 	// Combine header+nonce into a 64 byte seed
// 	seed := make([]byte, 40)
// 	copy(seed, hash)
// 	binary.LittleEndian.PutUint64(seed[32:], nonce)

// 	seed = crypto.Keccak512(seed)
// 	seedHead := binary.LittleEndian.Uint32(seed)

// 	// Start the mix with replicated seed
// 	mix := make([]uint32, mixBytes/4)
// 	for i := 0; i < len(mix); i++ {
// 		mix[i] = binary.LittleEndian.Uint32(seed[i%16*4:])
// 	}
// 	// Mix in random dataset nodes
// 	temp := make([]uint32, len(mix))

// 	for i := 0; i < loopAccesses; i++ {
// 		parent := fnv(uint32(i)^seedHead, mix[i%len(mix)]) % rows
// 		for j := uint32(0); j < mixBytes/hashBytes; j++ {
// 			copy(temp[j*hashWords:], lookup(2*parent+j))
// 		}
// 		fnvHash(mix, temp)
// 	}
// 	// Compress mix
// 	for i := 0; i < len(mix); i += 4 {
// 		mix[i/4] = fnv(fnv(fnv(mix[i], mix[i+1]), mix[i+2]), mix[i+3])
// 	}
// 	mix = mix[:len(mix)/4]

// 	digest := make([]byte, common.HashLength)
// 	for i, val := range mix {
// 		binary.LittleEndian.PutUint32(digest[i*4:], val)
// 	}
// 	return digest, crypto.Keccak256(append(seed, digest...))
// }

func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

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
