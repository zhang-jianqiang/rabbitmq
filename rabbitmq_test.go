package rabbitmq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDial(t *testing.T) {
	_, err := Dial("amqp://guest:guest@localhost:5672/")
	assert.NoError(t, err)
}
