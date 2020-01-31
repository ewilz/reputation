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
	StorageHash  string           `json:"solutionHash" yaml:"solutionHash"`    // solution hash of the reputation
}

// implement fmt.Stringer
func (s Reputation) String() string {
	return strings.TrimSpace(fmt.Sprintf(
	`Account: %s
	Score: %s
	StorageHash: %s`,
		s.Account,
		s.Score,
		s.StorageHash,
	))
}
