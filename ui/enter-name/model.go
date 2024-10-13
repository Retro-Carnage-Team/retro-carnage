package enter_name

type model struct {
	cursorVisible     bool   // true if cursor is visible
	playerIdx         int    // index of player
	playerName        string // player name
	playerNameDisplay string // name to be displayed - possibly including cursor char
}
