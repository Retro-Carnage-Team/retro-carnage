package network

import (
	"retro-carnage.net/logging"
)

type mockBackend struct{}

func (*mockBackend) StartGameSession() {
	logging.Info.Print("Started game session")
}

func (*mockBackend) ReportGameState(screenName string) {
	logging.Info.Printf("Reported game progress to %s", screenName)
}

func (*mockBackend) ReportError(error error) {
	logging.Info.Printf("Reported error %s", error)
}
