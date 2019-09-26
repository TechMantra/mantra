package mantra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMailbox(t *testing.T) {
	mailbox := newMailbox("test")
	assert.NotNil(t, mailbox)
}

func TestSend(t *testing.T) {
	msg := "test_msg"
	mailbox := newMailbox("test_send")
	mailbox.send(msg)
	assert.Len(t, mailbox.messages, 1)
	assert.Equal(t, msg, mailbox.messages[0].msg)
}