package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintCoins = "mint_coins"

var _ sdk.Msg = &MsgMintCoins{}

func NewMsgMintCoins(creator string, amt sdk.Coin) *MsgMintCoins {
	return &MsgMintCoins{
		Creator: creator,
		Amt:     amt,
	}
}

func (msg *MsgMintCoins) Route() string {
	return RouterKey
}

func (msg *MsgMintCoins) Type() string {
	return TypeMsgMintCoins
}

func (msg *MsgMintCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
