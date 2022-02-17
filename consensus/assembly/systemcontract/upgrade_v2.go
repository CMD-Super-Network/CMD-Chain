package systemcontract

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type IUpgradeAction2 interface {
	GetName() string
	Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) error
	Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) error
}

var (
	sysContracts2 []IUpgradeAction2
)

func init() {
	sysContracts2 = []IUpgradeAction2{
		// &hardForkSysGov{},
		// &hardForkAddressList{},
		// &forkValidatorsV2{},
		// &forkPunishV2{},
		// &forkProposalV2{},
		// &forkCMDV2{},
		// &hardForkPunishV1{},
		&forkCMDV3{},
		&forkValidatorsV3{},
		&forkPunishV3{},
		&forkProposalV3{},
	}
}

// func init2(){

// }

func ApplySystemContractUpgrade2(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {
	if config == nil || header == nil || state == nil {
		return
	}
	height := header.Number

	for _, contract := range sysContracts2 {
		log.Info("system contract v3 upgrade", "name", contract.GetName(), "height", height, "chainId", config.ChainID.String())

		err = contract.Update(config, height, state)
		if err != nil {
			log.Error("Upgrade system contract update error", "name", contract.GetName(), "err", err)
			return
		}

		log.Info("system contract upgrade v3 execution", "name", contract.GetName(), "height", header.Number, "chainId", config.ChainID.String())

		err = contract.Execute(state, header, chainContext, config)
		if err != nil {
			log.Error("Upgrade system contract execute error", "name", contract.GetName(), "err", err)
			return
		}
	}
	// Update the state with pending changes
	log.Info("update done!")
	state.Finalise(true)

	return
}
