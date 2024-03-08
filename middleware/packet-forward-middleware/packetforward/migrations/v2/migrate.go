package v2

import (
	sdkmath "cosmossdk.io/math"
	"fmt"

	"github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/exported"
	"github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"

	storetypes "cosmossdk.io/store/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrate migrates the x/packetforward module state from the consensus version 1 to
// version 2. Specifically, it takes the parameters that are currently stored
// and managed by the x/params modules and stores them directly into the x/packetforward
// module state.
func Migrate(
	_ sdk.Context,
	store storetypes.KVStore,
	_ exported.Subspace,
	cdc codec.BinaryCodec,
) error {
	var currParams types.Params
	currParams.FeePercentage = sdkmath.LegacyNewDec(0)
	//legacySubspace.GetParamSet(ctx, &currParams)

	if err := currParams.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&currParams)
	store.Set(types.ParamsKey, bz)

	return validate(store, cdc, currParams)
}

func validate(store storetypes.KVStore, cdc codec.BinaryCodec, currParams types.Params) error {
	var res types.Params
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return fmt.Errorf("expected params at key %s but not found", types.ParamsKey)
	}

	if err := cdc.Unmarshal(bz, &res); err != nil {
		return err
	}

	if !currParams.FeePercentage.Equal(res.FeePercentage) {
		return fmt.Errorf("expected %s but got %s", currParams, res)
	}

	return nil
}
