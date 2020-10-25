package money

import "fmt"

type Money struct {
	Amount         float64
	CurrencySymbol string
}

func (m *Money) Plus(add Money) (Money, error) {
	if m.CurrencySymbol != add.CurrencySymbol {
		return Money{}, fmt.Errorf("different currencies")
	}
	return Money{
		Amount:         m.Amount + add.Amount,
		CurrencySymbol: m.CurrencySymbol,
	}, nil
}

func (m *Money) Sub(sub Money) (Money, error) {
	if m.CurrencySymbol != sub.CurrencySymbol {
		return Money{}, fmt.Errorf("different currencies")
	}
	return Money{
		Amount:         m.Amount - sub.Amount,
		CurrencySymbol: m.CurrencySymbol,
	}, nil
}
