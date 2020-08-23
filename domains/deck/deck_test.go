package deck_test

import (
	"testing"

	"github.com/eduardogspereira/deck-api/domains/deck"
)

func TestNewDeckDefaultValues(t *testing.T) {
	options := deck.Options{}
	d, _ := deck.New(options)

	if d.Shuffled {
		t.Error("expect deck not to be shuffled")
	}

	expectedRemaining := 52
	if d.Remaining() != expectedRemaining {
		t.Errorf("d.Remaining = %v, want %v", d.Remaining(), expectedRemaining)
	}

	expectCardLengths := 52
	if len(d.Cards) != expectCardLengths {
		t.Errorf("len(d.Cards) = %v, want %v", len(d.Cards), expectCardLengths)
	}

	expectedFirstCardCode := "AS"
	if d.Cards[0].Code != expectedFirstCardCode {
		t.Errorf("len(d.Cards) = %v, want %v", d.Cards[0].Code, expectedFirstCardCode)
	}

	expectedLastCardCode := "KH"
	if d.Cards[len(d.Cards)-1].Code != expectedLastCardCode {
		t.Errorf("len(d.Cards) = %v, want %v", d.Cards[len(d.Cards)-1].Code, expectedLastCardCode)
	}
}

func TestNewDeckWithShuffledOption(t *testing.T) {
	options := deck.Options{
		Shuffle: true,
	}
	d, _ := deck.New(options)

	if d.Shuffled {
		t.Error("expect deck not to be shuffled")
	}

	expectedRemaining := 52
	if d.Remaining() != expectedRemaining {
		t.Errorf("d.Remaining = %v, want %v", d.Remaining(), expectedRemaining)
	}

	expectCardLengths := 52
	if len(d.Cards) != expectCardLengths {
		t.Errorf("len(d.Cards) = %v, want %v", len(d.Cards), expectCardLengths)
	}

	expectedFirstCardCode := "AS"
	expectedSecondCardCode := "2S"
	expectedThirthCardCode := "3S"
	expectedFourthCardCode := "4S"
	expectedFifthCardCode := "5S"
	if d.Cards[0].Code == expectedFirstCardCode &&
		d.Cards[1].Code == expectedSecondCardCode &&
		d.Cards[2].Code == expectedThirthCardCode &&
		d.Cards[3].Code == expectedFourthCardCode &&
		d.Cards[4].Code == expectedFifthCardCode {
		t.Error("expected deck to be shuffled")
	}
}

func TestNewDeckWantedCardOption(t *testing.T) {
	options := deck.Options{
		WantedCards: []string{"AS", "KD", "AC", "2C", "KH"},
	}
	d, _ := deck.New(options)

	expectedRemaining := 5
	if d.Remaining() != expectedRemaining {
		t.Errorf("d.Remaining = %v, want %v", d.Remaining(), expectedRemaining)
	}

	options = deck.Options{
		WantedCards: []string{"2Z"},
	}
	d, err := deck.New(options)
	if err == nil {
		t.Errorf("expect error not to be nil")
	}
}
