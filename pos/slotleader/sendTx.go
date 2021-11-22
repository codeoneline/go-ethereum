package slotleader

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	errRCNotReady = errors.New("rc is not ready")
)

type SendTxFn func(rc *rpc.Client, tx map[string]interface{})

func (s *SLS) sendSlotTx(payload []byte, posSender SendTxFn) error {
	if s.rc == nil {
		return errRCNotReady
	}

	to := vm.GetSlotLeaderSCAddress()
	data := hexutil.Bytes(payload)
	//gas := core.IntrinsicGas(data, &to, true)
	gas := core.IntrinsicGas_gwan(data, &to, true)

	arg := map[string]interface{}{}
	arg["from"] = s.key.Address
	arg["to"] = vm.GetSlotLeaderSCAddress()
	arg["value"] = (*hexutil.Big)(big.NewInt(0))
	arg["gas"] = (*hexutil.Big)(gas)
	arg["txType"] = types.POS_TX
	arg["data"] = data
	log.Debug("Write data of payload", "length", len(data))

	go posSender(s.rc, arg)
	return nil
}
