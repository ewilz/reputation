package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgRecordReputation{}

// MsgRecordReputation - struct for recording reputation score
type MsgRecordReputation struct {
	Account      sdk.AccAddress `json:"creator" yaml:"creator"`           // account address associated to reputation score
	Score        int            `json:"score" yaml:"score"`               // reputation score
	StorageHash  string         `json:"storageHash" yaml:"storageHash"`   // storage hash representing app associated with reputation score
}

// NewMsgRecordReputation is a constructor function for MsgGreet
func NewMsgRecordReputation(account sdk.AccAddress, score int, storageHash string) MsgRecordReputation {
	return MsgRecordReputation{
		Account:      account,
		Score:        score,
		StorageHash:  storageHash,
	}
}

const RecordReputationConst = "RecordReputation"

// Route should return the name of the module
func (msg MsgRecordReputation) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRecordReputation) Type() string { return RecordReputationConst }

func (msg MsgRecordReputation) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Account)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgRecordReputation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic runs stateless checks on the message
func (msg MsgRecordReputation) ValidateBasic() sdk.Error {
	if msg.Account.Empty() {
		return sdk.ErrInvalidAddress(msg.Account.String())
	}
  if msg.StorageHash == "" {
		return sdk.ErrUnknownRequest("StorageHash field cannot be empty")
	}
	return nil
}
