package starling

import (
	"context"
	"net/http"
	"strings"
	"time"
)

// Cards represents a collection of cards.
type cards struct {
	Cards []Card `json:"cards"`
}

// Card represents card details
type Card struct {
	CardUID                   string         `json:"cardUid"`
	PublicToken               string         `json:"publicToken"`
	Enabled                   bool           `json:"enabled"`
	Cancelled                 bool           `json:"cancelled"`
	ActivationRequested       bool           `json:"activationRequested"`
	Activated                 bool           `json:"activated"`
	WalletNotificationsEnaled bool           `json:"walletNotificationsEnabled"`
	PosEnabled                bool           `json:"posEnabled"`
	AtmEnabled                bool           `json:"atmEnabled"`
	OnlineEnabled             bool           `json:"onlineEnabled"`
	MobileWalletEnabled       bool           `json:"mobileWalletEnabled"`
	GamblingEnabled           bool           `json:"gamblingEnabled"`
	MagStripeEnabled          bool           `json:"magStripeEnabled"`
	EndOfCardNumber           string         `json:"endOfCardNumber"`
	CurrencyFlags             []CurrencyFlag `json:"currencyFlags"`
	CardAssociationUID        string         `json:"cardAssociationUid"`
	GamblingToBeEnabledAt     time.Time      `json:"gamblingToBeEnabledAt"`
}

type CurrencyFlag struct {
	Enabled  bool   `json:"enabled"`
	Currency string `json:"currency"`
}

type enabledRequest struct {
	Enabled bool `json:"enabled"`
}

// Cards returns a list of cards.
func (c *Client) Cards(ctx context.Context) ([]Card, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v2/cards", nil)
	if err != nil {
		return nil, nil, err
	}

	var cards cards
	resp, err := c.Do(ctx, req, &cards)
	if err != nil {
		return cards.Cards, resp, err
	}

	return cards.Cards, resp, nil
}

// EnableCard enables a card.
func (c *Client) EnableCard(ctx context.Context, cardUID string, en bool) (*http.Response, error) {
	req, err := c.NewRequest("PUT", "/api/v2/cards/"+cardUID+"/controls/enabled", enabledRequest{Enabled: en})
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// EnableCardOption enables a specific card option with the list of valid options being:
// atm, gambling, mag-stripe, mobile-wallet, online, pos
func (c *Client) EnableCardOption(ctx context.Context, cardUID, option string, en bool) (*http.Response, error) {
	req, err := c.NewRequest("PUT", "/api/v2/cards/"+cardUID+"/controls/"+strings.ToLower(option)+"-enabled", enabledRequest{Enabled: en})
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
