package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProvisionModel = "provision_model"

var _ sdk.Msg = &MsgProvisionModel{}

func NewMsgProvisionModel(signer string, vid int32, pid int32, softwareVersion uint32, softwareVersionString string, cDVersionNumber uint32, provisionalDate string, certificationType string, reason string) *MsgProvisionModel {
	return &MsgProvisionModel{
		Signer:                signer,
		Vid:                   vid,
		Pid:                   pid,
		SoftwareVersion:       softwareVersion,
		SoftwareVersionString: softwareVersionString,
		CDVersionNumber:       cDVersionNumber,
		ProvisionalDate:       provisionalDate,
		CertificationType:     certificationType,
		Reason:                reason,
	}
}

func (msg *MsgProvisionModel) Route() string {
	return RouterKey
}

func (msg *MsgProvisionModel) Type() string {
	return TypeMsgProvisionModel
}

func (msg *MsgProvisionModel) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{signer}
}

func (msg *MsgProvisionModel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProvisionModel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address (%s)", err)
	}
	return nil
}