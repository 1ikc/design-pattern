package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	e := NewErrorNotification(sender)
	err := e.Notify("test msg")

	assert.Nil(t, err)
}