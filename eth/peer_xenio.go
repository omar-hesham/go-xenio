package eth

import (
	"fmt"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/consensus/xenio"
	"github.com/xenioplatform/go-xenio/p2p"
	//"github.com/xenioplatform/go-xenio/log"
	"time"
)
const (
	stakeInterval = 120 * time.Second // 15 for dev tests
)
// RequestCoinbase fetches the coinbase of the remote node, if the node wants to share it.
func (p *peer) RequestCoinbase(adr common.Address) error {
	p.Log().Warn("Fetching remote nodes Coinbase")
	return p2p.Send(p.rw, GetNodeCoinbase, adr)
}

// TransmitCoinbase transmits our coinbase to the remote node.
func (p *peer) TransmitCoinbase(adr common.Address) error {
	p.Log().Warn("Transmitting Coinbase")
	err := p2p.Send(p.rw, TransmitCoinbase, adr)
	if err == nil {
		err = p.TransmitNodeList()
	}
	return err
}
func (p *peer) TransmitNodeList() error {
	p.Log().Warn("Transmitting NodeList")
	if common.StakerSnapShot != nil && len(common.StakerSnapShot.Stakers) > 0 {
		var toSend []common.StakerTransmit
		for key, value := range common.StakerSnapShot.Stakers {
			if xenio.StakerExpired(key) == false {
				_time := fmt.Sprint(value.LastSeen.Unix())
				toSend = append(toSend, common.StakerTransmit{key, _time})
			}
		}
		return p2p.Send(p.rw, TransmitNodeList, toSend)
	} else {
		return nil
	}
}

func (p *peer) transmitLoop() {
	ping := time.NewTicker(stakeInterval)
	defer p.stakeWait.Done()
	defer ping.Stop()
	for {
		select {
		case <-ping.C:
			if err := p.TransmitCoinbase(common.Address{}); err != nil {
				p.Peer.ProtoErr <- err
				return
			}
		case <-p.Peer.Closed:
			return
		}
	}
}
