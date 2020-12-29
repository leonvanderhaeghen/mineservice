package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Player struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Name string `json:"name" yaml:"name"`
    Invetory []Resource `json:"invetory" yaml:"invetory"`
    Mines []string `json:"mines" yaml:"mines"`
}