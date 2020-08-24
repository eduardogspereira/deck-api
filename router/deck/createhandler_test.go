package deck_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	router "github.com/eduardogspereira/deck-api/router/deck"
)

func responseToJSONMap(r string) map[string]interface{} {
	respJSON := make(map[string]interface{})
	json.Unmarshal([]byte(r), &respJSON)
	return respJSON
}

func TestCreateDeckHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/deck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	deckRepo := MockedDeckRepo{id: "ABC-DEF"}
	handler := http.HandlerFunc(router.CreateBuilder(deckRepo))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	respJSON := responseToJSONMap(rr.Body.String())
	expectedShuffled := false
	if respJSON["shuffled"] != expectedShuffled {
		t.Errorf("expect shuffled to be %v, got %v",
			expectedShuffled, respJSON["shuffled"])
	}

	expectedRemaining := 52.
	if respJSON["remaining"] != expectedRemaining {
		t.Errorf("expect remaining to be %v, got %v",
			expectedRemaining, respJSON["remaining"])
	}

	expectedID := "ABC-DEF"
	if respJSON["deck_id"] != expectedID {
		t.Errorf("expect deck_id to be %v, got %v",
			expectedID, respJSON["deck_id"])
	}
}

func TestCreateDeckHandlerFail(t *testing.T) {
	req, err := http.NewRequest("POST", "/deck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	deckRepo := MockedDeckRepo{returnErrorOnSave: true}
	handler := http.HandlerFunc(router.CreateBuilder(deckRepo))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

func TestCreateShuffledDeck(t *testing.T) {
	req, err := http.NewRequest("POST", "/deck?shuffled=true", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	deckRepo := MockedDeckRepo{returnErrorOnSave: true}
	handler := http.HandlerFunc(router.CreateBuilder(deckRepo))
	handler.ServeHTTP(rr, req)
}
