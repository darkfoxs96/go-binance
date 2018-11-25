/*

   Stores response structs for API functions user.go

*/

package binance

import (
	"github.com/shopspring/decimal"
)

type WSBalance struct {
	Asset  string          `json:"a"` // Asset
	Free   decimal.Decimal `json:"f"` // Free amount
	Locked decimal.Decimal `json:"l"` // Locked amount
}

type WSAccountUpdate struct {
	EventType            string       `json:"e"` // Event type
	EventTime            int32        `json:"E"` // Event time
	MakerCommissionRate  int32        `json:"m"` // Maker commission rate (bips)
	TakerCommissionRate  int32        `json:"t"` // Taker commission rate (bips)
	BuyerCommissionRate  int32        `json:"b"` // Buyer commission rate (bips)
	SellerCommissionRate int32        `json:"s"` // Seller commission rate (bips)
	CanTrad              bool         `json:"T"` // Can trade?
	CanWithdraw          bool         `json:"W"` // Can withdraw?
	CanDeposit           bool         `json:"D"` // Can deposit?
	LastAccountUpdate    int32        `json:"u"` // Time of last account update
	Balance              []*WSBalance `json:"B"`
}

type WSOrderUpdate struct {
	EventType                              string          `json:"e"` // Event type
	EventTime                              int32           `json:"E"` // Event time
	Symbol                                 string          `json:"s"` // Symbol
	ClientOrderID                          string          `json:"c"` // Client order ID
	Side                                   string          `json:"S"` // Side
	OrderType                              string          `json:"o"` // Order type
	TimeInForce                            string          `json:"f"` // Time in force
	OrderQuantity                          decimal.Decimal `json:"q"` // Order quantity
	OrderPrice                             decimal.Decimal `json:"p"` // Order price
	StopPrice                              decimal.Decimal `json:"P"` // Stop price
	IcebergQuantity                        decimal.Decimal `json:"F"` // Iceberg quantity
	Ignore1                                int8            `json:"g"` // Ignore // TODO: maybe int32 ??
	OriginalClientOrderID                  string          `json:"C"` // Original client order ID; This is the ID of the order being canceled
	CurrentExecutionType                   string          `json:"x"` // Current execution type
	// Execution types:
	// NEW
	// CANCELED
	// REPLACED (currently unused)
	// REJECTED
	// TRADE
	// EXPIRED
	CurrentOrderStatus                     string          `json:"X"` // Current order status
	OrderRejectReason                      string          `json:"r"` // Order reject reason; will be an error code.
	OrderID                                int64           `json:"i"` // Order ID
	LastExecutedQuantity                   decimal.Decimal `json:"l"` // Last executed quantity
	CumulativeFilledQuantity               decimal.Decimal `json:"z"` // Cumulative filled quantity
	LastExecutedPrice                      decimal.Decimal `json:"L"` // Last executed price
	CommissionAmount                       decimal.Decimal `json:"n"` // Commission amount
	CommissionAsset                        int32           `json:"N"` // Commission asset
	TransactionTime                        int32           `json:"T"` // Transaction time
	TradeID                                int32           `json:"t"` // Trade ID
	Ignore2                                int32           `json:"I"` // Ignore
	IsOrderWorking                         bool            `json:"w"` // Is the order working? Stops will have
	IsTradeMakerSide                       bool            `json:"m"` // Is this trade the maker side?
	Ignore3                                bool            `json:"M"` // Ignore
	OrderCreationTime                      int32           `json:"O"` // Order creation time
	CumulativeQuoteAssetTransactedQuantity decimal.Decimal `json:"Z"` // Cumulative quote asset transacted quantity
	LastQuoteAssetTransactedQuantity       decimal.Decimal `json:"Y"` // Last quote asset transacted quantity (i.e. lastPrice * lastQty)
}
