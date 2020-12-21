/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type Currency struct {
	// Currency name
	Currency string `json:"currency,omitempty"`
	// Whether currency is de-listed
	Delisted bool `json:"delisted,omitempty"`
	// Whether currency's withdrawal is disabled
	WithdrawDisabled bool `json:"withdraw_disabled,omitempty"`
	// Whether currency's withdrawal is delayed
	WithdrawDelayed bool `json:"withdraw_delayed,omitempty"`
	// Whether currency's deposit is disabled
	DepositDisabled bool `json:"deposit_disabled,omitempty"`
	// Whether currency's trading is disabled
	TradeDisabled bool `json:"trade_disabled,omitempty"`
}
