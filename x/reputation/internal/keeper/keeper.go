package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emag3m/reputation/x/reputation/internal/types"
)

// Keeper of the reputation store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
}

// NewKeeper creates a reputation keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}


// SetReputation sets a reputation
func (k Keeper) SetReputation(ctx sdk.Context, reputation types.Reputation) {
	storageHash := reputation.ApplicationID + reputation.Account.String()
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(reputation)
	key := []byte(storageHash)
	store.Set(key, bz)
}

// GetScavenge returns the scavenge information
func (k Keeper) GetReputation(ctx sdk.Context, storageHash string) (types.Reputation, error) {
	store := ctx.KVStore(k.storeKey)
	var reputation types.Reputation
	byteKey := []byte(storageHash)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &reputation)
	if err != nil {
		return reputation, err
	}
	return reputation, nil
}
