package systemcontract


import (
	// "bytes"
	// "errors"
	// "github.com/ethereum/go-ethereum/accounts/abi"
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

const (
	proposalV2Code = "0x60806040526004361061027c5760003560e01c806390a693941161014f578063c7e4b964116100c1578063e65044a71161007a578063e65044a714610847578063e823c81414610871578063efd8d8e214610886578063f26fdb291461089b578063f3982b29146108b0578063ff1b636d146108cd5761027c565b8063c7e4b964146107a2578063c967f90f146107b7578063cd15c722146107e3578063cd5c5342146107eb578063d9f18dc914610800578063db78dd28146108155761027c565b8063a224cee711610113578063a224cee7146106bc578063b3c77d8014610739578063b9419a3d1461074e578063bb5d40d814610763578063be64569214610778578063bfd25ce81461078d5761027c565b806390a693941461065357806396c7cd67146106685780639ba92fad1461067d578063a0e8ff6614610692578063a1a7536b146106a75761027c565b80633a061bd3116101f3578063667dd61d116101ac578063667dd61d146105b757806367105148146105cc57806378e97925146105e1578063811b0b39146105f6578063817da0fd1461060b57806382c4b3b2146106205761027c565b80633a061bd3146104b45780633a73c881146104c95780633fa3b8ee146104de57806349c2a1a6146105085780635e235a61146105785780636233be5d146105a25761027c565b806315ea27811161024557806315ea2781146103545780631a9667b7146103875780631b5e358c1461039c5780631d681516146103b15780632e89dce1146103c657806332ed5b12146103db5761027c565b80627ad093146102815780630aa01884146102a85780630cc06e2b146102d957806312041459146102ee578063158ef93e1461032b575b600080fd5b34801561028d57600080fd5b506102966108e2565b60408051918252519081900360200190f35b3480156102b457600080fd5b506102bd6108e8565b604080516001600160a01b039092168252519081900360200190f35b3480156102e557600080fd5b506102966108ee565b3480156102fa57600080fd5b506103296004803603604081101561031157600080fd5b506001600160a01b03813581169160200135166108fc565b005b34801561033757600080fd5b50610340610c4d565b604080519115158252519081900360200190f35b34801561036057600080fd5b506103406004803603602081101561037757600080fd5b50356001600160a01b0316610c56565b34801561039357600080fd5b50610329610d07565b3480156103a857600080fd5b506102bd610dd3565b3480156103bd57600080fd5b50610296610dd9565b3480156103d257600080fd5b50610329610ddf565b3480156103e757600080fd5b50610405600480360360208110156103fe57600080fd5b50356111c9565b60405180886001600160a01b031681526020018060200187815260200186151581526020018581526020018481526020018315158152602001828103825288818151815260200191508051906020019080838360005b8381101561047357818101518382015260200161045b565b50505050905090810190601f1680156104a05780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b3480156104c057600080fd5b506102bd6112a2565b3480156104d557600080fd5b506102966112a8565b3480156104ea57600080fd5b506103296004803603602081101561050157600080fd5b50356112af565b6103406004803603602081101561051e57600080fd5b81019060208101813564010000000081111561053957600080fd5b82018360208201111561054b57600080fd5b8035906020019184600183028401116401000000008311171561056d57600080fd5b50909250905061146c565b34801561058457600080fd5b506102966004803603602081101561059b57600080fd5b503561183c565b3480156105ae57600080fd5b506102bd61185a565b3480156105c357600080fd5b50610296611860565b3480156105d857600080fd5b506102bd611867565b3480156105ed57600080fd5b5061029661186d565b34801561060257600080fd5b506102bd611873565b34801561061757600080fd5b506102bd611879565b34801561062c57600080fd5b506103406004803603602081101561064357600080fd5b50356001600160a01b031661187f565b34801561065f57600080fd5b50610296611894565b34801561067457600080fd5b506102966118a2565b34801561068957600080fd5b506102966118b2565b34801561069e57600080fd5b506102966118b8565b3480156106b357600080fd5b506102966118c8565b3480156106c857600080fd5b50610329600480360360208110156106df57600080fd5b8101906020810181356401000000008111156106fa57600080fd5b82018360208201111561070c57600080fd5b8035906020019184602083028401116401000000008311171561072e57600080fd5b5090925090506118cd565b34801561074557600080fd5b506102bd611a4f565b34801561075a57600080fd5b50610296611a67565b34801561076f57600080fd5b50610296611a6c565b34801561078457600080fd5b50610296611a79565b34801561079957600080fd5b50610296611a88565b3480156107ae57600080fd5b506102bd611a98565b3480156107c357600080fd5b506107cc611aac565b6040805161ffff9092168252519081900360200190f35b610340611ab1565b3480156107f757600080fd5b50610296611ab6565b34801561080c57600080fd5b50610296611abb565b34801561082157600080fd5b5061082a611ac0565b6040805167ffffffffffffffff9092168252519081900360200190f35b34801561085357600080fd5b506103296004803603602081101561086a57600080fd5b5035611ac7565b34801561087d57600080fd5b50610296611de3565b34801561089257600080fd5b5061082a6108e2565b3480156108a757600080fd5b50610296611de9565b610340600480360360208110156108c657600080fd5b5035611df7565b3480156108d957600080fd5b506102bd611ffb565b61708081565b61a00181565b690a968163f0a57b40000081565b3361f0041461094d576040805162461bcd60e51b815260206004820152601860248201527756616c696461746f727320636f6e7472616374206f6e6c7960401b604482015290519081900360640190fd5b6001600160a01b03821660009081526002602052604090205460ff166109aa576040805162461bcd60e51b815260206004820152600d60248201526c3737ba103b30b634b230ba37b960991b604482015290519081900360640190fd5b6001600160a01b03808316600090815260026020526040808220805460ff1990811690915592841682529020805490911660011790556109e861238d565b6001600160a01b038216815260408051808201909152600c81526b031b430b733b290333937b6960a51b6020820152610a2990610a248561200a565b6120fe565b8160200181905250428160400181815250506000338383602001514260405160200180856001600160a01b031660601b8152601401846001600160a01b031660601b815260140183805190602001908083835b60208310610a9b5780518252601f199092019160209182019101610a7c565b51815160209384036101000a6000190180199092169116179052920193845250604080518085038152938201815283519382019390932060008181526004835293909320885181546001600160a01b0319166001600160a01b03909116178155888201518051949850899750909550610b1d94506001860193910191506123d7565b506040828101516002830155606083015160038301805491151560ff19928316179055608084015160048085019190915560a0850151600585015560c0909401516006909301805493151593909116929092179091556008546000848152602084815283822054845163415e9ec960e11b81526001600160a01b0391821696810196909652935161010090930493909316936382bd3d9293602480830194919391928390030190829087803b158015610bd557600080fd5b505af1158015610be9573d6000803e3d6000fd5b505050506040513d6020811015610bff57600080fd5b50506040805142815290516001600160a01b03808616929087169184917fc10f2f4d53a0e342536c6af3cce9c6ee25c32dbb323521ce0e1d4494a3e362e8919081900360200190a450505050565b60005460ff1681565b60003361f00414610ca9576040805162461bcd60e51b815260206004820152601860248201527756616c696461746f727320636f6e7472616374206f6e6c7960401b604482015290519081900360640190fd5b6001600160a01b038216600081815260026020908152604091829020805460ff19169055815142815291517f4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a9281900390910190a25060015b919050565b6003546001600160a01b03163314610d5f576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e206f70657261746560501b604482015290519081900360640190fd5b62030d404310610da05760405162461bcd60e51b81526004018080602001828103825260348152602001806124cd6034913960400191505060405180910390fd5b6040514790339082156108fc029083906000818181858888f19350505050158015610dcf573d6000803e3d6000fd5b5050565b61f00581565b6105b481565b3373d1a4e4fbc69e17cabb96fd0d45cf98889b85034714610e3f576040805162461bcd60e51b81526020600482015260156024820152746f6e6c7920676f7665726e6163652063616e20646f60581b604482015290519081900360640190fd5b60085460ff16610e84576040805162461bcd60e51b815260206004820152600b60248201526a1a5cc8199a5b9a5cda195960aa1b604482015290519081900360640190fd5b600654600754600090829081610e9657fe5b04905060005b828110156111ba57816004600060068481548110610eb657fe5b9060005260206000200154815260200190815260200160002060050154106110e257600860019054906101000a90046001600160a01b03166001600160a01b03166382bd3d926004600060068581548110610f0d57fe5b9060005260206000200154815260200190815260200160002060000160009054906101000a90046001600160a01b03166040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050602060405180830381600087803b158015610f7c57600080fd5b505af1158015610f90573d6000803e3d6000fd5b505050506040513d6020811015610fa657600080fd5b50506006805460019160029160009160049183919087908110610fc557fe5b60009182526020808320919091015483528281019390935260409182018120546001600160a01b031684529183019390935291018120805460ff1916921515929092179091556006805460019260049290918590811061102157fe5b9060005260206000200154815260200190815260200160002060030160006101000a81548160ff021916908315150217905550600460006006838154811061106557fe5b60009182526020808320909101548352820192909252604001902054600680546001600160a01b03909216918390811061109b57fe5b90600052602060002001547fc9d96d61eb62031865c523ae107f3c22f5ed445af237636bcd88bea1705c70d5426040518082815260200191505060405180910390a36111b2565b600160046000600684815481106110f557fe5b9060005260206000200154815260200190815260200160002060030160006101000a81548160ff021916908315150217905550600460006006838154811061113957fe5b60009182526020808320909101548352820192909252604001902054600680546001600160a01b03909216918390811061116f57fe5b90600052602060002001547fec955d77e6e7d74e18b1c91977ef0f6fd5a6d02a28d1979686339fe693997825426040518082815260200191505060405180910390a35b600101610e9c565b50506008805460ff1916905550565b6004602090815260009182526040918290208054600180830180548651600261010094831615949094026000190190911692909204601f81018690048602830186019096528582526001600160a01b039092169492939092908301828280156112735780601f1061124857610100808354040283529160200191611273565b820191906000526020600020905b81548152906001019060200180831161125657829003601f168201915b50505060028401546003850154600486015460058701546006909701549596929560ff92831695509093501687565b61f00481565b6206978081565b60085460ff16156112fe576040805162461bcd60e51b81526020600482015260146024820152731d9bdd19481a5cc81b9bdd08199a5b9a5cda195960621b604482015290519081900360640190fd5b6000818152600460205260409020546001600160a01b03163314611369576040805162461bcd60e51b815260206004820152601960248201527f796f752063616e2774207769746864726177616c20636f696e00000000000000604482015290519081900360640190fd5b60008181526004602052604090206006015460ff16156113d0576040805162461bcd60e51b815260206004820152601d60248201527f74686520636f696e20686164206265656e2077697468726177616c6564000000604482015290519081900360640190fd5b600081815260046020819052604080832060068101805460ff19166001179055909101549051339282156108fc02929190818181858888f1935050505015801561141e573d6000803e3d6000fd5b506000818152600460208181526040928390209091015482519081529151339284927f5448ccfcfb641b84b45c8fdcc6c5aa18fb824086d1293c111080bc87aa5763ad92918290030190a350565b60085460009060ff166114b4576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd081cdd185c9d195960aa1b604482015290519081900360640190fd5b3360009081526002602052604090205460ff16156115035760405162461bcd60e51b815260040180806020018281038252602981526020018061253c6029913960400191505060405180910390fd5b69d3c21bcecceda100000034101561154c5760405162461bcd60e51b815260040180806020018281038252602f81526020018061249e602f913960400191505060405180910390fd5b6000333385854260405160200180866001600160a01b031660601b8152601401856001600160a01b031660601b815260140184848082843791909101928352505060408051808303815260209283018252805190830120336000908152600990935291205490955060ff1615935061160992505050576040805162461bcd60e51b815260206004820152601760248201527650726f706f73616c20616c72656164792065786973747360481b604482015290519081900360640190fd5b610bb8831115611653576040805162461bcd60e51b815260206004820152601060248201526f44657461696c7320746f6f206c6f6e6760801b604482015290519081900360640190fd5b600081815260046020526040902060020154156116b1576040805162461bcd60e51b815260206004820152601760248201527650726f706f73616c20616c72656164792065786973747360481b604482015290519081900360640190fd5b6116b961238d565b338152604080516020601f87018190048102820181019092528581529086908690819084018382808284376000920182905250602086810195865242604080890191909152346080890181905260a089015288835260048252909120865181546001600160a01b0319166001600160a01b03909116178155945180518796955061174c94506001860193509101906123d7565b5060408201516002820155606082015160038201805491151560ff199283161790556080830151600483015560a0830151600583015560c09092015160069091018054911515919092161790556007546117a690346121be565b6007556006805460018181019092557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0183905533600081815260096020908152604091829020805460ff1916909417909355805142815290519192839286927fc10f2f4d53a0e342536c6af3cce9c6ee25c32dbb323521ce0e1d4494a3e362e892908290030190a46001925050505b92915050565b6006818154811061184957fe5b600091825260209091200154905081565b61f00681565b620d2f0081565b61f00781565b600a5481565b61a00281565b61f00381565b60026020526000908152604090205460ff1681565b69d3c21bcecceda100000081565b6b0d92289838d21a996800000081565b60075481565b6b015b6a759f4835dc2400000081565b600181565b60005460ff161561191b576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b6008805462f00400610100600160a81b0319909116179055600380546001600160a01b03191673d1a4e4fbc69e17cabb96fd0d45cf98889b85034717905560005b81811015611a2957600083838381811061197257fe5b905060200201356001600160a01b03166001600160a01b031614156119de576040805162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f72206164647265737300000000000000604482015290519081900360640190fd5b6001600260008585858181106119f057fe5b602090810292909201356001600160a01b0316835250810191909152604001600020805460ff191691151591909117905560010161195c565b505060088054600160ff19918216811790925542600a5560008054909116909117905550565b73d1a4e4fbc69e17cabb96fd0d45cf98889b85034781565b600a81565b683635c9adc5dea0000081565b6a0422ca8b0a00a42500000081565b6b07fdacf155df27a328c0000081565b60005461010090046001600160a01b031681565b603381565b600190565b603c81565b601e81565b6201518081565b60085460ff1615611b16576040805162461bcd60e51b81526020600482015260146024820152731d9bdd19481a5cc81b9bdd08199a5b9a5cda195960621b604482015290519081900360640190fd5b3360009081526005602090815260408083208484528252808320805482518185028101850190935280835260609492939192909184015b82821015611baf576000848152602090819020604080516080810182526004860290920180546001600160a01b03168352600180820154848601526002820154928401929092526003015460ff16151560608301529083529092019101611b4d565b50508251929350600091505081611c04576040805162461bcd60e51b81526020600482015260146024820152736e6f7468696e6720666f7220776974686472617760601b604482015290519081900360640190fd5b60005b82811015611c7757838181518110611c1b57fe5b602090810291909101015160600151611c6f576001848281518110611c3c57fe5b60200260200101516060019015159081151581525050838181518110611c5e57fe5b602002602001015160400151820191505b600101611c07565b506000848152600460209081526040808320546001600160a01b03168352600290915281205460ff1615611ccd57611cc682611cc06103e8611cba83601e61221f565b90612278565b906121be565b9050611ce5565b611ce282611cc06103e8611cba83600a61221f565b90505b60008111611d2f576040805162461bcd60e51b8152602060048201526012602482015271636f696e2063616e2774206265207a65726f60701b604482015290519081900360640190fd5b4781811015611d6f5760405162461bcd60e51b815260040180806020018281038252602a815260200180612586602a913960400191505060405180910390fd5b8115611ddb57604051339083156108fc029084906000818181858888f19350505050158015611da2573d6000803e3d6000fd5b50604080518381529051339188917f5448ccfcfb641b84b45c8fdcc6c5aa18fb824086d1293c111080bc87aa5763ad9181900360200190a35b505050505050565b60015481565b693ceab05409db274d269381565b60085460009060ff16611e3f576040805162461bcd60e51b815260206004820152600b60248201526a1b9bdd081cdd185c9d195960aa1b604482015290519081900360640190fd5b600082815260046020526040902060020154611e97576040805162461bcd60e51b8152602060048201526012602482015271141c9bdc1bdcd85b081b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b683635c9adc5dea00000341015611edf5760405162461bcd60e51b815260040180806020018281038252603b815260200180612501603b913960400191505060405180910390fd5b611ee7612455565b5060408051608081018252338082524260208084019182523484860181815260006060870181815295815260058085528882208b835285528882208054600180820183559184528684208a51600492830290910180546001600160a01b0319166001600160a01b039092169190911781559751918801919091559251600287015595516003909501805460ff19169515159590951790945588845290915293902001549091611f9691906121be565b600084815260046020526040902060050155600754611fb590346121be565b6007556040805142815290513491339186917f0b9c57abca55a6ad3059279543be201a6f54ed5374ef40fb92b456cad88b96ca919081900360200190a450600192915050565b6003546001600160a01b031681565b604080516028808252606082810190935282919060208201818036833701905050905060005b60148110156120f75760008160130360080260020a856001600160a01b03168161205657fe5b0460f81b9050600060108260f81c60ff168161206e57fe5b0460f81b905060008160f81c6010028360f81c0360f81b9050612090826122ba565b85856002028151811061209f57fe5b60200101906001600160f81b031916908160001a9053506120bf816122ba565b8585600202600101815181106120d157fe5b60200101906001600160f81b031916908160001a90535050600190920191506120309050565b5092915050565b60608083836040516020018083805190602001908083835b602083106121355780518252601f199092019160209182019101612116565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061217d5780518252601f19909201916020918201910161215e565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290508091505092915050565b600082820183811015612218576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b60008261222e57506000611836565b8282028284828161223b57fe5b04146122185760405162461bcd60e51b81526004018080602001828103825260218152602001806125656021913960400191505060405180910390fd5b600061221883836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506122eb565b6000600a60f883901c10156122da578160f81c60300160f81b9050610d02565b8160f81c60570160f81b9050610d02565b600081836123775760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561233c578181015183820152602001612324565b50505050905090810190601f1680156123695780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161238357fe5b0495945050505050565b6040518060e0016040528060006001600160a01b03168152602001606081526020016000815260200160001515815260200160008152602001600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061241857805160ff1916838001178555612445565b82800160010185558215612445579182015b8281111561244557825182559160200191906001019061242a565b50612451929150612488565b5090565b604051806080016040528060006001600160a01b0316815260200160008152602001600081526020016000151581525090565b5b80821115612451576000815560010161248956fe796f752063616e2774206372656174652070726f706f73616c2c20706c6561736520707574206d6f726520636f696e6e6f7420726561636865642074696d6520746f207472616e7366657220636f696e206261636b20746f20666f756e646174696f6e596f752063616e277420766f74652c20796f7520646f6e742720726561636820746865206d696e696d756d20766f746520636f696e20313030303044737420616c7265616479207061737365642c20596f752063616e207374617274207374616b696e67536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776e6f7420656e6f75676820636f696e206c6566742c20706c6561736520636f6e746163742061646d696ea2646970667358221220c335fd41e31e90e9159d54614d9ce4e8a20b5b35e45294122a01eff3df4b8ce064736f6c634300060c0033"
)

type forkProposalV2 struct {

}


func (f * forkProposalV2) GetName() string{
	return ProposalV2ContractName
}

func (f *forkProposalV2) Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) (err error) {
	contractCode := common.FromHex(proposalV2Code)

	//write code to sys contract
	state.SetCode(ProposalV2ContractAddr, contractCode)
	log.Debug("Write code to system contract account", "addr", ProposalV2ContractAddr.String(), "code", proposalV2Code)

	return
}

func (f *forkProposalV2) Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {

	v0 := NewValidatorV0()
	topVals, err := v0.GetTopValidators(state, header, chainContext, config)
	if err != nil {
		log.Error("getTopValidators from V0 failed", "err", err)
		return err
	}

	method := "initialize"
	data, err := GetInteractiveABI()[ProposalV2ContractName].Pack(method, topVals)
	if err != nil {
		log.Error("Can't pack data for initialize", "error", err)
		return err
	}

	msg := types.NewMessage(header.Coinbase, &ProposalV2ContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)
	_, err = caller.ExecuteMsg(msg, state, header, chainContext, config)

	return
}
