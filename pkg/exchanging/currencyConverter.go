package exchanging

import "errors"

// CurrencyConverter provides a service for converting money from one currency to another
type CurrencyConverter struct {
	rate ExchangeRater
}

// ExchangeRater represents service returns current exchange rate srt -> dst
type ExchangeRater interface {
	Rate(src string, dst string) (float32, error)
}

var (
	// ErrSrcCurrensyNotFound rises when source currency is not found
	ErrSrcCurrensyNotFound = errors.New("Source currency is not found")
	// ErrDstCurrensyNotFound rises when destination currency is not found
	ErrDstCurrensyNotFound = errors.New("Destination currency is not found")
	// ErrRateUnavailable rises when a exchange rate service is unavailable
	ErrRateUnavailable = errors.New("Exchange rate service is unavailable")
)

// NewCurrencyConverter creates a new CurrencyConverter instance
func NewCurrencyConverter(er ExchangeRater) *CurrencyConverter {
	return &CurrencyConverter{
		rate: er,
	}
}

// Convert converts given amount of money to desired curreyncy
func (cc *CurrencyConverter) Convert(amount float32, src string, dst string) (float32, error) {
	r, err := cc.rate.Rate(src, dst)
	if err != nil {
		return 0, err
	}

	return amount * r, nil
}
