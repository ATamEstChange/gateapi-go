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

type LoanRecord struct {
	// Loan record ID
	Id string `json:"id,omitempty"`
	// Loan ID the record belongs to
	LoanId string `json:"loan_id,omitempty"`
	// Loan time
	CreateTime string `json:"create_time,omitempty"`
	// Expiration time
	ExpireTime string `json:"expire_time,omitempty"`
	// Loan record status
	Status string `json:"status,omitempty"`
	// Garbled user ID
	BorrowUserId string `json:"borrow_user_id,omitempty"`
	// Loan currency
	Currency string `json:"currency,omitempty"`
	// Loan rate
	Rate string `json:"rate,omitempty"`
	// Loan amount
	Amount string `json:"amount,omitempty"`
	// Loan days
	Days int32 `json:"days,omitempty"`
	// Whether the record will auto renew on expiration
	AutoRenew bool `json:"auto_renew,omitempty"`
	// Repaid amount
	Repaid string `json:"repaid,omitempty"`
	// Repaid interest
	PaidInterest string `json:"paid_interest,omitempty"`
	// Interest not repaid
	UnpaidInterest string `json:"unpaid_interest,omitempty"`
}
