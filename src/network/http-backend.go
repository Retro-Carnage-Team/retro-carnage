package network

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"retro-carnage/logging"
	"sync"
	"time"
)

type httpBackend struct {
	gameId string
	mu     sync.Mutex
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
		logging.Warning.Printf("Failed to start game session: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			logging.Warning.Printf("Failed to start game session. Failed to read server response: %v", err)
			return
		}

		var usage usageResponse
		err = json.Unmarshal(bodyBytes, &usage)
		if err != nil {
			logging.Warning.Printf("Failed to start game session. Failed to parse server response: %v", err)
			return
		}

		hb.mu.Lock()
		hb.gameId = usage.GameId
		hb.mu.Unlock()

		logging.Info.Printf("Created game session with ID '%s'", hb.gameId)
	}
}

func (hb *httpBackend) ReportGameState(screenName string) {
	hb.mu.Lock()
	var gameId = hb.gameId
	hb.mu.Unlock()

	if "" != gameId {
		var url = backendUrl + "/usage/" + gameId + "/next-screen/" + screenName
		response, err := http.Post(url, "application/json", nil)
		if nil != err {
			logging.Warning.Printf("Failed to report game state: %v", err)
			return
		}
		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			logging.Info.Printf("Reported game progress to screen '%s'", screenName)
		} else {
			logging.Warning.Printf("Failed to report game progress to screen '%s'", screenName)
		}
	}
}

func (hb *httpBackend) ReportError(error error) {
	var url = backendUrl + "/script-errors/"

	errorMessage := error.Error()
	var errorRequest errorRequest
	errorRequest.Message = &errorMessage

	data, err := json.Marshal(errorRequest)
	if err != nil {
		logging.Error.Printf("Failed to report error. Unable to build request: %v", err)
		return
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if nil != err {
		logging.Warning.Printf("Failed to report error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusCreated {
		logging.Info.Printf("Reported error: '%v'", error)
	} else {
		logging.Warning.Printf("Failed to report error: %v", err)
	}
}
