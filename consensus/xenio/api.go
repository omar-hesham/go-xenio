package xenio

import (
	"github.com/xenioplatform/go-xenio/consensus"
)

type API struct {
	chain  consensus.ChainReader
	xenio *Xenio
}



