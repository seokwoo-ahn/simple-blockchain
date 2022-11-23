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

// NewTransaction creates an unsigned transaction
func NewTx(inner TxData) *Transaction {
	tx := new(Transaction)
	tx.setDecoded(inner, 0)
	return tx
}

// TxData is the underlying data of a transaction.
type TxData struct {
	Nonce uint64          // nonce of sender account
	To    *common.Address `rlp:"nil"` // nil means contract creation
	Value *big.Int        // wei amount
	Data  []byte          // contract invocation input data
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

func (tx *Transaction) TxData() TxData { return tx.inner }

func (tx *Transaction) TxTime() time.Time { return tx.time }
