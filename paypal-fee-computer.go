package paypal_fee_computer // import "src.techknowlogick.com/paypal-fee-computer"

import "math"

func Compute(amount, percentage, fixed float64) (float64, error) {
	// no error returned yet, just there as in future values may be validated
	return toFixed((amount + fixed) / (1.0 - percentage/100), 2), nil
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(math.Round(num * output)) / output
}
