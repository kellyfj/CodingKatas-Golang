package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, reverse("hello"), "olleh")
	assert.Equal(t, reverse("Go is awesome!"), "!emosewa si oG")
	assert.Equal(t, reverse("racecar"), "racecar")
	assert.Equal(t, reverse("a"), "a")
	assert.Equal(t, reverse(""), "")
}
