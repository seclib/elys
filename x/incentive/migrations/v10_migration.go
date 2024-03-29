package migrations

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (m Migrator) V10Migration(ctx sdk.Context) error {
	params := m.keeper.GetParams(ctx)
	params.DistributionInterval = 1200
	if params.LpIncentives != nil {
		params.LpIncentives.DistributionEpochInBlocks = sdk.NewInt(params.DistributionInterval)
	}
	if params.StakeIncentives != nil {
		params.StakeIncentives.DistributionEpochInBlocks = sdk.NewInt(params.DistributionInterval)
	}
	m.keeper.SetParams(ctx, params)
	return nil
}
