package types

import (
	"encoding/binary"
	"math/big"

	"sync/atomic"
	"time"

	"simple-blockchain/common"
	// "simple-blockchain/consensus"
)

var (
	EmptyRootHash = common.BytesToHash([]byte("simple-blockchain"))
)

// A BlockNonce is a 64-bit hash which proves (combined with the
// mix-hash) that a sufficient amount of computation has been carried
// out on a block.
type BlockNonce [8]byte

// Uint64 returns the integer value of a block nonce.
func (n BlockNonce) Uint64() uint64 {
	return binary.BigEndian.Uint64(n[:])
}

// Header represents a block header in the Ethereum blockchain.
type Header struct {
	ParentHash common.Hash `json:"parentHash"       gencodec:"required"`
	TxHash     common.Hash `json:"transactionsRoot" gencodec:"required"`
	Difficulty *big.Int    `json:"difficulty"       gencodec:"required"`
	Number     *big.Int    `json:"number"           gencodec:"required"`
	Time       uint64      `json:"timestamp"        gencodec:"required"`
	Nonce      BlockNonce  `json:"nonce"`
}

// Block represents an entire block in the Ethereum blockchain.
type Block struct {
	header       *Header
	transactions Transactions

	// caches
	hash atomic.Value
	// size atomic.Value

	// Td is used by package core to store the total difficulty
	// of the chain up to and including the block.
	td *big.Int

	// These fields are used by package eth to track
	// inter-peer block relay.
	ReceivedAt   time.Time
	ReceivedFrom interface{}
}

// NewBlock creates a new block. The input data is copied,
// changes to header and to the field values will not affect the
// block.
//
// The values of TxHash, UncleHash, ReceiptHash and Bloom in header
// are ignored and set to values derived from the given txs, uncles
// and receipts.
func NewBlock(header *Header, txs Transactions) *Block {
	b := &Block{header: header, td: new(big.Int)}

	if len(txs) == 0 {
		b.header.TxHash = EmptyRootHash
	} else {
		b.header.TxHash = txs.Hash()
		b.transactions = make(Transactions, len(txs))
		copy(b.transactions, txs)
	}

	return b
}

func (b *Block) Header() *Header            { return b.header }
func (b *Block) Transactions() Transactions { return b.transactions }

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
// TODO
func (h *Header) Hash() common.Hash {
	return common.Hash{}
}

// Hash returns the keccak256 hash of b's header.
// The hash is computed on the first call and cached thereafter.
func (b *Block) Hash() common.Hash {
	if hash := b.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := b.header.Hash()
	b.hash.Store(v)
	return v
}

// EncodeNonce converts the given integer to a block nonce.
func EncodeNonce(i uint64) BlockNonce {
	var n BlockNonce
	binary.BigEndian.PutUint64(n[:], i)
	return n
}

// WithSeal returns a new block with the data from b but the header replaced with
// the sealed one.
func (b *Block) WithSeal(header *Header) *Block {
	cpy := *header

	return &Block{
		header:       &cpy,
		transactions: b.transactions,
	}
}
