package vm

import (
	"github.com/Dipper-Protocol/x/vm/common"
	"github.com/Dipper-Protocol/x/vm/keeper"
	"github.com/Dipper-Protocol/x/vm/types"
)

const (
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	CodeKey           = types.CodeKey
	LogKey            = types.LogKey
	StoreDebugKey     = types.StoreDebugKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = keeper.DefaultParamspace
)

type (
	Keeper        = keeper.Keeper
	MsgContract   = types.MsgContract
	CommitStateDB = types.CommitStateDB
	Log           = types.Log

	GenesisState = types.GenesisState
)

var (
	NewKeeper        = keeper.NewKeeper
	NewCommitStateDB = types.NewCommitStateDB

	CreateAddress  = common.CreateAddress
	CreateAddress2 = common.CreateAddress2

	ValidateGenesis = types.ValidateGenesis

	ErrOutOfGas                 = types.ErrOutOfGas
	ErrCodeStoreOutOfGas        = types.ErrCodeStoreOutOfGas
	ErrDepth                    = types.ErrDepth
	ErrTraceLimitReached        = types.ErrTraceLimitReached
	ErrInsufficientBalance      = types.ErrInsufficientBalance
	ErrContractAddressCollision = types.ErrContractAddressCollision
	ErrNoCompatibleInterpreter  = types.ErrNoCompatibleInterpreter
	ErrEmptyInputs              = types.ErrEmptyInputs
	ErrNoCodeExist              = types.ErrNoCodeExist
	ErrMaxCodeSizeExceeded      = types.ErrMaxCodeSizeExceeded
	ErrWriteProtection          = types.ErrWriteProtection
	ErrReturnDataOutOfBounds    = types.ErrReturnDataOutOfBounds
	ErrExecutionReverted        = types.ErrExecutionReverted
	ErrInvalidJump              = types.ErrInvalidJump
	ErrGasUintOverflow          = types.ErrGasUintOverflow
	ErrNoPayload                = types.ErrNoPayload
)
