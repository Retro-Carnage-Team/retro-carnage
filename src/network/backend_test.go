package network

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type demoError struct{}

func (demoError) Error() string {
	return "fuck"
}

func TestCanCallBackend(t *testing.T) {
	// Remove this line to use httpBackend
	err := os.Setenv("target", "development")
	if nil != err {
		assert.Fail(t, err.Error())
	}

	var backend = NewBackend()
	backend.StartGameSession()
	backend.ReportGameState("test-screen")

	var demoErr demoError
	backend.ReportError(demoErr)
}
