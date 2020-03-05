package types

import (
	sdk "github.com/Dipper-Protocol/types"
)

// DefaultCodespace is the Module Name
const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeNameDoesNotExist sdk.CodeType = 101

	CodeNotEnoughTokenForBorrow sdk.CodeType = 201
	CodeTooMuchAmountToRepay sdk.CodeType = 202
	CodeNotEnoughAmountCoinForWithdraw sdk.CodeType = 203
)


func ErrNotEnoughTokenForBorrow(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNotEnoughTokenForBorrow, "Not enough token for borrow")
}

func ErrTooMuchAmountToRepay(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeTooMuchAmountToRepay, "too much amount to repay")
}

func ErrNotEnoughAmountCoinForWithdraw(codespace sdk.CodespaceType) sdk.Error{
	return sdk.NewError(codespace, CodeNotEnoughAmountCoinForWithdraw, "not enough amount or coin for withdraw")
}