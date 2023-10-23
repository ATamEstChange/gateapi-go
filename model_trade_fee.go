/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type TradeFee struct {
	// User ID
	UserId int64 `json:"user_id,omitempty"`
	// taker fee rate
	TakerFee string `json:"taker_fee,omitempty"`
	// maker fee rate
	MakerFee string `json:"maker_fee,omitempty"`
	// If GT deduction is enabled
	GtDiscount bool `json:"gt_discount,omitempty"`
	// Taker fee rate if using GT deduction. It will be 0 if GT deduction is disabled
	GtTakerFee string `json:"gt_taker_fee,omitempty"`
	// Maker fee rate if using GT deduction. It will be 0 if GT deduction is disabled
	GtMakerFee string `json:"gt_maker_fee,omitempty"`
	// Loan fee rate of margin lending
	LoanFee string `json:"loan_fee,omitempty"`
	// Point type. 0 - Initial version. 1 - new version since 202009
	PointType string `json:"point_type,omitempty"`
	// Futures trading taker fee
	FuturesTakerFee string `json:"futures_taker_fee,omitempty"`
	// Future trading maker fee
	FuturesMakerFee string `json:"futures_maker_fee,omitempty"`
	// Delivery trading taker fee
	DeliveryTakerFee string `json:"delivery_taker_fee,omitempty"`
	// Delivery trading maker fee
	DeliveryMakerFee string `json:"delivery_maker_fee,omitempty"`
}
