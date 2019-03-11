/*
 * Gate API v4
 *
 * APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * API version: 4.5.1
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

type Ticker struct {
	// Currency pair
	CurrencyPair string `json:"currency_pair,omitempty"`
	// Last trading price
	Last string `json:"last,omitempty"`
	// Lowest ask
	LowestAsk string `json:"lowest_ask,omitempty"`
	// Highest bid
	HighestBid string `json:"highest_bid,omitempty"`
	// Change percentage.
	ChangePercentage string `json:"change_percentage,omitempty"`
	// Base currency trade volume
	BaseVolume string `json:"base_volume,omitempty"`
	// Quote currency trade volume
	QuoteVolume string `json:"quote_volume,omitempty"`
	// Highest price in 24h
	High24h string `json:"high_24h,omitempty"`
	// Lowest price in 24h
	Low24h string `json:"low_24h,omitempty"`
}
