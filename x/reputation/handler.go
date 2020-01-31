package reputation

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/emag3m/reputation/x/reputation/internal/types"
)

// NewHandler creates an sdk.Handler for all the reputation type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgRecordReputation:
			return handleMsgRecordReputation(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgRecordReputation creates a new reputation and moves the reward into escrow
func handleMsgRecordReputation(ctx sdk.Context, k Keeper, msg MsgRecordReputation) sdk.Result {
	var reputation = types.Reputation{
		Account:       msg.Account,
		Score:         msg.Score,
		ApplicationID: msg.ApplicationID,
	}
		k.SetReputation(ctx, reputation)
		return sdk.Result{}
}
