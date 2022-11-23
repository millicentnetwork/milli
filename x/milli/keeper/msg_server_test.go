package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/millicentnetwork/milli/testutil/keeper"
	"github.com/millicentnetwork/milli/x/milli/keeper"
	"github.com/millicentnetwork/milli/x/milli/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MilliKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
