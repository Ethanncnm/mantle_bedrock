// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/mantlenetworkio/mantle/mt-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1020_storage\"},{\"astId\":1004,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1006,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1008,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap_reentrancy_guard\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)1019_storage\"},{\"astId\":1010,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"reentrancyLocks\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1017,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)1018_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1018_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[41]\",\"numberOfBytes\":\"1312\"},\"t_array(t_uint256)1019_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1020_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106100f35760003560e01c80638129fc1c1161008a578063b28ade2511610059578063b28ade25146102e1578063d764ad0b14610301578063da09a23214610314578063ecc704281461032757600080fd5b80638129fc1c146102265780639fce812c1461023d578063a4e7f8bd14610271578063b1b1b209146102b157600080fd5b80633f827a5a116100c65780633f827a5a146101b057806354fd4d50146101d85780636e296e45146101fa5780637dea7cc31461020f57600080fd5b8063028f85f7146100f85780630c5684981461012b5780630ff754ea146101415780632828d7e81461019a575b600080fd5b34801561010457600080fd5b5061010d601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b34801561013757600080fd5b5061010d6103e881565b34801561014d57600080fd5b506101757f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610122565b3480156101a657600080fd5b5061010d6103f881565b3480156101bc57600080fd5b506101c5600181565b60405161ffff9091168152602001610122565b3480156101e457600080fd5b506101ed61038c565b6040516101229190611b61565b34801561020657600080fd5b5061017561042f565b34801561021b57600080fd5b5061010d62030d4081565b34801561023257600080fd5b5061023b61051b565b005b34801561024957600080fd5b506101757f000000000000000000000000000000000000000000000000000000000000000081565b34801561027d57600080fd5b506102a161028c366004611b7b565b60ce6020526000908152604090205460ff1681565b6040519015158152602001610122565b3480156102bd57600080fd5b506102a16102cc366004611b7b565b60cb6020526000908152604090205460ff1681565b3480156102ed57600080fd5b5061010d6102fc366004611bf6565b610718565b61023b61030f366004611c6c565b610764565b61023b610322366004611cf2565b611018565b34801561033357600080fd5b5061037e60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b604051908152602001610122565b60606103b77f0000000000000000000000000000000000000000000000000000000000000000611412565b6103e07f0000000000000000000000000000000000000000000000000000000000000000611412565b6104097f0000000000000000000000000000000000000000000000000000000000000000611412565b60405160200161041b93929190611d72565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff2153016104fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000547501000000000000000000000000000000000000000000900460ff1615808015610566575060005460017401000000000000000000000000000000000000000090910460ff16105b806105985750303b158015610598575060005474010000000000000000000000000000000000000000900460ff166001145b610624576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016104f5565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000017905580156106aa57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b6106b2611547565b801561071557600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50565b600062030d40610729601085611e17565b6103e861073e6103f863ffffffff8716611e17565b6107489190611e76565b6107529190611e9d565b61075c9190611e9d565b949350505050565b60f087901c6002811061081f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a4016104f5565b8061ffff16600003610914576000610870878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611620915050565b600081815260cb602052604090205490915060ff1615610912576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c6179656400000000000000000060648201526084016104f5565b505b600061095a898989898989898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061163f92505050565b600081815260cf602052604090205490915060ff16156109d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c0060448201526064016104f5565b600081815260cf6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055610a15611662565b15610a4d57853414610a2957610a29611ec9565b600081815260ce602052604090205460ff1615610a4857610a48611ec9565b610b9f565b3415610b01576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a4016104f5565b600081815260ce602052604090205460ff16610b9f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c617965640000000000000000000000000000000060648201526084016104f5565b610ba887611786565b15610c5b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a4016104f5565b600081815260cb602052604090205460ff1615610cfa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c617965640000000000000000000060648201526084016104f5565b610d0661afc886611ef8565b5a1015610d95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d65737361676500000000000000000060648201526084016104f5565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a161790556000610e3188610de961138861afc8611f10565b5a610df49190611f10565b8988888080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117fd92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801515600103610ecc57600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2610fd9565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201610fd9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016104f5565b50600090815260cf6020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555050505050505050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff8716016111845761117f867f0000000000000000000000000000000000000000000000000000000000000000611075868686610718565b887fd764ad0b000000000000000000000000000000000000000000000000000000006110e160cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338b8d8a8d8d6040516024016110fd9796959493929190611f70565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611817565b61125b565b63ffffffff8616158061119d575063ffffffff86166002145b1561125b5761125b867f00000000000000000000000000000000000000000000000000000000000000006111d2868686610718565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061123e60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338b60008a8d8d6040516024016110fd9796959493929190611fcf565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856112e060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516112f2959493929190612024565b60405180910390a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff8716016113615760405185815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2611398565b6040516000815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a25b505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff00000000000000000000000000000000000000000000000000000000000090911617905550505050565b905090565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60608160000361145557505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561147f578061146981612072565b91506114789050600a836120aa565b9150611459565b60008167ffffffffffffffff81111561149a5761149a6120be565b6040519080825280601f01601f1916602001820160405280156114c4576020820181803683370190505b5090505b841561075c576114d9600183611f10565b91506114e6600a866120ed565b6114f1906030611ef8565b60f81b81838151811061150657611506612101565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611540600a866120aa565b94506114c8565b6000547501000000000000000000000000000000000000000000900460ff166115f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016104f5565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b600061162e858585856119af565b805190602001209050949350505050565b600061164f878787878787611a48565b8051906020012090509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156113f157507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639bf62d826040518163ffffffff1660e01b8152600401602060405180830381865afa158015611746573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061176a9190612130565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff82163014806117f757507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b92915050565b600080600080845160208601878a8af19695505050505050565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff63ffffffff8616016118f8576040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e9e05c42906000906118c1908890879089908590899060040161214d565b6000604051808303818588803b1580156118da57600080fd5b505af11580156118ee573d6000803e3d6000fd5b50505050506119a8565b6040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e9e05c4290849061197590889060009089908290899060040161214d565b6000604051808303818588803b15801561198e57600080fd5b505af11580156119a2573d6000803e3d6000fd5b50505050505b5050505050565b6060848484846040516024016119c894939291906121a5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b6060868686868686604051602401611a65969594939291906121ef565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b60005b83811015611b02578181015183820152602001611aea565b83811115611b11576000848401525b50505050565b60008151808452611b2f816020860160208601611ae7565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611b746020830184611b17565b9392505050565b600060208284031215611b8d57600080fd5b5035919050565b60008083601f840112611ba657600080fd5b50813567ffffffffffffffff811115611bbe57600080fd5b602083019150836020828501011115611bd657600080fd5b9250929050565b803563ffffffff81168114611bf157600080fd5b919050565b600080600060408486031215611c0b57600080fd5b833567ffffffffffffffff811115611c2257600080fd5b611c2e86828701611b94565b9094509250611c41905060208501611bdd565b90509250925092565b73ffffffffffffffffffffffffffffffffffffffff8116811461071557600080fd5b600080600080600080600060c0888a031215611c8757600080fd5b873596506020880135611c9981611c4a565b95506040880135611ca981611c4a565b9450606088013593506080880135925060a088013567ffffffffffffffff811115611cd357600080fd5b611cdf8a828b01611b94565b989b979a50959850939692959293505050565b60008060008060008060a08789031215611d0b57600080fd5b611d1487611bdd565b9550602087013594506040870135611d2b81611c4a565b9350606087013567ffffffffffffffff811115611d4757600080fd5b611d5389828a01611b94565b9094509250611d66905060808801611bdd565b90509295509295509295565b60008451611d84818460208901611ae7565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611dc0816001850160208a01611ae7565b60019201918201528351611ddb816002840160208801611ae7565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615611e3e57611e3e611de8565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff80841680611e9157611e91611e47565b92169190910492915050565b600067ffffffffffffffff808316818516808303821115611ec057611ec0611de8565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60008219821115611f0b57611f0b611de8565b500190565b600082821015611f2257611f22611de8565b500390565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152611fc260c083018486611f27565b9998505050505050505050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525060ff8616606083015263ffffffff8516608083015260c060a0830152611fc260c083018486611f27565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000612054608083018688611f27565b905083604083015263ffffffff831660608301529695505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036120a3576120a3611de8565b5060010190565b6000826120b9576120b9611e47565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826120fc576120fc611e47565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006020828403121561214257600080fd5b8151611b7481611c4a565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a06080820152600061219a60a0830184611b17565b979650505050505050565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250608060408301526121de6080830185611b17565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261223a60c0830184611b17565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}