package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/example/blog/x/blog/keeper"
	"github.com/example/blog/x/blog/types"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	//var value  = sdk.Coins{sdk.NewInt64Coin("token",5)}
	//BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	payment := sdk.Coins{sdk.NewInt64Coin("MYTOKEN", 5)}
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	//err1 := k.CoinKeeper.SubtractCoins(ctx, addr, sdk.Coins{sdk.NewInt64Coin("MYTOKEN", 5)})
	//err :=k.CoinKeeper.BurnCoins(ctx,"cosmos1e3mttyl2m4820tggftpkg6mpszu6nkvdx2h54g",value)
	err1 := k.CoinKeeper.SendCoins(ctx, addr, moduleAcct, payment)
	if err1 != nil {
		return nil, err1
	}
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
