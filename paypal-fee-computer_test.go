package paypal_fee_computer

import "testing"

func TestCompute(t *testing.T) {

	cases := []struct{ amount, percentage, fixed, want float64 }{
		{0.01, 2.9, 0.3, 0.32 },
		{0.1, 2.9, 0.3, 0.41 },
		{1.0, 2.9, 0.3, 1.34 },
		{10.0, 2.9, 0.3, 10.61 },
		{25.0, 2.9, 0.3, 26.06 },
		{100.0, 2.9, 0.3, 103.30 },
		{500.0, 2.9, 0.3, 515.24 },
		{1000.0, 2.9, 0.3, 1030.18 },
	}

	for _, c := range cases {
		got, err := Compute(c.amount, c.percentage, c.fixed)
		if err != nil {
			got = -1.0
		}
		if got != c.want {
			t.Errorf("Compute(%v, %v, %v) == %v, expected %v", c.amount, c.percentage, c.fixed, got, c.want)
		}
	}

}