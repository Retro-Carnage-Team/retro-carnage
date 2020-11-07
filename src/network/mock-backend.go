package network

import (
	"retro-carnage.net/util"
)

type mockBackend struct{}

func (*mockBackend) StartGameSession() {
	util.Info.Print("Started game session")
}

func (*mockBackend) ReportGameState(screenName string) {
	util.Info.Printf("Reported game progress to: %s", screenName)
}

func (*mockBackend) ReportError(error error) {
	util.Info.Printf("Reported error: %s", error)
}
