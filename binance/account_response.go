/*

   Stores response structs for API functions account.go

*/

package binance

import (
	"github.com/shopspring/decimal"
)

// Result from: GET /api/v3/account
type Account struct {
	MakerCommission  int64     `json:"makerCommission"`
	TakerCommission  int64     `json:"takerCommission"`
	BuyerCommission  int64     `json:"buyerCommission"`
	SellerCommission int64     `json:"sellerCommission"`
	CanTrade         bool      `json:"canTrade"`
	CanWithdraw      bool      `json:"canWithdraw"`
	CanDeposit       bool      `json:"canDeposit"`
	Balances         []Balance `json:"balances"`
}

type Balance struct {
	Asset  string
	Free   decimal.Decimal
	Locked decimal.Decimal
}

// Result from: POST /api/v3/order
type PlacedOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
	TransactTime  int64  `json:"transactTime"`
}

// Result from: DELETE /api/v3/order
type CanceledOrder struct {
	Symbol            string `json:"symbol"`
	OrigClientOrderId string `json:"origClientOrderId"`
	OrderId           int64  `json:"orderId"`
	ClientOrderId     string `json:"clientOrderId"`
}

// Result from: GET /api/v3/order
type OrderStatus struct {
	Symbol        string  `json:"symbol"`
	OrderId       int64   `json:"orderId"`
	ClientOrderId string  `json:"clientOrderId"`
	Price         float64 `json:"price,string"`
	OrigQty       float64 `json:"origQty,string"`
	ExecutedQty   float64 `json:"executedQty,string"`
	Status        string  `json:"status"`
	TimeInForce   string  `json:"timeInForce"`
	Type          string  `json:"type"`
	Side          string  `json:"side"`
	StopPrice     float64 `json:"stopPrice,string"`
	IcebergQty    float64 `json:"icebergQty,string"`
	Time          int64   `json:"time"`
}

// Result from: GET /api/v3/myTrades
type Trade struct {
	Id              int64   `json:"id"`
	Price           float64 `json:"price,string"`
	Quantity        float64 `json:"qty,string"`
	Commission      float64 `json:"commission,string"`
	CommissionAsset string  `json:"commissionAsset"`
	Time            int64   `json:"time"`
	IsBuyer         bool    `json:"isBuyer"`
	IsMaker         bool    `json:"isMaker"`
	IsBestMatch     bool    `json:"isBestMatch"`
}

// Result from: GET /api/v3/depositHistory
type Deposit struct {
	InsertTime int64   `json:"insertTime"`
	Amount     float64 `json:"amount"`
	Asset      string  `json:"asset"`
	Address    string  `json:"address"`
	TxId       string  `json:"txId"`
	Status     int64   `json:"status"`
}

// Result from: GET /api/v3/withdrawHistory
type Withdraw struct {
	Id        string  `json:"id"`
	Amount    float64 `json:"amount"`
	Address   string  `json:"address"`
	Asset     string  `json:"asset"`
	TxId      string  `json:"txId"`
	ApplyTime int64   `json:"applyTime"`
	Status    int64   `json:"status"`
}

type AssetDetail struct {
	MinWithdrawAmount decimal.Decimal `json:"minWithdrawAmount"`
	WithdrawFee       decimal.Decimal `json:"withdrawFee"`
	DepositStatus     bool            `json:"depositStatus"`
	WithdrawStatus    bool            `json:"withdrawStatus"`
	DepositTip        string          `json:"depositTip"`
}

type WithdrawList struct {
	Withdraws []Withdraw `json:"withdrawList"`
}

type DepositList struct {
	Deposits []Deposit `json:"depositList"`
}

type AssetDetailList struct {
	AssetDetail map[string]AssetDetail `json:"assetDetail"`
}
