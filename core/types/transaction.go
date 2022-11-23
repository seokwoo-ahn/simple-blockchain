package types

import (
	"math/big"
	"sync/atomic"
	"time"

	"simple-blockchain/common"
)

// Transactions implements DerivableList for transactions.
type Transactions []*Transaction

type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

// NewTx creates a new transaction.
func NewTx(inner TxData) *Transaction {
	tx := new(Transaction)
	tx.setDecoded(inner.copy(), 0)
	return tx
}

// TxData is the underlying data of a transaction.
//
// This is implemented by DynamicFeeTx, LegacyTx and AccessListTx.
type TxData interface {
	txType() byte // returns the type ID
	copy() TxData // creates a deep copy and initializes all fields

	data() []byte
	value() *big.Int
	nonce() uint64
	to() *common.Address

	// rawSignatureValues() (v, r, s *big.Int)
	// setSignatureValues(chainID, v, r, s *big.Int)
}

// Type returns the transaction type.
func (tx *Transaction) Type() uint8 {
	return tx.inner.txType()
}

// Data returns the input data of the transaction.
func (tx *Transaction) Data() []byte { return tx.inner.data() }

// Value returns the ether amount of the transaction.
func (tx *Transaction) Value() *big.Int { return new(big.Int).Set(tx.inner.value()) }

// Nonce returns the sender account nonce of the transaction.
func (tx *Transaction) Nonce() uint64 { return tx.inner.nonce() }

// To returns the recipient address of the transaction.
// For contract-creation transactions, To returns nil.
func (tx *Transaction) To() *common.Address {
	// Copy the pointed-to address.
	ito := tx.inner.to()
	if ito == nil {
		return nil
	}
	cpy := *ito
	return &cpy
}

// Hash returns the transaction hash.
func (tx *Transaction) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}

	var h common.Hash
	tx.hash.Store(h)
	return h
}

// From returns the transaction from.
func (tx *Transaction) From() common.Address {
	if from := tx.from.Load(); from != nil {
		return from.(common.Address)
	}

	var h common.Address
	tx.from.Store(h)
	return h
}

// setDecoded sets the inner transaction and size after decoding.
func (tx *Transaction) setDecoded(inner TxData, size int) {
	tx.inner = inner
	tx.time = time.Now()
	if size > 0 {
		tx.size.Store(common.StorageSize(size))
	}
}
