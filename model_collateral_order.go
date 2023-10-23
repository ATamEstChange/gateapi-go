/*
 * Gate API v4
 *
 * Welcome to Gate.io API  APIv4 provides spot, margin and futures trading operations. There are public APIs to retrieve the real-time market statistics, and private APIs which needs authentication to trade on user's behalf.
 *
 * Contact: support@mail.gate.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gateapi

// Collateral Order
type CollateralOrder struct {
	// Order ID
	OrderId int64 `json:"order_id,omitempty"`
	// Collateral
	CollateralCurrency string `json:"collateral_currency,omitempty"`
	// Collateral amount
	CollateralAmount string `json:"collateral_amount,omitempty"`
	// Borrowed currency
	BorrowCurrency string `json:"borrow_currency,omitempty"`
	// Borrowing amount
	BorrowAmount string `json:"borrow_amount,omitempty"`
	// Repaid amount
	RepaidAmount string `json:"repaid_amount,omitempty"`
	// Repaid principal
	RepaidPrincipal string `json:"repaid_principal,omitempty"`
	// Repaid interest
	RepaidInterest string `json:"repaid_interest,omitempty"`
	// The initial collateralization rate
	InitLtv string `json:"init_ltv,omitempty"`
	// The current collateralization rate
	CurrentLtv string `json:"current_ltv,omitempty"`
	// The liquidation collateralization rate
	LiquidateLtv string `json:"liquidate_ltv,omitempty"`
	// Order status: - initial: Initial state after placing the order - collateral_deducted: Collateral deduction successful - collateral_returning: Loan failed - Collateral return pending - lent: Loan successful - repaying: Repayment in progress - liquidating: Liquidation in progress - finished: Order completed - closed_liquidated: Liquidation and repayment completed
	Status string `json:"status,omitempty"`
	// Borrowing time, timestamp in seconds
	BorrowTime int64 `json:"borrow_time,omitempty"`
	// Outstanding principal and interest (outstanding principal + outstanding interest)
	LeftRepayTotal string `json:"left_repay_total,omitempty"`
	// outstanding principal
	LeftRepayPrincipal string `json:"left_repay_principal,omitempty"`
	// outstanding interest
	LeftRepayInterest string `json:"left_repay_interest,omitempty"`
}
