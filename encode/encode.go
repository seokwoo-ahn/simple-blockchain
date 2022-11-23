package encode

import (
	"encoding/json"
	"simple-blockchain/core/types"
)

func TxEncode(tx *types.Transaction) string {
	var encoded string
	txData := tx.TxData()
	if temp, err := json.Marshal(txData); err != nil {
		panic(err)
	} else {
		encoded = encoded + string(temp)
	}
	encoded = encoded + tx.TxTime().String()
	return encoded
}
