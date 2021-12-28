// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://3acbbc8a4aff5815b3dc7a199e14d5291986406b1a4cd460aadd02107d637e4f53c4e90ad74c42084831163b40ef25b95cdf802f3b298ee8bd0193ebcef802bf@46.166.162.25:32668",
	"enode://be19b1e071eb59e60caefaf653438c450cc1a060ae47a9c94b64d07e366d2d8f968560ac1921816538ffe7b2c504f1540a8fb80d85075b0a4fe44b6badf85162@27.122.57.180:32668",
	"enode://f78284ea1237a057896fca306c28c16698d193431962b881012e4f95f005a776db1599cecbfca0276300d2a4dac0c348df02cd007ab05121a9f30c67ac3cd782@198.252.110.8:32668",
	"enode://a8ccb632b3777cb01f08bb46f0682ce19f6be7c023aa3ebad2c2a9cc7801e355f1ec2151229575c500e0d1350f9b9a96674df92db5ed86dccd55aae2ab800b25@35.228.209.216:32668",
	"enode://3910d8a659413cc2161813e0e80dd6bae737048ea76ea235d48f5a719fc4ea067c77b8d9bfabc0d143d3c8d9c5b520c57ce3583412afec23681f76f99efa439c@35.246.44.242:32668",
	"enode://216a406c9765733967f5b7db2716c599854d50c112509a6b988af874b891ab3501ffd3e8a1837a18f548be79c07dcc556fd9763dc0dfff7b01d653c43ea4b0e3@18.134.226.222:32668",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var TestnetBootnodes = []string{
	// "enode://924543a43d18bc5759a8bdcd17fa9c7c35df63968e9333640b80b58dab94b17a012371c9d46bed10ce7508a607cac76828ca04685893958eee44ade83b856dc2@47.242.237.63:32668",
	// "enode://ebad898d980b520ef6adb54ffb6a68117686e7332f1ea01f7551b7a296a34dd945445a078d7cad019d864c5ef0e0b7f2b5777d94f93adf7dc59f798af72609ac@47.242.235.121:32668",
}

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	return ""
}
