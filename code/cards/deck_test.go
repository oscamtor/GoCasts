package main // Mismo paquete que .go (decks.go) a testar

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // Test + Nombre Funcion a Testar
	d := newDeck() // Create a new deck
	// Code to make sure that a deck is created with x number of cards
	if len(d) != 16 { // Write if statement to see if the deck has the right number of cards
		t.Errorf("Expected deck length of 16, but got %v", len(d)) // If it doesn't, tell the go test handler that something went wrong
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) { // Test de saveToFile() & newDeckFromFileFunctions
	os.Remove("_decktesting") // Delete any files in cwd with this name

	deck := newDeck()               // Create a dek
	deck.saveToFile("_decktesting") // Save it to file

	loadedDeck := newDeckFromFile("_decktesting") // Load from file

	if len(loadedDeck) != 16 { // Assert deck len
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting") // Delete any files in cwd with this name
}
