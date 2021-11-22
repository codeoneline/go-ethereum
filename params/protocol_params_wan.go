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

const (
	// ppow
	MaximumExtraDataSize uint64 = 97 // Maximum size extra data may be after Genesis.

	// builtin bn256
	ModExpQuadCoeffDiv      uint64 = 20     // Divisor for the quadratic particle of the big int modular exponentiation
	Bn256AddGas             uint64 = 500    // Gas needed for an elliptic curve addition
	Bn256ScalarMulGas       uint64 = 40000  // Gas needed for an elliptic curve scalar multiplication
	Bn256PairingBaseGas     uint64 = 100000 // Base price for an elliptic curve pairing check
	Bn256PairingPerPointGas uint64 = 80000  // Per-point price for an elliptic curve pairing check
	Bn256AddGasV2           uint64 = 50     // Gas needed for an elliptic curve addition
	Bn256ScalarMulGasV2     uint64 = 100
	RequiredGasPerMixPub    uint64 = 4000 // ring signature mix difficulty gas
	GetOTAMixSetMaxSize     uint64 = 20   // Max number of mix ota set size from once getting
	GasForSolEnhance        uint64 = 100
	S256AddGas              uint64 = 50 // Gas needed for an elliptic curve addition
	S256ScalarMulGas        uint64 = 100

	// private transaction.
	//SlsStgOnePerByteGas		uint64 = 20      // per byte gas for SlsStgOnePerByteGas
	SlsStgTwoPerByteGas uint64 = 20 // per byte gas for SlsStgOnePerByteGas

)
