package network

import "github.com/rs/zerolog/log"

type mockBackend struct{}

func (*mockBackend) StartGameSession() {
	log.Info().Str("class", "mockBackend").Str("method", "StartGameSession").Send()
}

func (*mockBackend) ReportGameState(screenName string) {
	log.Info().Str("class", "mockBackend").Str("method", "ReportGameState").Str("screenName", screenName).Send()
}

func (*mockBackend) ReportError(error error) {
	log.Info().Str("class", "mockBackend").Str("method", "ReportError").Err(error).Send()
}
