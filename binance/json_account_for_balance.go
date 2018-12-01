/*

   Stores response structs for API functions account.go

*/

package binance

import (
	"github.com/buger/jsonparser"
	"github.com/shopspring/decimal"
)

var pathToBalance = [][]string{
	{"balances"},
}

func getBalanceByAccount(data []byte) (balances []*Balance) {
	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0: // balances
			balances = getBalanceByAccountFor(value)
		}
	}, pathToBalance...)

	return
}

func getBalanceByAccountFor(data []byte) (balances []*Balance)  {
	balances = make([]*Balance, 0, 10)

	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		b := &Balance{}

		jsonparser.ObjectEach(value, func(key []byte, d []byte, dataType jsonparser.ValueType, offset int) error {

			k := string(key)
			if k == "asset" || k == "a" {
				b.Asset = string(d)
			} else if k == "free" || k == "f" {
				b.Free, _ = decimal.NewFromString(string(d))
			} else if k == "locked" || k == "l" {
				b.Locked, _ = decimal.NewFromString(string(d))
			}

			return nil
		})

		balances = append(balances, b)
	})

	return
}
