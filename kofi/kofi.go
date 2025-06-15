package kofi

import (
	"encoding/json"
	"net/http"
)

type PaymentType string

// Possible payment types
const (
	Commission   PaymentType = "Commission"
	Donation     PaymentType = "Donation"
	ShopOrder    PaymentType = "Shop Order"
	Subscription PaymentType = "Subscription"
)

type ShopItem struct {
	DirectLinkCode string `json:"direct_link_code"`
	Quantity       uint32 `json:"quantity"`
	VariationName  string `json:"variation_name"`
}

type Payload struct {
	// The verification token below will be included with every request to your webhook.
	VerificationToken          string      `json:"verification_token"`
	MessageID                  string      `json:"message_id"`
	Timestamp                  string      `json:"timestamp"`
	Type                       PaymentType `json:"type"`
	IsPublic                   bool        `json:"is_public"`
	FromName                   string      `json:"from_name"`
	Message                    string      `json:"message"`
	Amount                     string      `json:"amount"`
	Url                        string      `json:"url"`
	Email                      string      `json:"email"`
	Currency                   string      `json:"currency"`
	IsSubscriptionPayment      bool        `json:"is_subscription_payment"`
	IsFirstSubscriptionPayment bool        `json:"is_first_subscription_payment"`
	KoFiTransactionId          string      `json:"kofi_transaction_id"`
	ShopItems                  []ShopItem  `json:"shop_items"`
	TierName                   string      `json:"tier_name"`
	Shipping                   struct {
		FullName        string `json:"full_name"`
		StreetAddress   string `json:"street_address"`
		City            string `json:"city"`
		StateOrProvince string `json:"state_or_province"`
		PostalCode      string `json:"postal_code"`
		Country         string `json:"country"`
		CountryCode     string `json:"country_code"`
		Telephone       string `json:"telephone"`
	} `json:"shipping"`
}

func KofiWebhook(w http.ResponseWriter, req *http.Request) {
	// The webhook should only respond to POST requests
	if req.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)

		return
	}

	// The webhook should only accept the following content type
	if contentType := req.Header.Get("Content-Type"); contentType != "application/x-www-form-urlencoded" {
		http.Error(w, "", http.StatusUnsupportedMediaType)

		return
	}

	if err := req.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)

		return
	}

	jsonData := req.FormValue("data")
	if jsonData == "" {
		http.Error(w, "Missing `data` field", http.StatusBadRequest)

		return
	}

	var payload Payload
	if err := json.Unmarshal([]byte(jsonData), &payload); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusOK)
}
