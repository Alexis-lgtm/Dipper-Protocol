package vm

import (
	"github.com/Dipper-Protocol/x/vm/types"
	"math/big"
	"testing"

	sdk "github.com/Dipper-Protocol/types"
	"github.com/Dipper-Protocol/x/auth"
	"github.com/Dipper-Protocol/x/bank"
	"github.com/Dipper-Protocol/x/params"
)

type dummyContractRef struct {
	calledForEach bool
}

func (dummyContractRef) ReturnGas(*big.Int)       {}
func (dummyContractRef) Address() sdk.AccAddress  { return sdk.AccAddress{} }
func (dummyContractRef) Value() *big.Int          { return new(big.Int) }
func (dummyContractRef) SetCode(sdk.Hash, []byte) {}
func (d *dummyContractRef) ForEachStorage(callback func(key, value sdk.Hash) bool) {
	d.calledForEach = true
}
func (d *dummyContractRef) SubBalance(amount *big.Int) {}
func (d *dummyContractRef) AddBalance(amount *big.Int) {}
func (d *dummyContractRef) SetBalance(*big.Int)        {}
func (d *dummyContractRef) SetNonce(uint64)            {}
func (d *dummyContractRef) Balance() *big.Int          { return new(big.Int) }

type dummyStatedb struct {
	CommitStateDB
}

func (*dummyStatedb) GetRefund() uint64 { return 1337 }

func TestStoreCapture(t *testing.T) {
	keyAcc := sdk.NewKVStoreKey(auth.StoreKey)
	keyParams := sdk.NewKVStoreKey(params.StoreKey)
	tkeyParams := sdk.NewTransientStoreKey(params.TStoreKey)
	paramsKeeper := params.NewKeeper(types.ModuleCdc, keyParams, tkeyParams)
	accountKeeper := auth.NewAccountKeeper(types.ModuleCdc, keyAcc, paramsKeeper.Subspace(auth.DefaultParamspace), auth.ProtoBaseAccount)
	vmKeeper := NewKeeper(
		types.ModuleCdc,
		sdk.NewKVStoreKey(StoreKey),
		sdk.NewKVStoreKey(CodeKey),
		sdk.NewKVStoreKey(LogKey),
		sdk.NewKVStoreKey(StoreDebugKey),
		paramsKeeper.Subspace(bank.DefaultParamspace),
		accountKeeper)

	var (
		env      = NewEVM(Context{}, vmKeeper.StateDB, Config{})
		logger   = NewStructLogger(nil)
		mem      = NewMemory()
		stack    = newstack()
		contract = NewContract(&dummyContractRef{}, &dummyContractRef{}, new(big.Int), 0)
	)
	stack.push(big.NewInt(1))
	stack.push(big.NewInt(0))
	var index sdk.Hash
	logger.CaptureState(env, 0, SSTORE, 0, 0, mem, stack, contract, 0, nil)
	if len(logger.changedValues[contract.Address().String()]) == 0 {
		t.Fatalf("expected exactly 1 changed value on address %s, got %d", contract.Address().String(), len(logger.changedValues[contract.Address().String()]))
	}
	exp := sdk.BigToHash(big.NewInt(1))
	if logger.changedValues[contract.Address().String()][index] != exp {
		t.Errorf("expected %x, got %x", exp, logger.changedValues[contract.Address().String()][index])
	}
}
