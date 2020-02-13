/*

   Stores response structs for API functions account.go

*/

package binance

import (
	"fmt"
	"encoding/json"

	"github.com/buger/jsonparser"
	"github.com/shopspring/decimal"
)

var (
	pathsToEvent = [][]string{
		{"e"},
	}
	pathsInWSAccount = [][]string{
		{"B"},
	}
	pathsInWSOrder = [][]string{
		{"E"},
		{"s"},
		{"S"},
		{"O"},
		{"q"},
		{"p"},
		{"P"},
		{"F"},
		{"x"},
		{"X"},
		{"i"},
		{"l"},
		{"z"},
		{"L"},
		{"n"},
		{"T"},
		{"w"},
		{"O"},
		{"Z"},
		{"Y"},
	}
)

func isWSUpdateOrder(data []byte) (isOrder bool) {
	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0: // EventType
			isOrder = string(value) == "executionReport"
		}
	}, pathsToEvent...)

	return
}

func getWSUpdateAccount(data []byte) *WSAccountUpdate  {
	account := &WSAccountUpdate{}

	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0: // Balance
			account.Balance = getBalanceByAccountFor(value)
		}
	}, pathsInWSAccount...)

	return account
}

func getWSUpdateOrder(data []byte) *WSOrderUpdate {
	order := &WSOrderUpdate{}
	err := json.Unmarshal(data, order)
	if err != nil {
		fmt.Println("getWSUpdateOrder():", err)
	}

	return order

	jsonparser.EachKey(data, func(idx int, value []byte, vt jsonparser.ValueType, err error){
		switch idx {
		case 0: // EventTime
			order.EventTime, _ = jsonparser.GetInt(value)
		case 1: // Symbol
			order.Symbol = string(value)
		case 2: // Side
			order.Side = string(value)
		case 3: // OrderType
			order.OrderType = string(value)
		case 4: // OrderQuantity
			order.OrderQuantity, _ = decimal.NewFromString(string(value))
		case 5: // OrderPrice
			order.OrderPrice, _ = decimal.NewFromString(string(value))
		case 6: // StopPrice
			order.StopPrice, _ = decimal.NewFromString(string(value))
		case 7: // IcebergQuantity
			order.IcebergQuantity, _ = decimal.NewFromString(string(value))
		case 8: // CurrentExecutionType
			order.CurrentExecutionType = string(value)
		case 9: // CurrentOrderStatus
			order.CurrentOrderStatus = string(value)
		case 10: // OrderID
			order.OrderID, _ = jsonparser.GetInt(value)
		case 11: // LastExecutedQuantity
			order.LastExecutedQuantity, _ = decimal.NewFromString(string(value))
		case 12: // CumulativeFilledQuantity
			order.CumulativeFilledQuantity, _ = decimal.NewFromString(string(value))
		case 13: // LastExecutedPrice
			order.LastExecutedPrice, _ = decimal.NewFromString(string(value))
		case 14: // CommissionAmount
			order.CommissionAmount, _ = decimal.NewFromString(string(value))
		case 15: // TransactionTime
			order.TransactionTime, _ = jsonparser.GetInt(value)
		case 16: // IsOrderWorking
			order.IsOrderWorking, _ = jsonparser.GetBoolean(value)
		case 17: // OrderCreationTime
			order.OrderCreationTime, _ = jsonparser.GetInt(value)
		case 18: // CumulativeQuoteAssetTransactedQuantity
			order.CumulativeQuoteAssetTransactedQuantity, _ = decimal.NewFromString(string(value))
		case 19: // LastQuoteAssetTransactedQuantity
			order.LastQuoteAssetTransactedQuantity, _ = decimal.NewFromString(string(value))
		}
	}, pathsInWSOrder...)

	return order
}
