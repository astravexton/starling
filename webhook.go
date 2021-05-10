package starling

import (
	"bytes"
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// WebHookPayload defines the structure of the Starling web hook payload
type WebHookPayload struct {
	WebhookEventUID  string          `json:"webhookEventUid"`
	EventTimestamp   time.Time       `json:"eventTimestamp"`
	Content          WebHookFeedItem `json:"content"`
	AccountHolderUID string          `json:"accountHolderUid"`
}

// WebHookFeedItem defines the structure of the Starling web hook feed item
type WebHookFeedItem struct {
	FeedItem
	AccountUID            string             `json:"accountUid"`
	FeedItemFailureReason string             `json:"feedItemFailureReason"`
	MasterCardFeedDetails MasterCardFeedItem `json:"masterCardFeedDetails"`
}

// MasterCardFeedItem defines the structure of the MasterCard feed item
type MasterCardFeedItem struct {
	MerchantIdentifier string    `json:"merchantIdentifier"`
	MCC                int32     `json:"mcc"`
	PosTimestamp       time.Time `json:"posTimestamp"`
	CardLast4          string    `json:"cardLast4"`
}

// Validate takes an http request and a base64-encoded web-hook public key
// and validates the request signature matches the signature provided in
// the X-Hook-Signature. An error is returned if unable to parse the body
// of the request.
func Validate(r *http.Request, publicKey string) (bool, error) {
	if r.Body == nil {
		return false, fmt.Errorf("no body to validate")
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return false, err
	}

	key, err := publicKeyFrom64(publicKey)
	if err != nil {
		return false, err
	}
	reqSig, err := base64.StdEncoding.DecodeString(r.Header.Get("X-Hook-Signature"))
	if err != nil {
		return false, err
	}

	body := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = body

	digest := sha512.Sum512(buf)
	err = rsa.VerifyPKCS1v15(key, crypto.SHA512, digest[:], reqSig)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Convert the base64 encoded public key to *rsa.PublicKey
func publicKeyFrom64(key string) (*rsa.PublicKey, error) {
	b, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, err
	}
	pubInterface, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, err
	}

	return pub, nil
}
