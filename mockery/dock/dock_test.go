package dock_test

import (
	"testify/mockery/dock"
	"testify/mockery/dock/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	mockDock := new(mocks.Dock)

	// Call Greeting function and define return value
	mockDock.On("Greeting").Return("Test")

	// Call function
	val := mockDock.Greeting()
	assert.Equal(t, "Test", val)

	// Real Unit test
	greeting := dock.Greeting()
	assert.Equal(t, "Test", greeting)

	mockDock.AssertExpectations(t)
}
