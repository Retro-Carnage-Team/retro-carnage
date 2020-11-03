package network

import (
	"bytes"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"time"
)

type httpBackend struct {
	gameId string
}

type usageResponse struct {
	Id      string   `json:"id"`
	GameId  string   `json:"gameId"`
	Start   string   `json:"start"`
	Screens []string `json:"screens"`
}

type errorRequest struct {
	Id           *string    `json:"id"`
	TimeStamp    *time.Time `json:"timeStamp"`
	Message      *string    `json:"message"`
	Source       *string    `json:"source"`
	LineNumber   *int32     `json:"lineno"`
	ColumnNumber *int32     `json:"colno"`
	StackTrace   *string    `json:"stack"`
}

const backendUrl = "https://backend.retro-carnage.net"

func (hb *httpBackend) StartGameSession() {
	response, err := http.Post(backendUrl+"/usage/start-game", "application/json", nil)
	if nil != err {
		log.Error().Str("class", "httpBackend").Str("method", "StartGameSession").Err(err).Send()
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Error().Str("class", "httpBackend").Str("method", "StartGameSession").Err(err).Send()
		}

		var usage usageResponse
		err = json.Unmarshal(bodyBytes, &usage)
		if err != nil {
			log.Error().Str("class", "httpBackend").Str("method", "StartGameSession").Err(err).Send()
		}

		hb.gameId = usage.GameId
		log.Info().
			Str("class", "httpBackend").Str("method", "StartGameSession").
			Msg("Created game session with ID '" + hb.gameId + "'")
	}
}

func (hb *httpBackend) ReportGameState(screenName string) {
	var url = backendUrl + "/usage/" + hb.gameId + "/next-screen/" + screenName
	response, err := http.Post(url, "application/json", nil)
	if nil != err {
		log.Error().Str("class", "httpBackend").Str("method", "ReportGameState").Err(err).Send()
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		log.Info().
			Str("class", "httpBackend").Str("method", "ReportGameState").
			Msg("Reported game progress to screen '" + screenName + "'")
	} else {
		log.Warn().
			Str("class", "httpBackend").Str("method", "ReportGameState").
			Msg("Failed to report game progress to screen '" + screenName + "'")
	}
}

func (hb *httpBackend) ReportError(error error) {
	var url = backendUrl + "/script-errors/"

	errorMessage := error.Error()
	var errorRequest errorRequest
	errorRequest.Message = &errorMessage

	data, err := json.Marshal(errorRequest)
	if err != nil {
		log.Error().Str("class", "httpBackend").Str("method", "ReportError").Err(err).Send()
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if nil != err {
		log.Error().Str("class", "httpBackend").Str("method", "ReportError").Err(err).Send()
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		log.Info().
			Str("class", "httpBackend").Str("method", "ReportError").
			Msg("Reported error")
	} else {
		log.Warn().
			Str("class", "httpBackend").Str("method", "ReportError").
			Msg("Failed to report error")
	}
}
