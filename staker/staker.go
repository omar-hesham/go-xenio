// Copyright 2017 The go-xenio Authors
//
// This file is part of the go-xenio library.
//
// The go-xenio library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-xenio library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-xenio library. If not, see <http://www.gnu.org/licenses/>.

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
	"math/big"
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

	watcher *watcher

	coinbase common.Address
	staking   int32
	eth      Backend
	engine   consensus.Engine

	canStart    int32 // can start indicates whether we can start the staking operation
	shouldStart int32 // should start indicates whether we should start after sync

	isConnected bool //indicates that the staker is connected to a server node, probably gaming
	serverbase common.Address

}


func New(eth Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine) *Staker {
	staker := &Staker{
		eth:      eth,
		mux:      mux,
		engine:   engine,
		watcher:   newWatcher(config, engine, common.Address{}, eth, mux),
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
	self.watcher.setEtherbase(coinbase)
	self.coinbase = coinbase

	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start staker afterwards")
		return
	}

	atomic.StoreInt32(&self.staking, 1)

	log.Info("Starting staking operation")

	self.watcher.start()
	self.watcher.commitNewWork()
}


func (self *Staker) Stop() {
	self.watcher.stop()
	atomic.StoreInt32(&self.staking, 0)
	atomic.StoreInt32(&self.shouldStart, 0)
}

func (self *Staker) Register(agent Agent) {
	if self.Staking() {
		agent.Start()
	}
	self.watcher.register(agent)
}

func (self *Staker) Unregister(agent Agent) {
	self.watcher.unregister(agent)
}

func (self *Staker) Staking() bool {
	return atomic.LoadInt32(&self.staking) > 0
}


func (self *Staker) SetExtra(extra []byte) error {
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("Extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
	}
	self.watcher.setExtra(extra)
	return nil
}

// Pending returns the currently pending block and associated state.
func (self *Staker) Pending() (*types.Block, *state.StateDB) {
	return self.watcher.pending()
}

// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (self *Staker) PendingBlock() *types.Block {
	return self.watcher.pendingBlock()
}

func (self *Staker) SetEtherbase(addr common.Address) {
	self.coinbase = addr
	self.watcher.setEtherbase(addr)
}
func(self *Staker) SetServerbase(addr common.Address) {//TODO: should be changed to a server node
	self.serverbase = addr
}