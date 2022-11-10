package consensus

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
)

func TestMining(t *testing.T) {
	header := SampleHeader()
	abort := make(chan struct{})
	found := make(chan *types.Block, 10)

	block := types.NewBlock(header, nil, nil, nil, nil)
	Mine(block, uint64(0), abort, found)

	fmt.Println(<-found)
}
