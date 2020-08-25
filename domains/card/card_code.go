package card

import (
	"errors"
	"fmt"
)

var suitsByCode = map[string]string{
	"C": "CLUBS",
	"D": "DIAMONDS",
	"H": "HEARTS",
	"S": "SPADES",
}

var valuesByCode = map[string]string{
	"A":  "ACE",
	"1":  "1",
	"2":  "2",
	"3":  "3",
	"4":  "4",
	"5":  "5",
	"6":  "6",
	"7":  "7",
	"8":  "8",
	"9":  "9",
	"10": "10",
	"J":  "JACK",
	"Q":  "QUEEN",
	"K":  "KING",
}

type code struct {
	value string
	suit  string
}

// FromCode returns a new card whose values are based
// on the code provided.
func FromCode(s string) (Card, error) {
	var err error
	var card Card

	cardCode, err := parseCode(s)

	if err != nil {
		return card, fmt.Errorf("card.FromCode: %v", err)
	}

	card.Value = valuesByCode[cardCode.value]
	card.Suit = suitsByCode[cardCode.suit]
	card.Code = s

	return card, err
}

func parseCode(s string) (code, error) {
	var cardCode code
	var err error

	if len(s) == 0 {
		err = errors.New("parseCode: received an empty string")
		return cardCode, err
	}

	sLastIndex := len(s) - 1

	value := s[:sLastIndex]
	suit := s[sLastIndex:]

	cardCode.value = value
	cardCode.suit = suit

	if !IsValidCode(cardCode) {
		err = errors.New("parseCode " + s + ": the provided card code is not valid")
	}

	return cardCode, err
}

// IsValidCode validates if the code is valid
func IsValidCode(cardCode code) bool {
	_, hasValue := valuesByCode[cardCode.value]
	_, hasSuit := suitsByCode[cardCode.suit]

	return hasValue && hasSuit
}
