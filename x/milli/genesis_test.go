package milli_test

import (
	"testing"

	keepertest "github.com/millicentnetwork/milli/testutil/keeper"
	"github.com/millicentnetwork/milli/testutil/nullify"
	"github.com/millicentnetwork/milli/x/milli"
	"github.com/millicentnetwork/milli/x/milli/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MilliKeeper(t)
	milli.InitGenesis(ctx, *k, genesisState)
	got := milli.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
