package starling

import (
	"context"
	"net/http"
)

// Account represents bank account details
type Account struct {
	UID           string `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	SortCode      string `json:"sortCode"`
	Currency      string `json:"currency"`
	IBAN          string `json:"iban"`
	BIC           string `json:"bic"`
	CreatedAt     string `json:"createdAt"`
}

// AccountSummary represents the basic account details
type AccountSummary struct {
	UID             string `json:"accountUid"`
	DefaultCategory string `json:"defaultCategory"`
	Currency        string `json:"currency"`
	CreatedAt       string `json:"createdAt"`
}

// Accounts is a list containing all accounts for a customer
type accounts struct {
	Accounts []AccountSummary `json:"accounts"`
}

// Accounts returns the the accounts held by the current user.
func (c *Client) Accounts(ctx context.Context) ([]AccountSummary, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v2/accounts", nil)
	if err != nil {
		return nil, nil, err
	}

	var acts accounts
	resp, err := c.Do(ctx, req, &acts)
	if err != nil {
		return nil, resp, err
	}
	return acts.Accounts, resp, nil
}

// AccountID represents the identifiers for an individual account
type AccountID struct {
	ID     string `json:"accountIdentifier"`
	BankID string `json:"bankIdentifier"`
	IBAN   string `json:"iban"`
	BIC    string `json:"bic"`
}

// AccountID returns the identifiers for an individual account
func (c *Client) AccountID(ctx context.Context, accountUID string) (*AccountID, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v2/accounts/"+accountUID+"/identifiers", nil)
	if err != nil {
		return nil, nil, err
	}

	var actID *AccountID
	resp, err := c.Do(ctx, req, &actID)
	return actID, resp, err
}

// Balance represents the balance on an account
type Balance struct {
	Cleared     Amount `json:"clearedBalance"`
	Effective   Amount `json:"effectiveBalance"`
	PendingTxns Amount `json:"pendingTransactions"`
	Overdraft   Amount `json:"acceptedOverdraft"`
	Amount      Amount `json:"amount"`
}

// AccountBalance returns the the account balance for the current customer.
func (c *Client) AccountBalance(ctx context.Context, accountUID string) (*Balance, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v2/accounts/"+accountUID+"/balance", nil)
	if err != nil {
		return nil, nil, err
	}

	var b *Balance
	resp, err := c.Do(ctx, req, &b)
	if err != nil {
		return b, resp, err
	}
	return b, resp, nil
}
