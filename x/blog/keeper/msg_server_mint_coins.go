package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/blog/types"
)

func (k msgServer) MintCoins(goCtx context.Context, msg *types.MsgMintCoins) (*types.MsgMintCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	coins := sdk.NewCoins(msg.Amt)
	err = k.Bank.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		panic(err)
	}

	err = k.Bank.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, coins)
	if err != nil {
		panic(err)
	}

	return &types.MsgMintCoinsResponse{}, nil
}
