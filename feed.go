package starling

import (
	"context"
	"net/http"
	"time"
)

// Feed is a slice of Items representing customer transactions
type feed struct {
	Items []FeedItem `json:"feedItems"`
}

// Item is a single customer transaction in their feed
type FeedItem struct {
	FeedItemUID                       string      `json:"feedItemUid"`
	CategoryUID                       string      `json:"categoryUid"`
	AccountUID                        string      `json:"accountUid"`
	Amount                            Amount      `json:"amount"`
	SourceAmount                      Amount      `json:"sourceAmount"`
	Direction                         string      `json:"direction"`
	UpdatedAt                         time.Time   `json:"updatedAt"`
	TransactionTime                   time.Time   `json:"transactionTime"`
	SettlementTime                    time.Time   `json:"settlementTime"`
	RetryAllocationUntilTime          time.Time   `json:"retryAllocationUntilTime"`
	Source                            string      `json:"source"`
	SourceSubType                     string      `json:"sourceSubType"`
	Status                            string      `json:"status"`
	TransactionApplicationUserUID     string      `json:"transactionApplicationUserUid"`
	CounterPartyType                  string      `json:"counterPartyType"`
	CounterPartyUID                   string      `json:"counterPartyUid"`
	CounterPartyName                  string      `json:"counterPartyName"`
	CounterPartySubEntityUID          string      `json:"counterPartySubEntityUid"`
	CounterPartySubEntityName         string      `json:"counterPartySubEntityName"`
	CounterPartySubEntityIdentifier   string      `json:"counterPartySubEntityIdentifier"`
	CounterPartSubEntitySubIdentifier string      `json:"counterPartSubEntitySubIdentifier"`
	ExchangeRate                      float64     `json:"exchangeRate"`
	TotalFees                         float64     `json:"totalFees"`
	TotalFeeAmount                    Amount      `json:"totalFeeAmount"`
	Reference                         string      `json:"reference"`
	Country                           string      `json:"country"`
	SpendingCategory                  string      `json:"spendingCategory"`
	UserNote                          string      `json:"userNote"`
	RoundUp                           FeedRoundUp `json:"roundUp"`
	HasAttachment                     bool        `json:"hasAttachment"`
	ReceiptPresent                    bool        `json:"receiptPresent"`
}

type FeedRoundUp struct {
	GoalCategoryUID string `json:"goalCategoryUid"`
	Amount          Amount `json:"amount"`
}

// Feed returns a slice of Items for a given account and category. It returns an error if unable
// to retrieve the feed.
func (c *Client) Feed(ctx context.Context, act, cat string, since time.Time) ([]FeedItem, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v2/feed/account/"+act+"/category/"+cat, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Add("changesSince", since.Format(time.RFC3339Nano))
	req.URL.RawQuery = q.Encode()

	var f feed
	resp, err := c.Do(ctx, req, &f)
	if err != nil {
		return nil, resp, err
	}
	return f.Items, resp, nil
}

// FeedItem returns a feed Item for a given account and category. It returns an error if unable to
// retrieve the feed Item.
// Note: FeedItem uses the v2 API which is still under active development.
func (c *Client) FeedItem(ctx context.Context, act, cat, itm string) (*FeedItem, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v2/feed/account/"+act+"/category/"+cat+"/"+itm, nil)
	if err != nil {
		return nil, nil, err
	}

	var i FeedItem
	resp, err := c.Do(ctx, req, &i)
	if err != nil {
		return nil, resp, err
	}
	return &i, resp, nil
}
