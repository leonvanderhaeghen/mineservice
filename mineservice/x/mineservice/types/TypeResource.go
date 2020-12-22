package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Resource struct {
	Owner sdk.AccAddress `json:"owner" yaml:"owner"`
	ID      string         `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	MineID string `json:"mineID" yaml:"mineID"`
    Price sdk.Coins `json:"price" yaml:"price"`
    Selling bool `json:"selling" yaml:"selling"`
	MinPrice sdk.Coins `json:"minPrice" yaml:"minPrice"`
	Amount int `json:"amount" yaml:"amount"`

}