package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// get the community coins
func (k Keeper) GetFeePoolCommunityCoins(ctx sdk.Context) sdk.DecCoins {
	return k.GetFeePool(ctx).CommunityPool
}
