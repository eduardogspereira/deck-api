package deck

type createResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}

type cardResponse struct {
	Value string `json:"value"`
	Suit  string `json:"suit"`
	Code  string `json:"code"`
}

type loadResponse struct {
	DeckID    string         `json:"deck_id"`
	Shuffled  bool           `json:"shuffled"`
	Remaining int            `json:"remaining"`
	Cards     []cardResponse `json:"cards"`
}

type drawCardResponse struct {
	Cards []cardResponse `json:"cards"`
}
