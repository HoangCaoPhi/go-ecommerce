package basic

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAddOne(t *testing.T) {
	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
}
