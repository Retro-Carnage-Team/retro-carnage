package common

type ScreenName string

const (
	Loading                 ScreenName = "Loading screen"
	Start                   ScreenName = "Start screen"
	Title                   ScreenName = "Title screen"
	ConfigurationSelect     ScreenName = "Configuration select screen"
	ConfigurationResult     ScreenName = "Configuration result screen"
	ConfigurationOptions    ScreenName = "Configuration options screen"
	ConfigurationAudio      ScreenName = "Configuration audio settings screen"
	ConfigurationVideo      ScreenName = "Configuration video settings screen"
	ConfigurationControls   ScreenName = "Configuration controller selection screen"
	ConfigurationControlsP1 ScreenName = "Configuration controller settings screen (Player 1)"
	ConfigurationControlsP2 ScreenName = "Configuration controller settings screen (Player 2)"
	Mission                 ScreenName = "Mission screen"
	BuyYourWeaponsP1        ScreenName = "Buy your items screen (Player 1)"
	BuyYourWeaponsP2        ScreenName = "Buy your items screen (Player 2)"
	ShopP1                  ScreenName = "Shop screen (Player 1)"
	ShopP2                  ScreenName = "Shop screen (Player 2)"
	LetTheMissionBegin      ScreenName = "Let the mission begin screen"
	Game                    ScreenName = "Game screen"
	EnterNameP1             ScreenName = "Enter name screen (Player 1)"
	EnterNameP2             ScreenName = "Enter name screen (Player 2)"
	HighScore               ScreenName = "High score table screen"
)
