package main

import "testing"


func TestNewCards(t *testing.T) {
	d := newCards()
	if len(d) != 16 {
		t.Errorf("cards len should be 16 but it is %v", len(d))
	}
}
