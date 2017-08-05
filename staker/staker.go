// Copyright 2014 The go-xenio Authors
// This file is part of the go-xenio library.

package staker

import (
	"fmt"
	"sync/atomic"

	"github.com/xenioplatform/go-xenio/accounts"
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/consensus"
	"github.com/xenioplatform/go-xenio/core"
	"github.com/xenioplatform/go-xenio/core/state"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/eth/downloader"
	"github.com/xenioplatform/go-xenio/ethdb"
	"github.com/xenioplatform/go-xenio/event"
	"github.com/xenioplatform/go-xenio/log"
	"github.com/xenioplatform/go-xenio/params"
)

// Backend wraps all methods required for mining.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() ethdb.Database
}

type Staker struct {
	mux *event.TypeMux

	worker *worker

	coinbase common.Address
	staking   int32
	eth      Backend
	engine   consensus.Engine

	canStart    int32 // can start indicates whether we can start the mining operation
	shouldStart int32 // should start indicates whether we should start after sync
}

func New(eth Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine) *Staker {
	staker := &Staker{
		eth:      eth,
		mux:      mux,
		engine:   engine,
		worker:   newWorker(config, engine, common.Address{}, eth, mux),
		canStart: 1,
	}
	staker.Register(NewCpuAgent(eth.BlockChain(), engine))
	go staker.update()

	return staker
}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (self *Staker) update() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
out:
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&self.canStart, 0)
			if self.Staking() {
				self.Stop()
				atomic.StoreInt32(&self.shouldStart, 1)
				log.Info("Staking aborted due to sync")
			}
		case downloader.DoneEvent, downloader.FailedEvent:
			shouldStart := atomic.LoadInt32(&self.shouldStart) == 1

			atomic.StoreInt32(&self.canStart, 1)
			atomic.StoreInt32(&self.shouldStart, 0)
			if shouldStart {
				self.Start(self.coinbase)
			}
			// unsubscribe. we're only interested in this event once
			events.Unsubscribe()
			// stop immediately and ignore all further pending events
			break out
		}
	}
}


func (self *Staker) Start(coinbase common.Address) {
	atomic.StoreInt32(&self.shouldStart, 1)
	self.worker.setEtherbase(coinbase)
	self.coinbase = coinbase

	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start staker afterwards")
		return
	}
	atomic.StoreInt32(&self.staking, 1)

	log.Info("Starting staking operation")
	self.worker.start()
	self.worker.commitNewWork()
}


func (self *Staker) Stop() {
	self.worker.stop()
	atomic.StoreInt32(&self.staking, 0)
	atomic.StoreInt32(&self.shouldStart, 0)
}

func (self *Staker) Register(agent Agent) {
	if self.Staking() {
		agent.Start()
	}
	self.worker.register(agent)
}

func (self *Staker) Unregister(agent Agent) {
	self.worker.unregister(agent)
}

func (self *Staker) Staking() bool {
	return atomic.LoadInt32(&self.staking) > 0
}

/*func (self *Staker) HashRate() (tot int64) {
	if pow, ok := self.engine.(consensus.PoW); ok {
		tot += int64(pow.Hashrate())
	}
	// do we care this might race? is it worth we're rewriting some
	// aspects of the worker/locking up agents so we can get an accurate
	// hashrate?
	for agent := range self.worker.agents {
		if _, ok := agent.(*CpuAgent); !ok {
			tot += agent.GetHashRate()
		}
	}
	return
}*/

func (self *Staker) SetExtra(extra []byte) error {
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("Extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
	}
	self.worker.setExtra(extra)
	return nil
}

// Pending returns the currently pending block and associated state.
func (self *Staker) Pending() (*types.Block, *state.StateDB) {
	return self.worker.pending()
}

// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (self *Staker) PendingBlock() *types.Block {
	return self.worker.pendingBlock()
}

func (self *Staker) SetEtherbase(addr common.Address) {
	self.coinbase = addr
	self.worker.setEtherbase(addr)
}
