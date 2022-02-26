package systemcontract

import (
	// "bytes"
	// "errors"
	// "github.com/ethereum/go-ethereum/accounts/abi"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/assembly/caller"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	// "sort"
)

const (
	punishV3Code = "0x608060405234801561001057600080fd5b50600436106102525760003560e01c80638129fc1c11610146578063c967f90f116100c3578063e0d8ea5311610087578063e0d8ea5314610430578063ea7221a114610438578063efd8d8e21461045e578063f26fdb2914610466578063f342da331461046e578063f62af26c1461047657610252565b8063c967f90f146103c7578063cb1ea725146103e6578063cd5c5342146102c9578063d93d2cb9146103ee578063db78dd281461040b57610252565b8063b3c77d801161010a578063b3c77d801461039f578063be645692146103a7578063bfd1a8a1146103af578063bfd25ce8146103b7578063c7e4b964146103bf57610252565b80638129fc1c14610375578063817da0fd1461037f57806396c7cd6714610387578063a0e8ff661461038f578063a6943ab21461039757610252565b8063399a87eb116101d457806363e1d4511161019857806363e1d4511461032f578063667dd61d14610355578063671051481461035d5780637e30e98614610365578063811b0b391461036d57610252565b8063399a87eb146103075780633a061bd31461030f5780633a73c8811461031757806344c1aa991461031f5780636233be5d1461032757610252565b80631b5e358c1161021b5780631b5e358c146102c15780631d681516146102c95780632897183d146102d1578063308fe582146102d957806332f3c17f146102e157610252565b80627ad09314610257578063079982a2146102715780630aa01884146102955780630cc06e2b1461029d578063158ef93e146102a5575b600080fd5b61025f610493565b60408051918252519081900360200190f35b610279610499565b604080516001600160a01b039092168252519081900360200190f35b61027961049f565b61025f6104a5565b6102ad6104b3565b604080519115158252519081900360200190f35b6102796104bc565b61025f6104cb565b61025f6104d1565b6102796104d7565b61025f600480360360208110156102f757600080fd5b50356001600160a01b03166104dd565b6102796104f8565b6102796104fe565b61025f610512565b61025f610519565b61027961051f565b6102ad6004803603602081101561034557600080fd5b50356001600160a01b031661052e565b61025f61075a565b610279610761565b610279610767565b61027961076d565b61037d610773565b005b6102796107f3565b61025f6107f9565b61025f610809565b610279610819565b61027961081f565b61025f610837565b610279610846565b61025f61084c565b61027961085c565b6103cf61086b565b6040805161ffff9092168252519081900360200190f35b61025f610870565b61037d6004803603602081101561040457600080fd5b5035610876565b610413610b19565b6040805167ffffffffffffffff9092168252519081900360200190f35b61025f610b20565b61037d6004803603602081101561044e57600080fd5b50356001600160a01b0316610b26565b610413610493565b61025f610e34565b610279610e42565b6102796004803603602081101561048c57600080fd5b5035610e48565b61708081565b61f00581565b61a00181565b690a968163f0a57b40000081565b60005460ff1681565b6001546001600160a01b031681565b6105b481565b60065481565b61f00781565b6001600160a01b031660009081526008602052604090205490565b61f00881565b60005461010090046001600160a01b031681565b6206978081565b60055481565b6002546001600160a01b031681565b6000805460ff16610575576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b3361f008146105cb576040805162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c790000000000000000604482015290519081900360640190fd5b6001600160a01b03821660009081526008602052604090205415610603576001600160a01b0382166000908152600860205260408120555b6001600160a01b03821660009081526008602052604090206002015460ff16801561062f575060095415155b15610752576009546001600160a01b038316600090815260086020526040902060010154600019909101146106f9576009805460009190600019810190811061067457fe5b60009182526020808320909101546001600160a01b03868116845260089092526040909220600101546009805492909316935083929181106106b257fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790558583168252600890526040808220600190810154949093168252902001555b600980548061070457fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038416825260089052604081206001810191909155600201805460ff191690555b506001919050565b620d2f0081565b61f00b81565b61f00981565b61a00281565b60005460ff16156107c1576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b600780546001600160a01b03191661f0081790556018600481905560306005556006556000805460ff19166001179055565b61f00381565b6b0d92289838d21a996800000081565b6b015b6a759f4835dc2400000081565b61f00681565b73d1a4e4fbc69e17cabb96fd0d45cf98889b85034781565b6a0422ca8b0a00a42500000081565b61f00a81565b6b07fdacf155df27a328c0000081565b6003546001600160a01b031681565b603381565b60045481565b3341146108b7576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b436000908152600b602052604090205460ff1615610910576040805162461bcd60e51b8152602060048201526011602482015270105b1c9958591e48191958dc99585cd959607a1b604482015290519081900360640190fd5b60005460ff16610956576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b8080438161096057fe5b06156109a6576040805162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b604482015290519081900360640190fd5b436000908152600b60205260409020805460ff191660011790556009546109cc57610b15565b60005b600954811015610aea57600654600554816109e657fe5b0460086000600984815481106109f857fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020541115610aa95760065460055481610a3057fe5b046008600060098481548110610a4257fe5b60009182526020808320909101546001600160a01b03168352820192909252604001812054600980549390910392600892919085908110610a7f57fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055610ae2565b60006008600060098481548110610abc57fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020555b6001016109cf565b506040517f181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db464190600090a15b5050565b6201518081565b60095490565b334114610b67576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b60005460ff16610bad576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b436000908152600a602052604090205460ff1615610c05576040805162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c1d5b9a5cda195960821b604482015290519081900360640190fd5b436000908152600a60209081526040808320805460ff191660011790556001600160a01b0384168352600890915290206002015460ff16610cae57600980546001600160a01b038316600081815260086020526040812060018082018590558085019095557f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af90930180546001600160a01b0319168317905552600201805460ff191690911790555b6001600160a01b03811660009081526008602052604090208054600101908190556005549081610cda57fe5b06610d635760075460408051632566392f60e01b81526001600160a01b03848116600483015291519190921691632566392f91602480830192600092919082900301818387803b158015610d2d57600080fd5b505af1158015610d41573d6000803e3d6000fd5b5050506001600160a01b03821660009081526008602052604081205550610df2565b6004546001600160a01b03821660009081526008602052604090205481610d8657fe5b06610df257600754604080516305dd095960e41b81526001600160a01b03848116600483015291519190921691635dd0959091602480830192600092919082900301818387803b158015610dd957600080fd5b505af1158015610ded573d6000803e3d6000fd5b505050505b6040805142815290516001600160a01b038316917f770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e919081900360200190a250565b693ceab05409db274d269381565b61f00481565b60098181548110610e5557fe5b6000918252602090912001546001600160a01b031690508156fea26469706673582212209e986d44ee40cdee125a6ed094c5f5143cdc39b3db6d018819f0b4016268ce0464736f6c634300060c0033"
)

type forkPunishV3 struct {
}

func (f *forkPunishV3) GetName() string {
	return PunishV2ContractName
}

func (f *forkPunishV3) Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) (err error) {
	contractCode := common.FromHex(punishV3Code)

	//write code to sys contract
	state.SetCode(PunishV3ContractAddr, contractCode)
	log.Debug("Write code to system contract account", "addr", PunishV3ContractAddr.String(), "code", punishV3Code)

	return
}

func (f *forkPunishV3) Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {

	method := "initialize"
	data, err := GetInteractiveABI()[PunishV3ContractName].Pack(method)
	if err != nil {
		log.Error("Can't pack data for initialize", "error", err)
		return err
	}

	msg := types.NewMessage(header.Coinbase, &PunishV3ContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	_, err = caller.ExecuteMsg(msg, state, header, chainContext, config)

	return
}
