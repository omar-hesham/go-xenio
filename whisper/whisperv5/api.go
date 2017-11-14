// Copyright 2017 The go-xenio Authors
// Copyright 2016 The go-ethereum Authors
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

package whisperv5

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"sync"
	"time"
	"strings"
	//mrand "math/rand"
	"encoding/hex"

	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/common/hexutil"
	"github.com/xenioplatform/go-xenio/crypto"
	"github.com/xenioplatform/go-xenio/log"
	"github.com/xenioplatform/go-xenio/p2p/discover"
	"github.com/xenioplatform/go-xenio/rpc"
)

const (
	filterTimeout = 300 // filters are considered timeout out after filterTimeout seconds
)

var (
	ErrSymAsym              = errors.New("specify either a symmetric or an asymmetric key")
	ErrInvalidSymmetricKey  = errors.New("invalid symmetric key")
	ErrInvalidPublicKey     = errors.New("invalid public key")
	ErrInvalidSigningPubKey = errors.New("invalid signing public key")
	ErrTooLowPoW            = errors.New("message rejected, PoW too low")
	ErrNoTopics             = errors.New("missing topic(s)")
)

// PublicWhisperAPI provides the whisper RPC service that can be
// use publicly without security implications.
type PublicWhisperAPI struct {
	w *Whisper

	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.
}

// NewPublicWhisperAPI create a new RPC whisper service.
func NewPublicWhisperAPI(w *Whisper) *PublicWhisperAPI {
	api := &PublicWhisperAPI{
		w:        w,
		lastUsed: make(map[string]time.Time),
	}

	go api.run()
	return api
}

// run the api event loop.
// this loop deletes filter that have not been used within filterTimeout
func (api *PublicWhisperAPI) run() {
	timeout := time.NewTicker(2 * time.Minute)
	for {
		<-timeout.C

		api.mu.Lock()
		for id, lastUsed := range api.lastUsed {
			if time.Since(lastUsed).Seconds() >= filterTimeout {
				delete(api.lastUsed, id)
				if err := api.w.Unsubscribe(id); err != nil {
					log.Error("could not unsubscribe whisper filter", "error", err)
				}
				log.Debug("delete whisper filter (timeout)", "id", id)
			}
		}
		api.mu.Unlock()
	}
}

// Version returns the Whisper sub-protocol version.
func (api *PublicWhisperAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

// Info contains diagnostic information.
type Info struct {
	Memory         int     `json:"memory"`         // Memory size of the floating messages in bytes.
	Messages       int     `json:"messages"`       // Number of floating messages.
	MinPow         float64 `json:"minPow"`         // Minimal accepted PoW
	MaxMessageSize uint32  `json:"maxMessageSize"` // Maximum accepted message size
}

// Info returns diagnostic information about the whisper node.
func (api *PublicWhisperAPI) Info(ctx context.Context) Info {
	stats := api.w.Stats()
	return Info{
		Memory:         stats.memoryUsed,
		Messages:       len(api.w.messageQueue) + len(api.w.p2pMsgQueue),
		MinPow:         api.w.MinPow(),
		MaxMessageSize: api.w.MaxMessageSize(),
	}
}

// SetMaxMessageSize sets the maximum message size that is accepted.
// Upper limit is defined in whisperv5.MaxMessageSize.
func (api *PublicWhisperAPI) SetMaxMessageSize(ctx context.Context, size uint32) (bool, error) {
	return true, api.w.SetMaxMessageSize(size)
}

// SetMinPow sets the minimum PoW for a message before it is accepted.
func (api *PublicWhisperAPI) SetMinPoW(ctx context.Context, pow float64) (bool, error) {
	return true, api.w.SetMinimumPoW(pow)
}

// MarkTrustedPeer marks a peer trusted. , which will allow it to send historic (expired) messages.
// Note: This function is not adding new nodes, the node needs to exists as a peer.
func (api *PublicWhisperAPI) MarkTrustedPeer(ctx context.Context, enode string) (bool, error) {
	n, err := discover.ParseNode(enode)
	if err != nil {
		return false, err
	}
	return true, api.w.AllowP2PMessagesFromPeer(n.ID[:])
}

// NewKeyPair generates a new public and private key pair for message decryption and encryption.
// It returns an ID that can be used to refer to the keypair.
func (api *PublicWhisperAPI) NewKeyPair(ctx context.Context) (string, error) {
	return api.w.NewKeyPair()
}


func (api *PublicWhisperAPI) MessageSend(ctx context.Context,  message string,  topic string, hexKey string, targetPeer string) (string, error) {

	//symKey_id, err := api.w.GenerateSymKey()
	//symKey, err := api.w.GetSymKey(symKey_id)
	//"0x0f99d381030a59edc477bddf741069e0d9158c191f04f3d8d079c01db9792ee2"
	//var randomKey []byte //make([]byte, aesKeyLength)
	if hexKey == "" {
		hexKey = "7c8d019192c24224e2cafccae3a61fb586b14323a6bc8f9e7df1d929333ff993"
	}

	randomKey, err := hex.DecodeString(hexKey)
	var _message NewMessage
	_message.Payload = []byte(message)
	_message.PowTarget = 2
	_message.PowTime = 3
	_message.TTL = 600
	_message.Topic, err = api.MessageGenerateTopic(ctx, topic)

	if targetPeer != ""{
		_message.TargetPeer = targetPeer
	}
	log.Info("key: " + hex.EncodeToString(randomKey))
	log.Info("topic: " + _message.Topic.String())
	symKey_id, err := api.AddSymKey(ctx, randomKey)
	symKey, err := api.w.GetSymKey(symKey_id)

	log.Info("sumkeyid: " + string(symKey_id))
	//log.Info(string(symKey))
	_message.SymKeyID = symKey_id


	// post example
	//shh.sendMessage('boohoo','tooopic', '')
	_, err = api.Post(ctx, _message)

	return hex.EncodeToString(symKey), err
}

func (api *PublicWhisperAPI) MessageRead(ctx context.Context, messages []*Message) {
	if len(messages) > 0 {
		for _, msg := range messages {
			msg.ActualMessage = BytesToString(msg.Payload)
			log.Info(msg.ActualMessage)
		}
	} else {
		log.Warn("read failed")
	}
}

func (api *PublicWhisperAPI) MessageReceive(ctx context.Context, messageFilterID string) ([]*Message, error) {
	messages, err := api.GetFilterMessages(messageFilterID)
	return messages, err
}

func (api *PublicWhisperAPI) MessageSetupListener(ctx context.Context, topic string, hexKey string) (string, error) {
	// Allowed number of topics
	const topicNum = 1

	var _filter Filter
	var _criteria Criteria

	if hexKey == "" {
		hexKey = "7c8d019192c24224e2cafccae3a61fb586b14323a6bc8f9e7df1d929333ff993"
	}

	topicHex, err := api.MessageGenerateTopic(ctx, topic)
	randomKey, err := hex.DecodeString(hexKey)
	//log.Info(hex.EncodeToString(randomKey))

	symKey_id, err := api.AddSymKey(ctx, randomKey)
	symKey, err := api.w.GetSymKey(symKey_id)

	//log.Info(string(symKey_id))

	// Create a new Subscription
	_filter.Topics = make([][]byte, topicNum)
	log.Info("Topic String: " + topic)
	for i := 0; i < topicNum; i++ {
		_filter.Topics[i] = []byte(topicHex.String())
		log.Info("Topic Filter: " + string(_filter.Topics[i]))
	}
	_filter.KeySym = symKey
	_filter.AllowP2P = true

	// Set a new filter in order to poll new messages
	_criteria.Topics = make([]TopicType, 1)
	_criteria.Topics[0] = topicHex
	_criteria.SymKeyID = symKey_id
	log.Info("Topic: " + _criteria.Topics[0].String())

	messageFilterID, err := api.NewMessageFilter(_criteria)
	log.Info("Filter ID: " + messageFilterID)
	if err != nil {
		log.Error(fmt.Sprintf("%v\n", err))
	}

	return messageFilterID, err
}

// Expects topic string in plain text or in hex format
func (api *PublicWhisperAPI) MessageGenerateTopic(ctx context.Context, topic string) (TopicType, error) {
	// Handle topic string that is already in hex format
	if strings.HasPrefix(topic, "0x") {
		topic = strings.Replace(topic, "0x", "", 1)
		topicHex, err := hex.DecodeString(topic)
		log.Info("Topic given recognised as already in hex format")
		return BytesToTopic(topicHex), err
	} else {
		// Set default topic
		if topic == "" {
			topic = DefaultTopic
			log.Info("Topic not given, using default topic: \"" + topic + "\"")
		}

		topicHex := BytesToTopic([]byte(topic))
		log.Info("Generated topic \"" + topicHex.String() + "\" from value \"" + topic +"\"")
		return topicHex, nil
	}
}

func BytesToString(data []byte) string {
	return string(data[:])
}
// AddPrivateKey imports the given private key.
func (api *PublicWhisperAPI) AddPrivateKey(ctx context.Context, privateKey hexutil.Bytes) (string, error) {
	key, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return "", err
	}
	return api.w.AddKeyPair(key)
}

// DeleteKeyPair removes the key with the given key if it exists.
func (api *PublicWhisperAPI) DeleteKeyPair(ctx context.Context, key string) (bool, error) {
	if ok := api.w.DeleteKeyPair(key); ok {
		return true, nil
	}
	return false, fmt.Errorf("key pair %s not found", key)
}

// HasKeyPair returns an indication if the node has a key pair that is associated with the given id.
func (api *PublicWhisperAPI) HasKeyPair(ctx context.Context, id string) bool {
	return api.w.HasKeyPair(id)
}

// GetPublicKey returns the public key associated with the given key. The key is the hex
// encoded representation of a key in the form specified in section 4.3.6 of ANSI X9.62.
func (api *PublicWhisperAPI) GetPublicKey(ctx context.Context, id string) (hexutil.Bytes, error) {
	key, err := api.w.GetPrivateKey(id)
	if err != nil {
		return hexutil.Bytes{}, err
	}
	return crypto.FromECDSAPub(&key.PublicKey), nil
}

// GetPublicKey returns the private key associated with the given key. The key is the hex
// encoded representation of a key in the form specified in section 4.3.6 of ANSI X9.62.
func (api *PublicWhisperAPI) GetPrivateKey(ctx context.Context, id string) (hexutil.Bytes, error) {
	key, err := api.w.GetPrivateKey(id)
	if err != nil {
		return hexutil.Bytes{}, err
	}
	return crypto.FromECDSA(key), nil
}

// NewSymKey generate a random symmetric key.
// It returns an ID that can be used to refer to the key.
// Can be used encrypting and decrypting messages where the key is known to both parties.
func (api *PublicWhisperAPI) NewSymKey(ctx context.Context) (string, error) {
	return api.w.GenerateSymKey()
}

// AddSymKey import a symmetric key.
// It returns an ID that can be used to refer to the key.
// Can be used encrypting and decrypting messages where the key is known to both parties.
func (api *PublicWhisperAPI) AddSymKey(ctx context.Context, key hexutil.Bytes) (string, error) {
	return api.w.AddSymKeyDirect([]byte(key))
}

// GenerateSymKeyFromPassword derive a key from the given password, stores it, and returns its ID.
func (api *PublicWhisperAPI) GenerateSymKeyFromPassword(ctx context.Context, passwd string) (string, error) {
	return api.w.AddSymKeyFromPassword(passwd)
}

// HasSymKey returns an indication if the node has a symmetric key associated with the given key.
func (api *PublicWhisperAPI) HasSymKey(ctx context.Context, id string) bool {
	return api.w.HasSymKey(id)
}

// GetSymKey returns the symmetric key associated with the given id.
func (api *PublicWhisperAPI) GetSymKey(ctx context.Context, id string) (hexutil.Bytes, error) {
	return api.w.GetSymKey(id)
}

// DeleteSymKey deletes the symmetric key that is associated with the given id.
func (api *PublicWhisperAPI) DeleteSymKey(ctx context.Context, id string) bool {
	return api.w.DeleteSymKey(id)
}

//go:generate gencodec -type NewMessage -field-override newMessageOverride -out gen_newmessage_json.go

// NewMessage represents a new whisper message that is posted through the RPC.
type NewMessage struct {
	SymKeyID   string    `json:"symKeyID"`
	PublicKey  []byte    `json:"pubKey"`
	Sig        string    `json:"sig"`
	TTL        uint32    `json:"ttl"`
	Topic      TopicType `json:"topic"`
	Payload    []byte    `json:"payload"`
	Padding    []byte    `json:"padding"`
	PowTime    uint32    `json:"powTime"`
	PowTarget  float64   `json:"powTarget"`
	TargetPeer string    `json:"targetPeer"`
}

type newMessageOverride struct {
	PublicKey hexutil.Bytes
	Payload   hexutil.Bytes
	Padding   hexutil.Bytes
}

// Post a message on the Whisper network.
func (api *PublicWhisperAPI) Post(ctx context.Context, req NewMessage) (bool, error) {
	var (
		symKeyGiven = len(req.SymKeyID) > 0
		pubKeyGiven = len(req.PublicKey) > 0
		err         error
	)

	// user must specify either a symmetric or an asymmetric key
	if (symKeyGiven && pubKeyGiven) || (!symKeyGiven && !pubKeyGiven) {
		return false, ErrSymAsym
	}

	params := &MessageParams{
		TTL:      req.TTL,
		Payload:  req.Payload,
		Padding:  req.Padding,
		WorkTime: req.PowTime,
		PoW:      req.PowTarget,
		Topic:    req.Topic,
	}

	// Set key that is used to sign the message
	if len(req.Sig) > 0 {
		if params.Src, err = api.w.GetPrivateKey(req.Sig); err != nil {
			return false, err
		}
	}

	// Set symmetric key that is used to encrypt the message
	if symKeyGiven {
		if params.Topic == (TopicType{}) { // topics are mandatory with symmetric encryption
			return false, ErrNoTopics
		}
		if params.KeySym, err = api.w.GetSymKey(req.SymKeyID); err != nil {
			return false, err
		}
		if !validateSymmetricKey(params.KeySym) {
			return false, ErrInvalidSymmetricKey
		}
	}

	// Set asymmetric key that is used to encrypt the message
	if pubKeyGiven {
		params.Dst = crypto.ToECDSAPub(req.PublicKey)
		if !ValidatePublicKey(params.Dst) {
			return false, ErrInvalidPublicKey
		}
	}

	// encrypt and sent message
	whisperMsg, err := NewSentMessage(params)
	if err != nil {
		return false, err
	}

	env, err := whisperMsg.Wrap(params)
	if err != nil {
		return false, err
	}

	// send to specific node (skip PoW check)
	if len(req.TargetPeer) > 0 {
		n, err := discover.ParseNode(req.TargetPeer)
		if err != nil {
			return false, fmt.Errorf("failed to parse target peer: %s", err)
		}
		return true, api.w.SendP2PMessage(n.ID[:], env)
	}

	// ensure that the message PoW meets the node's minimum accepted PoW
	if req.PowTarget < api.w.MinPow() {
		return false, ErrTooLowPoW
	}

	return true, api.w.Send(env)
}

//go:generate gencodec -type Criteria -field-override criteriaOverride -out gen_criteria_json.go

// Criteria holds various filter options for inbound messages.
type Criteria struct {
	SymKeyID     string      `json:"symKeyID"`
	PrivateKeyID string      `json:"privateKeyID"`
	Sig          []byte      `json:"sig"`
	MinPow       float64     `json:"minPow"`
	Topics       []TopicType `json:"topics"`
	AllowP2P     bool        `json:"allowP2P"`
}

type criteriaOverride struct {
	Sig hexutil.Bytes
}

// Messages set up a subscription that fires events when messages arrive that match
// the given set of criteria.
func (api *PublicWhisperAPI) Messages(ctx context.Context, crit Criteria) (*rpc.Subscription, error) {
	var (
		symKeyGiven = len(crit.SymKeyID) > 0
		pubKeyGiven = len(crit.PrivateKeyID) > 0
		err         error
	)

	// ensure that the RPC connection supports subscriptions
	notifier, supported := rpc.NotifierFromContext(ctx)
	if !supported {
		return nil, rpc.ErrNotificationsUnsupported
	}

	// user must specify either a symmetric or an asymmetric key
	if (symKeyGiven && pubKeyGiven) || (!symKeyGiven && !pubKeyGiven) {
		return nil, ErrSymAsym
	}

	filter := Filter{
		PoW:      crit.MinPow,
		Messages: make(map[common.Hash]*ReceivedMessage),
		AllowP2P: crit.AllowP2P,
	}

	if len(crit.Sig) > 0 {
		filter.Src = crypto.ToECDSAPub(crit.Sig)
		if !ValidatePublicKey(filter.Src) {
			return nil, ErrInvalidSigningPubKey
		}
	}

	for i, bt := range crit.Topics {
		if len(bt) == 0 || len(bt) > 4 {
			return nil, fmt.Errorf("subscribe: topic %d has wrong size: %d", i, len(bt))
		}
		filter.Topics = append(filter.Topics, bt[:])
	}

	// listen for message that are encrypted with the given symmetric key
	if symKeyGiven {
		if len(filter.Topics) == 0 {
			return nil, ErrNoTopics
		}
		key, err := api.w.GetSymKey(crit.SymKeyID)
		if err != nil {
			return nil, err
		}
		if !validateSymmetricKey(key) {
			return nil, ErrInvalidSymmetricKey
		}
		filter.KeySym = key
		filter.SymKeyHash = crypto.Keccak256Hash(filter.KeySym)
	}

	// listen for messages that are encrypted with the given public key
	if pubKeyGiven {
		filter.KeyAsym, err = api.w.GetPrivateKey(crit.PrivateKeyID)
		if err != nil || filter.KeyAsym == nil {
			return nil, ErrInvalidPublicKey
		}
	}

	id, err := api.w.Subscribe(&filter)
	if err != nil {
		return nil, err
	}

	// create subscription and start waiting for message events
	rpcSub := notifier.CreateSubscription()
	go func() {
		// for now poll internally, refactor whisper internal for channel support
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if filter := api.w.GetFilter(id); filter != nil {
					for _, rpcMessage := range toMessage(filter.Retrieve()) {
						if err := notifier.Notify(rpcSub.ID, rpcMessage); err != nil {
							log.Error("Failed to send notification", "err", err)
						}
					}
				}
			case <-rpcSub.Err():
				api.w.Unsubscribe(id)
				return
			case <-notifier.Closed():
				api.w.Unsubscribe(id)
				return
			}
		}
	}()

	return rpcSub, nil
}

//go:generate gencodec -type Message -field-override messageOverride -out gen_message_json.go

// Message is the RPC representation of a whisper message.
type Message struct {
	Sig       []byte    `json:"sig,omitempty"`
	TTL       uint32    `json:"ttl"`
	Timestamp uint32    `json:"timestamp"`
	Topic     TopicType `json:"topic"`
	Payload   []byte    `json:"payload"`
	Padding   []byte    `json:"padding"`
	PoW       float64   `json:"pow"`
	Hash      []byte    `json:"hash"`
	Dst       []byte    `json:"recipientPublicKey,omitempty"`
	ActualMessage string `json:"actualMessage,omitempty"`
}

type messageOverride struct {
	Sig     hexutil.Bytes
	Payload hexutil.Bytes
	Padding hexutil.Bytes
	Hash    hexutil.Bytes
	Dst     hexutil.Bytes
}

// ToWhisperMessage converts an internal message into an API version.
func ToWhisperMessage(message *ReceivedMessage) *Message {
	msg := Message{
		Payload:   message.Payload,
		Padding:   message.Padding,
		Timestamp: message.Sent,
		TTL:       message.TTL,
		PoW:       message.PoW,
		Hash:      message.EnvelopeHash.Bytes(),
		Topic:     message.Topic,
	}

	if message.Dst != nil {
		b := crypto.FromECDSAPub(message.Dst)
		if b != nil {
			msg.Dst = b
		}
	}

	if isMessageSigned(message.Raw[0]) {
		b := crypto.FromECDSAPub(message.SigToPubKey())
		if b != nil {
			msg.Sig = b
		}
	}

	return &msg
}

// toMessage converts a set of messages to its RPC representation.
func toMessage(messages []*ReceivedMessage) []*Message {
	msgs := make([]*Message, len(messages))
	for i, msg := range messages {
		msgs[i] = ToWhisperMessage(msg)
	}
	return msgs
}

// GetFilterMessages returns the messages that match the filter criteria and
// are received between the last poll and now.
func (api *PublicWhisperAPI) GetFilterMessages(id string) ([]*Message, error) {
	api.mu.Lock()
	f := api.w.GetFilter(id)
	if f == nil {
		api.mu.Unlock()
		return nil, fmt.Errorf("filter not found")
	}
	api.lastUsed[id] = time.Now()
	api.mu.Unlock()

	receivedMessages := f.Retrieve()
	messages := make([]*Message, 0, len(receivedMessages))
	for _, msg := range receivedMessages {
		messages = append(messages, ToWhisperMessage(msg))
	}

	return messages, nil
}

// DeleteMessageFilter deletes a filter.
func (api *PublicWhisperAPI) DeleteMessageFilter(id string) (bool, error) {
	api.mu.Lock()
	defer api.mu.Unlock()

	delete(api.lastUsed, id)
	return true, api.w.Unsubscribe(id)
}

// NewMessageFilter creates a new filter that can be used to poll for
// (new) messages that satisfy the given criteria.
func (api *PublicWhisperAPI) NewMessageFilter(req Criteria) (string, error) {
	var (
		src     *ecdsa.PublicKey
		keySym  []byte
		keyAsym *ecdsa.PrivateKey
		topics  [][]byte

		symKeyGiven  = len(req.SymKeyID) > 0
		asymKeyGiven = len(req.PrivateKeyID) > 0

		err error
	)

	// user must specify either a symmetric or an asymmetric key
	if (symKeyGiven && asymKeyGiven) || (!symKeyGiven && !asymKeyGiven) {
		return "", ErrSymAsym
	}

	if len(req.Sig) > 0 {
		src = crypto.ToECDSAPub(req.Sig)
		if !ValidatePublicKey(src) {
			return "", ErrInvalidSigningPubKey
		}
	}

	if symKeyGiven {
		if keySym, err = api.w.GetSymKey(req.SymKeyID); err != nil {
			return "", err
		}
		if !validateSymmetricKey(keySym) {
			return "", ErrInvalidSymmetricKey
		}
	}

	if asymKeyGiven {
		if keyAsym, err = api.w.GetPrivateKey(req.PrivateKeyID); err != nil {
			return "", err
		}
	}

	if len(req.Topics) > 0 {
		topics = make([][]byte, 1)
		for _, topic := range req.Topics {
			topics = append(topics, topic[:])
		}
	}

	f := &Filter{
		Src:      src,
		KeySym:   keySym,
		KeyAsym:  keyAsym,
		PoW:      req.MinPow,
		AllowP2P: req.AllowP2P,
		Topics:   topics,
		Messages: make(map[common.Hash]*ReceivedMessage),
	}

	id, err := api.w.Subscribe(f)
	if err != nil {
		return "", err
	}

	api.mu.Lock()
	api.lastUsed[id] = time.Now()
	api.mu.Unlock()

	return id, nil
}
