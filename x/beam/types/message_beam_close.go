package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCloseBeam{}

// NewMsgCloseBeam Build a MsgCloseBeam instance
func NewMsgCloseBeam(updater string, id string) *MsgCloseBeam {
	return &MsgCloseBeam{
		Updater: updater,
		Id:      id,
	}
}

// Route dunno
func (msg MsgCloseBeam) Route() string {
	return RouterKey
}

// Type Return the message type
func (msg MsgCloseBeam) Type() string {
	return "CloseBeam"
}

// GetSigners Return the list of signers for the given message
func (msg *MsgCloseBeam) GetSigners() []sdk.AccAddress {
	updater, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{updater}
}

// GetSignBytes Return the generated bytes from the signature
func (msg *MsgCloseBeam) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic Validate the msg payload
func (msg *MsgCloseBeam) ValidateBasic() error {
	// Ensure the address is correct and that we are able to acquire it
	_, err := sdk.AccAddressFromBech32(msg.Updater)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	return nil
}