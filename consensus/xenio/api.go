package xenio

import (
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/consensus"
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

