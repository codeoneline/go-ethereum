// Copyright 2016 The go-ethereum Authors
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

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const MainnetPow2PosUpgradeBlockNumber = 4046000
const TestnetPow2PosUpgradeBlockNumber = 3560000
const InternalPow2PosUpgradeBlockNumber = 200

const MAINNET_CHAIN_ID = 1
const TESTNET_CHAIN_ID = 3

// INTERNAL_CHAIN_ID is Private chain pow -> pos mode chainId
const INTERNAL_CHAIN_ID = 4

// PLUTO_CHAIN_ID is Private chain pos mode chainId
const PLUTO_CHAIN_ID = 6

// PLUTODEV_CHAIN_ID is Private chain pos mode single node chainId
const PLUTODEV_CHAIN_ID = 6

// JUPITER_MAINNET_CHAIN_ID is mainnet chainId after jupiter fork
const JUPITER_MAINNET_CHAIN_ID = 888

// JUPITER_TESTNET_CHAIN_ID is testnet chainId after jupiter fork
const JUPITER_TESTNET_CHAIN_ID = 999
const JUPITER_INTERNAL_CHAIN_ID = 777
//const JUPITER_PLUTO_CHAIN_ID = 6
const JUPITER_PLUTO_CHAIN_ID = 666
//const JUPITER_PLUTODEV_CHAIN_ID = 6
const JUPITER_PLUTODEV_CHAIN_ID = 555

// NOT_JUPITER_CHAIN_ID is used for compare
const NOT_JUPITER_CHAIN_ID = 0xffffffff


// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash = common.HexToHash("0x0376899c001618fc7d5ab4f31cfd7f57ca3a896ccc1581a57d8f129ecf40b840") // Mainnet genesis hash to enforce below configs on
	TestnetGenesisHash = common.HexToHash("0xa37b811609a9d1e898fb49b3901728023e5e72e18e58643d9a7a82db483bfeb0") // Testnet genesis hash to enforce below configs on
	PlutoGenesisHash   = common.HexToHash("0x7b67a3f28e0d12b57e5fdaa445c4d6dbe68bffa9b808e944e5c67726669d62b6") // Pluto genesis hash to enforce below configs on
	InternalGenesisHash = common.HexToHash("0xb1dc31a86510003c23b9ddee0e194775807262529b8dafa6dc23d9315364d2b3")
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID:             big.NewInt(1),
		HomesteadBlock:      big.NewInt(0),
		DAOForkBlock:        big.NewInt(0),
		DAOForkSupport:      false,
		EIP150Block:         big.NewInt(0),
		EIP150Hash:          common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:         big.NewInt(0),
		EIP158Block:         big.NewInt(0),
		ByzantiumBlock:      big.NewInt(0),
		ConstantinopleBlock: big.NewInt(999_000_000_000),
		PetersburgBlock:     big.NewInt(999_000_000_000),
		IstanbulBlock:       big.NewInt(999_000_000_000),
		MuirGlacierBlock:    big.NewInt(999_000_000_000),
		BerlinBlock:         big.NewInt(999_000_000_000),
		Ethash:              new(EthashConfig),
	}
	/*

		WanchainChainConfig = &ChainConfig{
			ChainID: big.NewInt(1),
			//HomesteadBlock: big.NewInt(0),
			//DAOForkBlock:   nil,
			//DAOForkSupport: true,
			//EIP150Block:    big.NewInt(0),
			//EIP150Hash:     common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
			//EIP155Block:    big.NewInt(0),
			//EIP158Block:    big.NewInt(0),
			ByzantiumBlock: big.NewInt(0),
			Ethash:         new(EthashConfig),
		}

		// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
		TestnetChainConfig = &ChainConfig{
			ChainID: big.NewInt(3),
			//HomesteadBlock: big.NewInt(0),
			//DAOForkBlock:   nil,
			//DAOForkSupport: true,
			//EIP150Block:    big.NewInt(0),
			//EIP150Hash:     common.HexToHash("0x41941023680923e0fe4d74a34bdac8141f2540e3ae90623718e47d66d1ca4a2d"),
			//EIP155Block:    big.NewInt(10),
			//EIP158Block:    big.NewInt(10),
			ByzantiumBlock: big.NewInt(0),

			Ethash: new(EthashConfig),
		}

		// RinkebyChainConfig contains the chain parameters to run a node on the Rinkeby test network.
		InternalChainConfig = &ChainConfig{
			ChainID: big.NewInt(4),
			//HomesteadBlock: big.NewInt(1),
			//DAOForkBlock:   nil,
			//DAOForkSupport: true,
			//EIP150Block:    big.NewInt(2),
			//EIP150Hash:     common.HexToHash("0x9b095b36c15eaf13044373aef8ee0bd3a382a5abb92e402afa44b8249c3a90e9"),
			//EIP155Block:    big.NewInt(3),
			//EIP158Block:    big.NewInt(3),
			ByzantiumBlock: big.NewInt(0),

			Ethash: new(EthashConfig),
		}
		// PlutoChainConfig contains the chain parameters to run a node on the Pluto test network.
		PlutoChainConfig = &ChainConfig{
			ChainID: big.NewInt(6),
			//HomesteadBlock: big.NewInt(0),
			//DAOForkBlock:   nil,
			//DAOForkSupport: true,
			//EIP150Block:    big.NewInt(0),
			//EIP150Hash:     common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
			//EIP155Block:    big.NewInt(3),
			//EIP158Block:    big.NewInt(3),
			ByzantiumBlock: big.NewInt(0),

			Pluto: &PlutoConfig{
				Period: 6,
				Epoch:  6000,
			},
		}
	*/

	// AllProtocolChanges contains every protocol change (EIPs)
	// introduced and accepted by the Ethereum core developers.
	//
	// This configuration is intentionally not using keyed fields.
	// This configuration must *always* have all forks enabled, which
	// means that all fields must be set at all times. This forces
	// anyone adding flags to the config to also have to set these
	// fields.
	//AllProtocolChanges = &ChainConfig{big.NewInt(1337), big.NewInt(0), big.NewInt(100), false, new(EthashConfig), nil, nil}
	//
	//TestChainConfig = &ChainConfig{
	//	ChainId:        big.NewInt(MAINNET_CHAIN_ID),
	//	ByzantiumBlock: big.NewInt(0),
	//	Ethash:         new(EthashConfig),
	//	PosFirstBlock:  big.NewInt(TestnetPow2PosUpgradeBlockNumber), // set as n * epoch_length
	//	IsPosActive:    false,
	//}
	//
	//TestRules = TestChainConfig.Rules(new(big.Int))

	noStaking = false
)


func IsNoStaking() bool {
	return noStaking
}
