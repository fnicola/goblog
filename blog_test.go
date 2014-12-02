package main

import (
	"testing"
)

func TestGetagFromPath(t *testing.T) {

	tag, _ := getagFromPath("bari_cities")
	if tag != "cities" {
		t.Errorf("Expected 'cities', but got '%v'\n", tag)
	}

}
