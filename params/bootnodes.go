// Copyright 2015 The go-rue Authors
// This file is part of the go-rue library.
//
// The go-rue library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-rue library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-rue library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Rue network.
var MainnetBootnodes = []string{
	// Rue Foundation Go Bootnodes
<<<<<<< HEAD
	"enode://5538ada83435de5935040f80e1f9e8297cda04f7b213d7a8c45e2550d401a90df21653586ee60ebb58b762a56ee0334ceb710821cd450fd731b36b877dc3ff04@173.33.8.90:30304",
=======
	"enode://a979fb575495b8d6db44f750317d0f4622bf4c2aa3365d6af7c284339968eef29b69ad0dce72a4d8db5ebb4968de0e3bec910127f134779fbcb0cb6d3331163c@52.16.188.185:30303", // IE
	"enode://3f1d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@13.93.211.84:30303",  // US-WEST
	"enode://78de8a0916848093c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@191.235.84.50:30303", // BR
	"enode://158f8aab45f6d19c6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@13.75.154.138:30303", // AU
	"enode://1118980bf48b0a3640bdba04e0fe78b1add18e1cd99bf22d53daac1fd9972ad650df52176e7c7d89d1114cfef2bc23a2959aa54998a46afcf7d91809f0855082@52.74.57.123:30303",  // SG

	// Rue Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
>>>>>>> e30d8edd0de13be18f8ba77cb45dd7ccd9cb1ebf
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://5538ada83435de5935040f80e1f9e8297cda04f7b213d7a8c45e2550d401a90df21653586ee60ebb58b762a56ee0334ceb710821cd450fd731b36b877dc3ff04@173.33.8.90:30304",
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://5538ada83435de5935040f80e1f9e8297cda04f7b213d7a8c45e2550d401a90df21653586ee60ebb58b762a56ee0334ceb710821cd450fd731b36b877dc3ff04@173.33.8.90:30304",
}

// RinkebyV5Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network for the experimental RLPx v5 topic-discovery network.
var RinkebyV5Bootnodes = []string{
	"enode://5538ada83435de5935040f80e1f9e8297cda04f7b213d7a8c45e2550d401a90df21653586ee60ebb58b762a56ee0334ceb710821cd450fd731b36b877dc3ff04@173.33.8.90:30304",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://5538ada83435de5935040f80e1f9e8297cda04f7b213d7a8c45e2550d401a90df21653586ee60ebb58b762a56ee0334ceb710821cd450fd731b36b877dc3ff04@173.33.8.90:30304",
}
