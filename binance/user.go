/*

   Stores response structs for API functions user.go

*/

package binance

import (
	"errors"
	"fmt"
	"sync"
	"time"

	goBinance "github.com/OlegFX/go-binance"
	"github.com/gorilla/websocket"
)

const (
	userDataWSUrl = "wss://stream.binance.com:9443/ws/"
)

type DoneChannel struct {
	ch       chan interface{}
	isClosed bool
	mutex    sync.RWMutex
}

func (d DoneChannel) IsClosed() bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.isClosed
}

func (d DoneChannel) Write(data interface{}) bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.isClosed {
		return false
	}

	d.ch <- data
	return true
}

func (d DoneChannel) Read() <-chan interface{} {
	return d.ch
}

func (d DoneChannel) Close() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.isClosed {
		return
	}

	d.isClosed = true
	close(d.ch)
}

func NewDoneChannel() *DoneChannel {
	return &DoneChannel{
		ch: make(chan interface{}, 2),
	}
}

var (
	keyUserData       = ""
	userDataStreamUrl = fmt.Sprintf("api/v1/userDataStream") // TODO: delete
	clientGoBinance   *goBinance.Client
)

// NewUserWSChannel POST /api/v1/userDataStream
func (c *Binance) NewUserWSChannel() (err error,
	accountUpdate chan *WSAccountUpdate,
	orderUpdate chan *WSOrderUpdate,
	done *DoneChannel) {
	clientGoBinance = goBinance.NewClient(c.client.key, c.client.secret)
	clientGoBinanceWS := clientGoBinance.NewStartUserStreamService()

	keyUserData, err = clientGoBinanceWS.Do(emptyContext{})
	if err != nil {
		return errors.New("Binance WS clientGoBinanceWS.Do(emptyContext{}) ERROR: " + err.Error()), nil, nil, nil
	}
	if keyUserData == "" {
		return errors.New("Binance WS ERROR: key for ws user data == nil"), nil, nil, nil
	}

	conn, _, err := websocket.DefaultDialer.Dial(getUserDataWSUrl(keyUserData), nil)
	if err != nil {
		return
	}

	done = NewDoneChannel()
	accountUpdate = make(chan *WSAccountUpdate, 50)
	orderUpdate = make(chan *WSOrderUpdate, 50)

	go func() {
		defer func() {
			if r := recover(); r != nil {
				done.Write(r)
			}

			done.Close()
			close(accountUpdate)
			close(orderUpdate)
			c.StopUserWSChannel()
		}()

		for {
			_, b, err := conn.ReadMessage()
			if err != nil {
				done.Write(err)
				break
			}

			if isWSUpdateOrder(b) {
				if data := getWSUpdateOrder(b); !done.IsClosed() {
					orderUpdate <- data
				}
			} else {
				if data := getWSUpdateAccount(b); !done.IsClosed() {
					accountUpdate <- data
				}
			}
		}
	}()

	go c.UpdateUserWSChannel(done)

	return
}

// UpdateUserWSChannel PUT /api/v1/userDataStream
func (c *Binance) UpdateUserWSChannel(done *DoneChannel) {
	for {
		keepalive := clientGoBinance.NewKeepaliveUserStreamService()
		keepalive.ListenKey(keyUserData)
		err := keepalive.Do(emptyContext{})
		if err != nil {
			done.Write(errors.New("Binance error update userData stream: " + err.Error()))
			done.Close()
			return
		}

		time.Sleep(time.Minute * 30)
	}
}

// StopUserWSChannel DELETE /api/v1/userDataStream
func (c *Binance) StopUserWSChannel() error {
	if keyUserData != "" {
		closeUserStream := clientGoBinance.NewCloseUserStreamService()
		closeUserStream.ListenKey(keyUserData)
		keyUserData = ""

		return closeUserStream.Do(emptyContext{})
	}

	return nil
}

func getUserDataWSUrl(key string) string {
	return fmt.Sprintf(userDataWSUrl + key)
}

type emptyContext struct{}

func (c emptyContext) Deadline() (deadline time.Time, ok bool) { return }
func (c emptyContext) Done() <-chan struct{}                   { return make(chan struct{}) }
func (c emptyContext) Err() error                              { return errors.New("") }
func (c emptyContext) Value(key interface{}) interface{}       { return "" }
