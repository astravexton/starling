package starling

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var cardTestCases = []struct {
	name string
	mock string
}{
	{
		name: "sample cards",
		mock: `{
			"cards": [
				{
					"cardUid": "ddeeddee-ddee-ddee-ddee-ddeeddeeddee",
					"publicToken": "123456789",
					"enabled": true,
					"walletNotificationEnabled": true,
					"posEnabled": true,
					"atmEnabled": true,
					"onlineEnabled": true,
					"mobileWalletEnabled": true,
					"gamblingEnabled": true,
					"magStripeEnabled": true,
					"cancelled": true,
					"activationRequested": true,
					"activated": true,
					"endOfCardNumber": "59312",
					"currencyFlags": [
						{
							"enabled": true,
							"currency": "string"
						}
					],
					"cardAssociationUid": "aaaaaaaa-aaaa-4aaa-aaaa-aaaaaaaaaaaa",
					"gamblingToBeEnabledAt": "2021-05-10T13:34:22.322Z"
				}
			]
		}`,
	},
	{
		name: "multiple sample cards",
		mock: `{
			"cards": [
				{
					"cardUid": "ddeeddee-ddee-ddee-ddee-ddeeddeeddee",
					"publicToken": "123456789",
					"enabled": true,
					"walletNotificationEnabled": true,
					"posEnabled": true,
					"atmEnabled": true,
					"onlineEnabled": true,
					"mobileWalletEnabled": true,
					"gamblingEnabled": true,
					"magStripeEnabled": true,
					"cancelled": true,
					"activationRequested": true,
					"activated": true,
					"endOfCardNumber": "59312",
					"currencyFlags": [
						{
							"enabled": true,
							"currency": "string"
						}
					],
					"cardAssociationUid": "aaaaaaaa-aaaa-4aaa-aaaa-aaaaaaaaaaaa",
					"gamblingToBeEnabledAt": "2021-05-10T13:34:22.322Z"
				},
				{
					"cardUid": "ddeeddee-ddee-ddee-ddee-ddeeddeeddee",
					"publicToken": "987654321",
					"enabled": true,
					"walletNotificationEnabled": true,
					"posEnabled": true,
					"atmEnabled": true,
					"onlineEnabled": true,
					"mobileWalletEnabled": true,
					"gamblingEnabled": true,
					"magStripeEnabled": true,
					"cancelled": true,
					"activationRequested": true,
					"activated": true,
					"endOfCardNumber": "59312",
					"currencyFlags": [
						{
							"enabled": true,
							"currency": "string"
						}
					],
					"cardAssociationUid": "aaaaaaaa-aaaa-4aaa-aaaa-aaaaaaaaaaaa",
					"gamblingToBeEnabledAt": "2021-05-10T13:34:22.322Z"
				}
			]
		}`,
	},
}

func TestCard(t *testing.T) {
	for _, tc := range cardTestCases {
		t.Run(tc.name, func(st *testing.T) {
			testCard(st, tc.name, tc.mock)
		})
	}
}

func testCard(t *testing.T, name, mock string) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v2/cards", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		fmt.Fprint(w, mock)
	})

	got, _, err := client.Cards(context.Background())
	checkNoError(t, err)

	want := &cards{}
	json.Unmarshal([]byte(mock), want)

	if !reflect.DeepEqual(got, want.Cards) {
		t.Error("should return cards matching the mock response", cross)
	}
}

func TestCardForbidden(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v2/cards", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusForbidden)
	})

	got, resp, err := client.Cards(context.Background())
	checkHasError(t, err)

	if resp.StatusCode != http.StatusForbidden {
		t.Error("should return HTTP 403 status")
	}

	if got != nil {
		t.Error("should not return a card")
	}
}

func TestEnableCard(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	cardID := "ddeeddee-ddee-ddee-ddee-ddeeddeeddee"
	mockReq := enabledRequest{Enabled: true}
	mockResp := &statusResponse{}
	mux.HandleFunc("/api/v2/cards/ddeeddee-ddee-ddee-ddee-ddeeddeeddee/controls/enabled", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPut)
		fmt.Fprintln(w, mockResp)

		var ecReq = enabledRequest{}
		err := json.NewDecoder(r.Body).Decode(&ecReq)
		if err != nil {
			t.Fatal("should send a request that the API can parse", cross, err)
		}

		if !reflect.DeepEqual(mockReq, ecReq) {
			t.Error("should send a request that matches the mock", cross)
		}
	})

	resp, err := client.EnableCard(context.Background(), cardID, true)
	if err != nil {
		t.Fatal("should be able to make the request", cross, err)
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("should receive a %d status code %s %d", want, cross, got)
	}
}

func TestEnableCardOption(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	cardID := "ddeeddee-ddee-ddee-ddee-ddeeddeeddee"
	mockReq := enabledRequest{Enabled: true}
	mockResp := &statusResponse{}
	mux.HandleFunc("/api/v2/cards/ddeeddee-ddee-ddee-ddee-ddeeddeeddee/controls/atm-enabled", func(w http.ResponseWriter, r *http.Request) {
		checkMethod(t, r, http.MethodPut)
		fmt.Fprintln(w, mockResp)

		var ecReq = enabledRequest{}
		err := json.NewDecoder(r.Body).Decode(&ecReq)
		if err != nil {
			t.Fatal("should send a request that the API can parse", cross, err)
		}

		if !reflect.DeepEqual(mockReq, ecReq) {
			t.Error("should send a request that matches the mock", cross)
		}
	})

	resp, err := client.EnableCardOption(context.Background(), cardID, "atm", true)
	if err != nil {
		t.Fatal("should be able to make the request", cross, err)
	}

	if got, want := resp.StatusCode, http.StatusOK; got != want {
		t.Errorf("should receive a %d status code %s %d", want, cross, got)
	}
}
