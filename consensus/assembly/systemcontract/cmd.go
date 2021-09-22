package systemcontract

import (
	// "bytes"
	"errors"
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

type CMD struct {
	abi          abi.ABI
	contractAddr common.Address
}


func NewCMD() *CMD {
	return &CMD{
		abi:          abiMap[CMDContractName],
		contractAddr: CMDContractAddr,
	}
}

func (v *CMD) GetTotalSupply(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (*big.Int, error) {
	method := "totalSupply"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for totalSupply", "error", err)
		return big.NewInt(0), err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return new(big.Int), err
	}
	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return new(big.Int), err
	}
	count, ok := ret[0].(*big.Int)
	if !ok {
		return new(big.Int), errors.New("invalid output")
	}
	return count, nil
}

func (v *CMD) GetCirculation(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (*big.Int, error) {
	method := "circulation"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for circulation", "error", err)
		return big.NewInt(0), err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return new(big.Int), err
	}
	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return new(big.Int), err
	}
	count, ok := ret[0].(*big.Int)
	if !ok {
		return new(big.Int), errors.New("invalid output")
	}
	return count, nil
}


func (v *CMD) AdjustRemainderMine(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (bool, error) {
	method := "adjustRemainderMine"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for adjustRemainderMine", "error", err)
		return false, err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		log.Error("1")
		return false, err
	}
	// unpack data
	_, err = v.abi.Unpack(method, result)
	if err != nil {
		log.Error("2")
		return false, err
	}

	return true, nil
}

func (v *CMD) GetValidatorReward(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (*big.Int,error){
	method := "getValidatorReward"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for getValidatorReward", "error", err)
		return big.NewInt(0), err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return new(big.Int), err
	}
	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return new(big.Int), err
	}
	count, ok := ret[0].(*big.Int)
	if !ok {
		return new(big.Int), errors.New("invalid output")
	}
	return count, nil
}

func (v *CMD) GetMinerReward(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (*big.Int,error){
	method := "getMinerReward"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for getMinerReward", "error", err)
		return big.NewInt(0), err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return new(big.Int), err
	}
	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return new(big.Int), err
	}
	count, ok := ret[0].(*big.Int)
	if !ok {
		return new(big.Int), errors.New("invalid output")
	}
	return count, nil
}