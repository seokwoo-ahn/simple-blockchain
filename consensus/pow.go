package consensus

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"simple-blockchain/crypto"

	"simple-blockchain/core/types"
)

// two256 is a big integer representing 2^256
var two256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))

// mine is the actual proof-of-work miner that searches for a nonce starting from
// seed that results in correct final block difficulty.
func Mine(block *types.Block, seed uint64, abort chan struct{}, found chan *types.Block) {
	// Extract some data from the header
	var (
		header = block.Header()
		hash   = SealHash(header).Bytes()
		target = new(big.Int).Div(two256, header.Difficulty)
	)
	// Start generating random nonces until we abort or find a good one
	var (
		attempts  = int64(0)
		nonce     = seed
		powBuffer = new(big.Int)
	)
search:
	for {
		attempts++
		// Compute the PoW value of this nonce
		result := hashimoto(hash, nonce)
		if powBuffer.SetBytes(result).Cmp(target) <= 0 {
			// Correct nonce found, create a new header with it
			header = types.CopyHeader(header)
			header.Nonce = types.EncodeNonce(nonce)

			// Seal and return a block (if still needed)
			select {
			case found <- block.WithSeal(header):
				fmt.Println("Ethash nonce found and reported", "attempts", attempts, "nonce", nonce)
			case <-abort:
				fmt.Println("Ethash nonce found but discarded", "attempts", attempts, "nonce", nonce)
			}
			break search
		}
		nonce++
	}
}

// hashimoto aggregates data from the full dataset in order to produce our final
// value for a particular header hash and nonce.
func hashimoto(hash []byte, nonce uint64) []byte {
	// Combine header+nonce into a 64 byte seed
	seed := make([]byte, 40)
	copy(seed, hash)
	binary.LittleEndian.PutUint64(seed[32:], nonce)
	seed = crypto.Keccak512(seed)

	return crypto.Keccak256(seed)
}
