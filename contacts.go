package starling

import (
	"context"
	"net/http"
)

// Contact represents the details of a payee
type Contact struct {
	UID  string `json:"id"`
	Name string `json:"name"`
}

// Contacts are a list of payees
type Contacts struct {
	Contacts []Contact
}

// HALContacts is a HAL wrapper around the Contacts type.
type HALContacts struct {
	Links    struct{}  `json:"_links"`
	Embedded *Contacts `json:"_embedded"`
}

// ContactAccount holds payee account details
type ContactAccount struct {
	UID           string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	AccountNumber string `json:"accountNumber"`
	SortCode      string `json:"sortCode"`
}

// ContactAccounts holds a list of accounts for a payee
type ContactAccounts struct {
	ContactAccounts []ContactAccount `json:"contactAccounts"`
}

// GetContacts returns the contacts for the current customer.
func (c *Client) GetContacts(ctx context.Context) (*Contacts, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v1/contacts", nil)
	if err != nil {
		return nil, nil, err
	}

	var halResp *HALContacts
	var contacts *Contacts
	resp, err := c.Do(ctx, req, &halResp)
	if err != nil {
		return contacts, resp, err
	}

	if halResp.Embedded != nil {
		contacts = halResp.Embedded
	}

	return contacts, resp, nil
}

// GetContact returns an individual contact for the current customer.
func (c *Client) GetContact(ctx context.Context, uid string) (*Contact, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v1/contacts/"+uid, nil)
	if err != nil {
		return nil, nil, err
	}

	contact := new(Contact)
	resp, err := c.Do(ctx, req, contact)
	return contact, resp, err
}

// DeleteContact deletes an individual contact for the current customer. It returns http.StatusNoContent
// on success. No payload is returned.
func (c *Client) DeleteContact(ctx context.Context, uid string) (*http.Response, error) {
	req, err := c.NewRequest("DELETE", "/api/v1/contacts/"+uid, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}

// GetContactAccounts returns the accounts for a given contact.
func (c *Client) GetContactAccounts(ctx context.Context, uid string) (*ContactAccounts, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v1/contacts/"+uid+"/accounts", nil)
	if err != nil {
		return nil, nil, err
	}

	var cas *ContactAccounts
	resp, err := c.Do(ctx, req, &cas)
	return cas, resp, err
}

// GetContactAccount returns the specified account for a given contact.
func (c *Client) GetContactAccount(ctx context.Context, cUID, aUID string) (*ContactAccount, *http.Response, error) {
	req, err := c.NewRequest("GET", "/api/v1/contacts/"+cUID+"/accounts/"+aUID, nil)
	if err != nil {
		return nil, nil, err
	}

	var ca *ContactAccount
	resp, err := c.Do(ctx, req, &ca)
	return ca, resp, nil
}

// PostContactAccount returns the specified account for a given contact.
func (c *Client) PostContactAccount(ctx context.Context, ca ContactAccount) (*http.Response, error) {
	req, err := c.NewRequest("POST", "/api/v1/contacts", ca)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(ctx, req, nil)
	return resp, err
}
