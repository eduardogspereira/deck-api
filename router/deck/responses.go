package deck

// CreateResponse implements the object response interface
// for the POST /deck.
type CreateResponse struct {
	DeckID    string `json:"deck_id"`
	Shuffled  bool   `json:"shuffled"`
	Remaining int    `json:"remaining"`
}
