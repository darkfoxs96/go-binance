/*

   Stores response structs for API functions user_responce.go

*/

package binance

import (
	"github.com/shopspring/decimal"
)

type WSAccountUpdate struct {
	//EventType            string       `json:"e"` // Event type
	//EventTime            int64        `json:"E"` // Event time
	//MakerCommissionRate  int64        `json:"m"` // Maker commission rate (bips)
	//TakerCommissionRate  int64        `json:"t"` // Taker commission rate (bips)
	//BuyerCommissionRate  int64        `json:"b"` // Buyer commission rate (bips)
	//SellerCommissionRate int64        `json:"s"` // Seller commission rate (bips)
	//CanTrad              bool         `json:"T"` // Can trade?
	//CanWithdraw          bool         `json:"W"` // Can withdraw?
	//CanDeposit           bool         `json:"D"` // Can deposit?
	//LastAccountUpdate    int64      `json:"u"` // Time of last account update
	Balance              []*Balance `json:"B"`
}

type WSOrderUpdate struct {
	EventType                              string          `json:"e"` // Event type
	EventTime                              int64           `json:"E"` // Event time
	Symbol                                 string          `json:"s"` // Symbol
	//ClientOrderID                          string          `json:"c"` // Client order ID
	Side                                   string          `json:"S"` // Side
	OrderType                              string          `json:"o"` // Order type
	TimeInForce                            string          `json:"f"` // Time in force
	OrderQuantity                          decimal.Decimal `json:"q"` // Order quantity
	Q                          			   decimal.Decimal `json:"Q"`
	OrderPrice                             decimal.Decimal `json:"p"` // Order price
	StopPrice                              decimal.Decimal `json:"P"` // Stop price
	IcebergQuantity                        decimal.Decimal `json:"F"` // Iceberg quantity
	//Ignore1                                int64            `json:"g"` // Ignore // TODO: maybe int32 ??
	//OriginalClientOrderID                  string          `json:"C"` // Original client order ID; This is the ID of the order being canceled
	CurrentExecutionType                   string          `json:"x"` // Current execution type
	// CurrentExecutionType:
	// NEW
	// CANCELED
	// REPLACED (currently unused)
	// REJECTED
	// TRADE
	// EXPIRED
	CurrentOrderStatus                     string          `json:"X"` // Current order status
	//OrderRejectReason                      string          `json:"r"` // Order reject reason; will be an error code.
	OrderID                                int64           `json:"i"` // Order ID
	LastExecutedQuantity                   decimal.Decimal `json:"l"` // Last executed quantity
	CumulativeFilledQuantity               decimal.Decimal `json:"z"` // Cumulative filled quantity
	LastExecutedPrice                      decimal.Decimal `json:"L"` // Last executed price
	CommissionAmount                       decimal.Decimal `json:"n"` // Commission amount
	CommissionAsset                        *string           `json:"N"` // Commission asset
	TransactionTime                        int64           `json:"T"` // Transaction time
	TradeID                                int64           `json:"t"` // Trade ID
	Ignore2                                int64           `json:"I"` // Ignore
	IsOrderWorking                         bool            `json:"w"` // Is the order working? Stops will have
	//IsTradeMakerSide                       bool            `json:"m"` // Is this trade the maker side?
	//Ignore3                                bool            `json:"M"` // Ignore
	OrderCreationTime                      int64           `json:"O"` // Order creation time
	CumulativeQuoteAssetTransactedQuantity decimal.Decimal `json:"Z"` // Cumulative quote asset transacted quantity
	LastQuoteAssetTransactedQuantity       decimal.Decimal `json:"Y"` // Last quote asset transacted quantity (i.e. lastPrice * lastQty)
}
