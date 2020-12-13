package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNotSelling = sdkerrors.Register(ModuleName,1,"This item is currently not for sale")
)
