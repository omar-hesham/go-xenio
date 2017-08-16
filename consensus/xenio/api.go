package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/consensus"
	"github.com/xenioplatform/go-xenio/log"

	//"github.com/xenioplatform/go-xenio/core/types"
	//"github.com/xenioplatform/go-xenio/rpc"
)

type API struct {
	chain  consensus.ChainReader
	xenio *Xenio
}

// Discard drops a currently running staker.
func (api *API) Discard(address common.Address) {
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()

	delete(api.xenio.stakers, address)
}

// Propose injects a new authorization proposal that the signer will attempt to
// push through.
func (api *API) Propose(address common.Address, auth bool) {
	log.Warn("Propose: propose")
	api.xenio.lock.Lock()
	defer api.xenio.lock.Unlock()

	api.xenio.proposals[address] = auth
}

func (api *API) Proposals() map[common.Address]bool {
	api.xenio.lock.RLock()
	defer api.xenio.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range api.xenio.proposals {
		proposals[address] = auth
	}
	return proposals
}