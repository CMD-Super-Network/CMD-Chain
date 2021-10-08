package systemcontract

import (
	"bytes"
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
	"sort"
)

type Validator struct {
	abi          abi.ABI
	contractAddr common.Address
}

func NewValidatorV0() *Validator {
	return &Validator{
		abi:          abiMap[ValidatorsContractName],
		contractAddr: ValidatorsContractAddr,
	}
}

func (v *Validator) GetTopValidators(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) ([]common.Address, error) {
	method := "getTopValidators"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for getTopValidators", "error", err)
		return []common.Address{}, err
	}

	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return []common.Address{}, err
	}

	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return []common.Address{}, err
	}
	if len(ret) != 1 {
		return []common.Address{}, errors.New("invalid params length")
	}
	validators, ok := ret[0].([]common.Address)
	if !ok {
		return []common.Address{}, errors.New("invalid validators format")
	}
	sort.Slice(validators, func(i, j int) bool {
		return bytes.Compare(validators[i][:], validators[j][:]) < 0
	})
	return validators, err
}

func (v *Validator) GetValidatorFeeAddr(val common.Address, statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (common.Address, error) {
	method := "getValidatorInfo"
	data, err := v.abi.Pack(method, val)
	if err != nil {
		log.Error("Can't pack data for GetValidatorInfo", "error", err)
		return common.Address{}, err
	}
	msg := types.NewMessage(header.Coinbase, &v.contractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	// use parent
	result, err := caller.ExecuteMsg(msg, statedb, header, chainContext, config)
	if err != nil {
		return common.Address{}, err
	}
	// unpack data
	ret, err := v.abi.Unpack(method, result)
	if err != nil {
		return common.Address{}, err
	}
	feeAddr, ok := ret[0].(common.Address)
	if !ok {
		return common.Address{}, errors.New("invalid output")
	}
	return feeAddr, nil
}

func (v *Validator) CountStakedActiveValidator(statedb *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (*big.Int, error) {
	method := "countStakedActiveValidator"
	data, err := v.abi.Pack(method)
	if err != nil {
		log.Error("Can't pack data for countStakedActiveValidator", "error", err)
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