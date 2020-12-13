package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Mine struct {
	ID      string         `json:"id" yaml:"id"`
    Name string `json:"name" yaml:"name"`
    Price sdk.Coins `json:"price" yaml:"price"`
    Owner sdk.AccAddress `json:"owner" yaml:"owner"`
    Selling bool `json:"selling" yaml:"selling"`
    Efficiency int `json:"efficiency" yaml:"efficiency"`
    Invetory string `json:"invetory" yaml:"invetory"` //later veranderen naar type invetory
	Resources []string `json:"resources" yaml:"resources"` //later veranderen naar type resources
	ResourceCounter int `json:"ResourceCounter" yaml:"ResourceCounter"`
    UraniumCost int `json:"uraniumCost" yaml:"uraniumCost"`
}