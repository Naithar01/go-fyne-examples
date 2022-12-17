package main

import "testing"

func TestGold_GetPrice(t *testing.T) {
	g := Gold{
		Prices: nil,
	}

	_, err := g.GetPrices()

	if err != nil {
		t.Error(err)
	}
}
