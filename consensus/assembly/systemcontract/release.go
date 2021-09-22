package systemcontract

import (
	// "bytes"
	// "errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/assembly/caller"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math"
	"math/big"
	// "sort"
)

type Release struct {
	abi          abi.ABI
	contractAddr common.Address
}

func NewRelease() *Release {
	return &Release{
		abi:          abiMap[ReleaseContractName],
		contractAddr: ReleaseContractAddr,
	}
}

func (v *Release) Trigger(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (bool, error) {
	method := "trigger"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for trigger", "error", err)
		return false, err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return false, err
	}
	// unpack data
	_, err = v.abi.Unpack(method, result)
	if err != nil {
		return false, err
	}
	return true, nil
}