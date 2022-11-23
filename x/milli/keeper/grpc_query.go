package keeper

import (
	"github.com/millicentnetwork/milli/x/milli/types"
)

var _ types.QueryServer = Keeper{}
