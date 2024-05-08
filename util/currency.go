package util


// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	TRY = "TRY"
	BYN = "BYN"
)


//IsSupportedCurrency returns true if the currency is supported

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, TRY, BYN:
		return true
	}
	return false
}
