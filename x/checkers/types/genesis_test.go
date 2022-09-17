package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&GenesisState{
			StoredGameList: []*StoredGame{},
			NextGame:       &NextGame{"", uint64(1)},
		},
		DefaultGenesis())
}
