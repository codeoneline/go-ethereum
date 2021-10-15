package eth

import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

// StatusPacket is the network packet for the status message for eth/64 and later.
type StatusData struct {
	ProtocolVersion uint32
	NetworkID       uint64
	TD              *big.Int
	Head            common.Hash
	Genesis         common.Hash
}
