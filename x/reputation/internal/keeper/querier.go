package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emag3m/reputation/x/reputation/internal/types"
)

// NewQuerier creates a new querier for reputation clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryGetReputation:
			return getReputation(ctx, path[1:], k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown reputation query endpoint")
		}
	}
}

func getReputation(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError sdk.Error) {
	storagePath := path[0]
	reputation, err := k.GetReputation(ctx, storagePath)
	if err != nil {
		return nil, sdk.NewError(types.DefaultCodespace, types.CodeInvalid, err.Error())
	}

	res, err = codec.MarshalJSONIndent(k.cdc, reputation)
	if err != nil {
		return nil, sdk.NewError(types.DefaultCodespace, types.CodeInvalid, "Could not marshal result to JSON")
	}
	return res, nil
}
