package types

import (
	"math/big"
	"sync/atomic"

	"simple-blockchain/common"
)

// Transactions implements DerivableList for transactions.
type Transactions []*Transaction

type Transaction struct {
	inner TxData // Consensus contents of a transaction
	// time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	// size atomic.Value
	// from atomic.Value
}

// TxData is the underlying data of a transaction.
//
// This is implemented by DynamicFeeTx, LegacyTx and AccessListTx.
type TxData interface {
	txType() byte // returns the type ID
	copy() TxData // creates a deep copy and initializes all fields

	chainID() *big.Int
	data() []byte
	gas() uint64
	gasPrice() *big.Int
	gasTipCap() *big.Int
	gasFeeCap() *big.Int
	value() *big.Int
	nonce() uint64
	to() *common.Address

	rawSignatureValues() (v, r, s *big.Int)
	setSignatureValues(chainID, v, r, s *big.Int)
}

// Data returns the input data of the transaction.
func (tx *Transaction) Data() []byte { return tx.inner.data() }

// Hash returns the transaction hash.
func (tx *Transaction) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}

	var h common.Hash
	tx.hash.Store(h)
	return h
}
