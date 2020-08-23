package card_test

import (
	"testing"

	"github.com/eduardogspereira/deck-api/card"
)

func TestNewCard(t *testing.T) {
	c, _ := card.FromCode("QH")

	expectedCardCode := "QH"
	if c.Code != expectedCardCode {
		t.Errorf("c.Code = %v, want %v", c.Code, expectedCardCode)
	}

	expectedCardValue := "QUEEN"
	if c.Value != expectedCardValue {
		t.Errorf("card.Value = %v, want %v", c.Value, expectedCardValue)
	}

	expectedCardSuit := "HEARTS"
	if c.Suit != expectedCardSuit {
		t.Errorf("card.Suit = %v, want %v", c.Suit, expectedCardSuit)
	}

	c, _ = card.FromCode("4D")

	expectedCardCode = "4D"
	if c.Code != expectedCardCode {
		t.Errorf("card.Code = %v, want %v", c.Code, expectedCardCode)
	}

	expectedCardValue = "4"
	if c.Value != expectedCardValue {
		t.Errorf("card.Value = %v, want %v", c.Value, expectedCardValue)
	}

	expectedCardSuit = "DIAMONDS"
	if c.Suit != expectedCardSuit {
		t.Errorf("card.Suit = %v, want %v", c.Suit, expectedCardSuit)
	}

	c, _ = card.FromCode("10S")

	expectedCardCode = "10S"
	if c.Code != expectedCardCode {
		t.Errorf("card.Code = %v, want %v", c.Code, expectedCardCode)
	}

	expectedCardValue = "10"
	if c.Value != expectedCardValue {
		t.Errorf("card.Value = %v, want %v", c.Value, expectedCardValue)
	}

	expectedCardSuit = "SPADES"
	if c.Suit != expectedCardSuit {
		t.Errorf("card.Suit = %v, want %v", c.Suit, expectedCardSuit)
	}

	_, err := card.FromCode("EZ")
	if err == nil {
		t.Errorf("expect error not to be nil")
	}

	_, err = card.FromCode("")
	if err == nil {
		t.Errorf("expect error not to be nil")
	}
}
