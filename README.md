# go-paypal-fee-computer
Compute paypal fees, see list of fees here: https://www.paypal.com/en/webapps/mpp/paypal-fees

## Example

```go
package main

import (
	"fmt"
	"github.com/techknowlogick/go-paypal-fee-computer"
)

func main() {
	sending := 5.00
	transaction_percentage_fee := 2.9
	fixed_fee := 0.3
	total, _ := paypal_fee_computer.Compute(sending, transaction_percentage_fee, fixed_fee)

	fmt.Printf("To send $%.2f it will cost you $%.2f", sending, total)
	// output: To send $5.00 it will cost you $5.46

}
```

## Prior Art
http://strk.kbt.io/utils/paypalfeecomputer.html
