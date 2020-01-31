package reputation

import (
	"github.com/emag3m/reputation/x/reputation/internal/keeper"
	"github.com/emag3m/reputation/x/reputation/internal/types"
)

const (
	// TODO: define constants that you would like exposed from the internal package
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc
	NewMsgRecordReputation = types.NewMsgRecordReputation
)

type (
	Keeper       = keeper.Keeper
	CodeType     = types.CodeType
	GenesisState = types.GenesisState

	MsgRecordReputation = types.MsgRecordReputation
)
