package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMobileSMS_Send(t *testing.T) {
	ms := &MobileSMS{}
	assert.Nil(t, ms.Send(13100011122, "123"))
}

func ExampleMobileSMS_Send() {
	ms := &MobileSMS{}
	_ = ms.Send(13100011122, "123")
	// Output:
	// valid from MobileSMS
	// send from MobileSMS
}

func TestTelecomSMS_Send(t *testing.T) {
	ts := &TelecomSMS{}
	assert.Nil(t, ts.Send(13100011122, "123"))
}

func ExampleTelecomSMS_Send() {
	ts := &TelecomSMS{}
	_ = ts.Send(13100011122, "123")
	// Output:
	// valid from TelecomSMS
	// send from TelecomSMS
}