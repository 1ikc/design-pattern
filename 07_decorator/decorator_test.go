package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestColorSquare_Draw(t *testing.T) {
	sq := Square{}
	csq := NewColorSquare(sq, "blue")
	got := csq.Draw()

	want := "blueSquare"
	assert.Equal(t, want, got)
}