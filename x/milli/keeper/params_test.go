package keeper_test

import (
	"testing"

	testkeeper "github.com/millicentnetwork/milli/testutil/keeper"
	"github.com/millicentnetwork/milli/x/milli/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MilliKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
