/*

   Stores response structs for API functions user.go

*/

package binance

import (
	"fmt"
	"errors"
	"time"
	"encoding/json"

	"github.com/gorilla/websocket"
	goBinance "github.com/OlegFX/go-binance"
)

const (
	userDataWSUrl     = "wss://stream.binance.com:9443/ws/"
)

var (
	keyUserData       = ""
	userDataStreamUrl = fmt.Sprintf("api/v1/userDataStream") // TODO: delete
)

// NewUserWSChannel POST /api/v1/userDataStream
func (c *Binance) NewUserWSChannel() (err error,
								      accountUpdate chan *WSAccountUpdate,
									  orderUpdate chan *WSOrderUpdate,
									  done chan struct{}) {
	clientGoBinance := goBinance.NewClient(c.client.key, c.client.secret)
	clientGoBinanceWS := clientGoBinance.NewStartUserStreamService()

	keyUserData, err = clientGoBinanceWS.Do(emptyContext{})
	if err != nil {
		fmt.Println("Binance WS ERROR: ", err)
	}
	if keyUserData == "" {
		return errors.New("Binance WS ERROR: key for ws user data == nil"), nil, nil, nil
	}

	conn, _, err := websocket.DefaultDialer.Dial(getUserDataWSUrl(keyUserData), nil)
	if err != nil {
		return
	}

	done = make(chan struct{})
	accountUpdate = make(chan *WSAccountUpdate, 2)
	orderUpdate = make(chan *WSOrderUpdate, 2)

	go func() {
		defer close(done)

		fmt.Println("Binance userData start =>>>>>>")
		for {
			_, b, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Binance ws read: ", err)
				continue
			}

			fmt.Printf("Binance ws recv: %s", string(b))

			var accountUp *WSAccountUpdate
			err = json.Unmarshal(b, accountUp)
			if err != nil {
				var orderUp *WSOrderUpdate
				err = json.Unmarshal(b, orderUp)
				if err != nil {
					fmt.Printf("Binance ws read: not found structures, error: " + err.Error())
					return
				}

				orderUpdate <- orderUp
			} else {
				accountUpdate <- accountUp
			}
		}
	}()

	return
}

// UpdateUserWSChannel PUT /api/v1/userDataStream
func (c *Binance) UpdateUserWSChannel() {
	c.client.do("PUT", userDataStreamUrl, "", false, nil)
}

// StopUserWSChannel DELETE /api/v1/userDataStream
func (c *Binance) StopUserWSChannel() {
	c.client.do("DELETE", userDataStreamUrl, "", false, nil)
}

func getUserDataWSUrl(key string) string  {
	return fmt.Sprintf(userDataWSUrl + key)
}

type emptyContext struct {}
func (c emptyContext) Deadline() (deadline time.Time, ok bool) { return }
func (c emptyContext) Done() <-chan struct{} { return make(chan struct{}) }
func (c emptyContext) Err() error { return errors.New("") }
func (c emptyContext) Value(key interface{}) interface{} { return "" }
