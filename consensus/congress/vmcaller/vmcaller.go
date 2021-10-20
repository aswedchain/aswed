package vmcaller

import (
	"github.com/aswedchain/aswed/core"
	"github.com/aswedchain/aswed/core/state"
	"github.com/aswedchain/aswed/core/types"
	"github.com/aswedchain/aswed/core/vm"
	"github.com/aswedchain/aswed/params"
)

// ExecuteMsg executes transaction sent to system contracts.
func ExecuteMsg(msg core.Message, state *state.StateDB, header *types.Header, chainContext core.ChainContext, chainConfig *params.ChainConfig) (ret []byte, err error) {
	// Set gas price to zero
	context := core.NewEVMContext(msg, header, chainContext, &(header.Coinbase))
	vmenv := vm.NewEVM(context, state, chainConfig, vm.Config{})

	ret, _, err = vmenv.Call(vm.AccountRef(msg.From()), *msg.To(), msg.Data(), msg.Gas(), msg.Value())

	return ret, err
}
