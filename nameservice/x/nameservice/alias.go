package nameservice

import (
	"cosmos-sdk/nameservice/x/nameservice/internal/keeper"
	"cosmos-sdk/nameservice/x/nameservice/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper        = keeper.NewKeeper
	NewQuerier       = keeper.NewQuerier
	NewMsgBuyName    = types.NewMsgBuyName
	NewMsgSetName    = types.NewMsgSetName
	NewMsgDeleteName = types.NewMsgDeleteName
	NewWhois         = types.NewWhois
	ModuleCdc        = types.ModuleCdc
	RegisterCodec    = types.RegisterCodec
)

type (
	Keeper          = keeper.Keeper
	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
)
