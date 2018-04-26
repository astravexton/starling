package starling

import "time"

// ErrorDetail holds the details of an error message
type ErrorDetail struct {
	Message string
}

// Amount represents the value and currency of a monetary amount
type Amount struct {
	Currency   string `json:"currency"`   // ISO-4217 3 character currency code
	MinorUnits int64  `json:"minorUnits"` // Amount in the minor units of the given currency; eg pence in GBP, cents in EUR
}

// RecurrenceRule defines the pattern for recurring events
type RecurrenceRule struct {
	StartDate string `json:"startDate"`
	Frequency string `json:"frequency"`
	Interval  int32  `json:"interval,omitempty"`
	Count     int32  `json:"count,omitempty"`
	UntilDate string `json:"untilDate,omitempty"`
	WeekStart string `json:"weekStart"`
}

// Photo is a photo associated to a savings goal
type Photo struct {
	Base64EncodedPhoto string `json:"base64EncodedPhoto"` // A text (base 64) encoded picture to associate with the savings goal
}

// MastercardTransactionPayload is the webhook payload for mastercard transactions
type MastercardTransactionPayload struct {
	WebhookNotificationUID string          `json:"webhookNotificationUid"` // Unique identifier of the webhook dispatch event
	CustomerUID            string          `json:"customerUid"`            // Unique identifier of the customer
	WebhookType            string          `json:"webhookType"`            // The type of the event
	EventUID               string          `json:"eventUid"`               // Unique identifier of the customer transaction event
	TransactionAmount      Amount          `json:"transactionAmount"`
	SourceAmount           Amount          `json:"sourceAmount"`
	Direction              string          `json:"direction"`            // The cashflow direction of the card transaction
	Description            string          `json:"description"`          // The transaction description, usually the name of the merchant
	MerchantUID            string          `json:"merchantUid"`          // The unique identifier of the merchant
	MerchantLocationUID    string          `json:"merchantLocationUid"`  // The unique identifier of the merchant location
	Status                 string          `json:"status"`               // The status of the transaction
	TransactionMethod      string          `json:"transactionMethod"`    // The method of card usage
	TransactionTimestamp   string          `json:"transactionTimestamp"` // Timestamp of the card transaction
	MerchantPosData        MerchantPosData `json:"merchantPosData"`
}

// MerchantPosData is data relating to the merchant at the point-of-sale
type MerchantPosData struct {
	PosTimestamp       string `json:"posTimestamp"`       // The transaction time as reported at the point of sale
	CardLast4          string `json:"cardLast4"`          // The last 4 digits of the mastercard PAN
	AuthorisationCode  string `json:"authorisationCode"`  // The authorisation code of the transaction, as reported at the point of sale
	Country            string `json:"country"`            // The country of the transaction, in ISO-3 format
	MerchantIdentifier string `json:"merchantIdentifier"` // The merchant identifier as reported by Mastercard AKA mid
}

// HALLink is a link to another resource
type HALLink struct {
	HREF        string `json:"href"`
	Templated   bool   `json:"templated"`
	Type        string `json:"type"`
	Deprecation string `json:"deprecation"`
	Name        string `json:"name"`
	Profile     string `json:"profile"`
	Title       string `json:"title"`
	HREFLang    string `json:"hreflang"`
}

// Merchant represents details of a merchant
type Merchant struct {
	UID             string `json:"merchantUid"`
	Name            string `json:"name"`
	Website         string `json:"website"`
	PhoneNumber     string `json:"phoneNumber"`
	TwitterUsername string `json:"twitterUsername"`
}

// MerchantLocation represents details of a merchant location
type MerchantLocation struct {
	UID                            string  `json:"merchantLocationUid"`
	MerchantUID                    string  `json:"merchantUid"`
	Merchant                       HALLink `json:"merchant"`
	MerchantName                   string  `json:"merchantName"`
	LocationName                   string  `json:"locationName"`
	Address                        string  `json:"address"`
	PhoneNumber                    string  `json:"phoneNUmber"`
	GooglePlaceID                  string  `json:"googlePlaceId"`
	MastercardMerchantCategoryCode int32   `json:"mastercardMerchantCategoryCode"`
}

// PaymentAmount represents the currency and amount of a payment
type PaymentAmount struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

// ScheduledPayment represents a scheduled payment
type ScheduledPayment struct {
	LocalPayment
	RecurrenceRule RecurrenceRule `json:"recurrenceRule"`
}

// LocalPayment represents a local payment
type LocalPayment struct {
	Payment               PaymentAmount `json:"payment"`
	DestinationAccountUID string        `json:"destinationAccountUid"`
	Reference             string        `json:"reference"`
}

// PaymentOrder is a single PaymentOrder
type PaymentOrder struct {
	UID                        string         `json:"paymentOrderId"`
	Currency                   string         `json:"currency"`
	Amount                     float64        `json:"amount"`
	Reference                  string         `json:"reference"`
	ReceivingContactAccountUID string         `json:"receivingContactAccountId"`
	RecipientName              string         `json:"recipientName"`
	Immediate                  bool           `json:"immediate"`
	RecurrenceRule             RecurrenceRule `json:"recurrenceRule"`
	StartDate                  string         `json:"startDate"`
	NextDate                   string         `json:"nextDate"`
	CancelledAt                string         `json:"cancelledAt"`
	PaymentType                string         `json:"paymentType"`
	MandateUID                 string         `json:"mandateId"`
}

// PaymentOrders is a list of PaymentOrders
type PaymentOrders struct {
	NextPage      HALLink        `json:"nextPage"`
	PaymentOrders []PaymentOrder `json:"paymentOrders"`
}

// SpendingCategory is the category associated with a transaction
type SpendingCategory struct {
	SpendingCategory string `json:"spendingCategory"`
}

// Receipt is a receipt for a transaction
type Receipt struct {
	UID                string        `json:"receiptUid"`
	EventUID           string        `json:"eventUid"`
	MetadataSource     string        `json:"metadataSource"`
	ReceiptIdentifier  string        `json:"receiptIdentifier"`
	MerchantIdentifier string        `json:"merchantIdentifier"`
	TotalAmount        float64       `json:"totalAmount"`
	TotalTax           float64       `json:"totalTax"`
	AuthCode           string        `json:"authCode"`
	CardLast4          string        `json:"cardLast4"`
	Items              []ReceiptItem `json:"items"`
}

// ReceiptItem is a single item on a Receipt
type ReceiptItem struct {
	UID         string  `json:"receiptItemUid"`
	Description string  `json:"description"`
	Quantity    int32   `json:"quantity"`
	Amount      float64 `json:"amount"`
	Tax         float64 `json:"tax"`
}

// ReceiptUID is an un-used type
type ReceiptUID struct{}

// OptionalTransactionSummary indicates the presence of a TransactionSummary
type OptionalTransactionSummary optional

type optional struct {
	Present bool `json:"present"`
}

// DateRange holds two dates that represent a range. It is typically
// used when providing a range when querying the API.
type DateRange struct {
	From time.Time
	To   time.Time
}
