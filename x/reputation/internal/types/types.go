package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Reputation is the Reputation struct
type Reputation struct {
	Account      sdk.AccAddress   `json:"account" yaml:"account"`              // address of the reputation holder
	Score        int              `json:"score" yaml:"account"`                // integer reputation score
	ApplicationID  string         `json:"applicationId" yaml:"applicationID"`  // application id for which reputation corresponds
}

// implement fmt.Stringer
func (s Reputation) String() string {
	return strings.TrimSpace(fmt.Sprintf(
	`Account: %s
	Score: %s
	ApplicationID: %s`,
		s.Account,
		s.Score,
		s.ApplicationID,
	))
}
