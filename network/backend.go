package network

import (
	"os"
)

type Backend interface {
	StartGameSession()
	ReportGameState(screenName string)
	ReportError(error error)
}

func NewBackend() Backend {
	if "development" == os.Getenv("target") {
		var result mockBackend
		return &result
	} else {
		var result httpBackend
		return &result
	}
}
