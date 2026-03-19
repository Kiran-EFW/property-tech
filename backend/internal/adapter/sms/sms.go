package sms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/zerolog/log"
)

// SMSProvider defines the interface for sending SMS messages.
type SMSProvider interface {
	SendSMS(phone, message string) error
}

// TwilioProvider sends SMS messages via the Twilio API.
type TwilioProvider struct {
	accountSID string
	authToken  string
	fromNumber string
	httpClient *http.Client
}

// NewTwilioProvider creates a new TwilioProvider.
func NewTwilioProvider(accountSID, authToken, fromNumber string) *TwilioProvider {
	return &TwilioProvider{
		accountSID: accountSID,
		authToken:  authToken,
		fromNumber: fromNumber,
		httpClient: &http.Client{},
	}
}

// SendSMS sends an SMS message via Twilio.
func (t *TwilioProvider) SendSMS(phone, message string) error {
	twilioURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.accountSID)

	data := url.Values{}
	data.Set("To", phone)
	data.Set("From", t.fromNumber)
	data.Set("Body", message)

	req, err := http.NewRequest(http.MethodPost, twilioURL, strings.NewReader(data.Encode()))
	if err != nil {
		return fmt.Errorf("twilio: failed to create request: %w", err)
	}

	req.SetBasicAuth(t.accountSID, t.authToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("twilio: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var errResp struct {
			Message string `json:"message"`
			Code    int    `json:"code"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err == nil {
			return fmt.Errorf("twilio: API error (code %d): %s", errResp.Code, errResp.Message)
		}
		return fmt.Errorf("twilio: API returned status %d", resp.StatusCode)
	}

	log.Info().
		Str("to", phone).
		Msg("SMS sent via Twilio")

	return nil
}

// NoopSMSProvider is a no-op implementation for development and testing.
type NoopSMSProvider struct{}

// SendSMS logs the message without actually sending it.
func (n *NoopSMSProvider) SendSMS(phone, message string) error {
	log.Info().
		Str("provider", "sms-noop").
		Str("to", phone).
		Str("message", message).
		Msg("SMS (noop)")
	return nil
}
