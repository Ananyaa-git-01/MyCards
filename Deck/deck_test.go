package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		//t.Errorf("Expected deck lenght of 20 , but got %v", len(d)) for error if the len is not 20
		t.Errorf("Expected deck lenght of 16 , but got %v", len(d))

	}

}
