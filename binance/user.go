/*

   Stores response structs for API functions user.go

*/

package binance

import (
	"errors"
	"fmt"
	"time"

	goBinance "github.com/OlegFX/go-binance"
	"github.com/gorilla/websocket"
)

const (
	userDataWSUrl = "wss://stream.binance.com:9443/ws/"
)

var (
	keyUserData       = ""
	userDataStreamUrl = fmt.Sprintf("api/v1/userDataStream") // TODO: delete
	clientGoBinance   *goBinance.Client
)

// NewUserWSChannel POST /api/v1/userDataStream
func (c *Binance) NewUserWSChannel() (err error,
	accountUpdate chan *WSAccountUpdate,
	orderUpdate chan *WSOrderUpdate,
	done chan struct{}) {
	clientGoBinance = goBinance.NewClient(c.client.key, c.client.secret)
	clientGoBinanceWS := clientGoBinance.NewStartUserStreamService()

	keyUserData, err = clientGoBinanceWS.Do(emptyContext{})
	if err != nil {
		fmt.Println("Binance WS clientGoBinanceWS.Do(emptyContext{}) ERROR: ", err)
	}
	if keyUserData == "" {
		return errors.New("Binance WS ERROR: key for ws user data == nil"), nil, nil, nil
	}

	conn, _, err := websocket.DefaultDialer.Dial(getUserDataWSUrl(keyUserData), nil)
	if err != nil {
		return
	}

	done = make(chan struct{}, 2)
	accountUpdate = make(chan *WSAccountUpdate, 10)
	orderUpdate = make(chan *WSOrderUpdate, 10)

	go func() {
		defer func() {
			if r := recover(); r != nil {
			}

			close(done)
			close(accountUpdate)
			close(orderUpdate)
			c.StopUserWSChannel()
		}()

		fmt.Println("Binance start userData stream =>>>>>>")

		for {
			_, b, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Binance ws read: ", err)
				break
			}

			if isWSUpdateOrder(b) {
				go func() {
					orderUpdate <- getWSUpdateOrder(b)
				}()
			} else {
				go func() {
					accountUpdate <- getWSUpdateAccount(b)
				}()
			}
		}
	}()

	go c.UpdateUserWSChannel()

	return
}

// UpdateUserWSChannel PUT /api/v1/userDataStream
func (c *Binance) UpdateUserWSChannel() {
	for {
		keepalive := clientGoBinance.NewKeepaliveUserStreamService()
		keepalive.ListenKey(keyUserData)
		err := keepalive.Do(emptyContext{})
		if err != nil {
			fmt.Println("Binance error update userData stream: " + err.Error())
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
