package config

type AudioConfiguration struct {
	PlayEffects bool
	PlayMusic   bool
}

func newDefaultAudioConfiguration() AudioConfiguration {
	return AudioConfiguration{
		PlayEffects: true,
		PlayMusic:   true,
	}
}
