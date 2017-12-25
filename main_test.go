package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceCN(t *testing.T) {
	templateJoke := "Chuck Norris doesn't read books. He stares them down until he gets the information he wants."
	newJoke := ReplaceCN(templateJoke, "Jimmy The King")
	expectedJoke := "Jimmy The King doesn't read books. He stares them down until he gets the information he wants."
	assert.Equal(t, newJoke, expectedJoke, "Did not replace properly")

}
