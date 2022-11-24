package consensus

import (
	"fmt"
	"testing"

	"simple-blockchain/core/types"
)

func TestMining(t *testing.T) {
	header := sampleHeader()
	abort := make(chan struct{})
	found := make(chan *types.Block, 10)

	block := types.NewBlock(header, nil)
	Mine(block, uint64(0), abort, found)

	fmt.Println(<-found)
}
