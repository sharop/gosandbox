package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Error("Se esperaban 52")
	}

	if d[0] != "Ace of Spades" {
		t.Error("Se esperaba Ace of Spades")
	}

}
