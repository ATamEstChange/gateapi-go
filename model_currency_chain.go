/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type CurrencyChain struct {
	// Chain name
	Chain string `json:"chain,omitempty"`
	// Chain name in Chinese
	NameCn string `json:"name_cn,omitempty"`
	// Chain name in English
	NameEn string `json:"name_en,omitempty"`
	// If it is disabled. 0 means NOT being disabled
	IsDisabled int32 `json:"is_disabled,omitempty"`
	// Is deposit disabled. 0 means not
	IsDepositDisabled int32 `json:"is_deposit_disabled,omitempty"`
	// Is withdrawal disabled. 0 means not
	IsWithdrawDisabled int32 `json:"is_withdraw_disabled,omitempty"`
}
