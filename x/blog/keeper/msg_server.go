package keeper

import (
	"planet/x/blog/types"
)

type msgServer struct {
	Keeper
	Bank types.BankKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, bank types.BankKeeper) types.MsgServer {
	return &msgServer{Keeper: keeper, Bank: bank}
}

var _ types.MsgServer = msgServer{}
