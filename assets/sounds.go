package assets

import (
	"math/rand"
)

type SoundEffect string
type Song string

const (
	FxNone           SoundEffect = "" // this will not play any sound
	FxAk47           SoundEffect = "AK47.mp3"
	FxAr10           SoundEffect = "AR10.mp3"
	FxBar            SoundEffect = "BAR.mp3"
	FxCash           SoundEffect = "cash-register.mp3"
	FxCheatSwitch    SoundEffect = "cheat-switch.mp3"
	FxDeathEnemy0    SoundEffect = "enemy-death-0.mp3"
	FxDeathEnemy1    SoundEffect = "enemy-death-1.mp3"
	FxDeathEnemy2    SoundEffect = "enemy-death-2.mp3"
	FxDeathEnemy3    SoundEffect = "enemy-death-3.mp3"
	FxDeathEnemy4    SoundEffect = "enemy-death-4.mp3"
	FxDeathEnemy5    SoundEffect = "enemy-death-5.mp3"
	FxDeathEnemy6    SoundEffect = "enemy-death-6.mp3"
	FxDeathEnemy7    SoundEffect = "enemy-death-7.mp3"
	FxDeathPlayer1   SoundEffect = "death-player-1.mp3"
	FxDeathPlayer2   SoundEffect = "death-player-2.mp3"
	FxError          SoundEffect = "error.mp3"
	FxFnfal          SoundEffect = "FNFAL.mp3"
	FxG36            SoundEffect = "G36.mp3"
	FxG95k           SoundEffect = "G95K.mp3"
	FxGrenade1       SoundEffect = "grenade.mp3"
	FxGrenade2       SoundEffect = "grenade2.mp3"
	FxGrenade3       SoundEffect = "grenade3.mp3"
	FxHk21           SoundEffect = "HK21.mp3"
	FxLoading        SoundEffect = "loading.mp3"
	FxM28A1          SoundEffect = "M28A1.mp3"
	FxMg4            SoundEffect = "MG4.mp3"
	FxMg42           SoundEffect = "MG42.mp3"
	FxMp5            SoundEffect = "MP5.mp3"
	FxMp7            SoundEffect = "MP7.mp3"
	FxOutOfAmmo      SoundEffect = "outofammo.mp3"
	FxPistol1        SoundEffect = "pistol.mp3"
	FxPistol2        SoundEffect = "pistol2.mp3"
	FxRocketLauncher SoundEffect = "rlauncher.mp3"
	FxTankMoving     SoundEffect = "tank-moving.mp3"
	FxUzi            SoundEffect = "UZI.mp3"

	BackgroundSong1  Song = "All-We-Ever-See-of-Stars.mp3"
	BackgroundSong2  Song = "Beatdown-City.mp3"
	BackgroundSong3  Song = "Cracked-Streets-And-Broken-Windows.mp3"
	BackgroundSong4  Song = "Dance-Harder.mp3"
	BackgroundSong5  Song = "Die-Historic.mp3"
	BackgroundSong6  Song = "Drive-Fast.mp3"
	BackgroundSong7  Song = "Gaining-Traction.mp3"
	BackgroundSong8  Song = "Heavy-Traffic.mp3"
	BackgroundSong9  Song = "Hot-Nights-In-Los-Angeles.mp3"
	BackgroundSong10 Song = "It-Cant-Be-Bargained-With.mp3"
	BackgroundSong11 Song = "Missing-You.mp3"
	BackgroundSong12 Song = "Raging-Streets.mp3"
	BackgroundSong13 Song = "The-Only-Me-is-Me.mp3"
	GameOverSong     Song = "Aether.mp3"
	GameWonSong      Song = "Explosions-In-The-Sky.mp3"
	ThemeSong        Song = "The-Contra-Chop.mp3"
)

var SoundEffects = []SoundEffect{FxAk47, FxAr10, FxBar, FxCash, FxCheatSwitch, FxDeathEnemy0, FxDeathEnemy1,
	FxDeathEnemy2, FxDeathEnemy3, FxDeathEnemy4, FxDeathEnemy5, FxDeathEnemy6, FxDeathEnemy7, FxDeathPlayer1,
	FxDeathPlayer2, FxError, FxFnfal, FxG36, FxG95k, FxGrenade1, FxGrenade2, FxGrenade3, FxHk21, FxLoading, FxM28A1,
	FxMg4, FxMg42, FxMp5, FxMp7, FxOutOfAmmo, FxPistol1, FxPistol2, FxRocketLauncher, FxTankMoving, FxUzi}

var LoopingSoundEffects = []SoundEffect{FxAk47, FxAr10, FxBar, FxFnfal, FxG36, FxG95k, FxHk21, FxMg4, FxMg42, FxMp5,
	FxMp7, FxTankMoving, FxUzi}

var EnemyDeathsSoundEffects = []SoundEffect{FxDeathEnemy0, FxDeathEnemy1, FxDeathEnemy2, FxDeathEnemy3, FxDeathEnemy4,
	FxDeathEnemy5, FxDeathEnemy6, FxDeathEnemy7}

var Music = []Song{BackgroundSong1, BackgroundSong2, BackgroundSong3, BackgroundSong4, BackgroundSong5, BackgroundSong6,
	BackgroundSong7, BackgroundSong8, BackgroundSong9, BackgroundSong10, BackgroundSong11, BackgroundSong12,
	BackgroundSong13, GameOverSong, GameWonSong, ThemeSong}

func RandomEnemyDeathSoundEffect() SoundEffect {
	return EnemyDeathsSoundEffects[rand.Intn(len(EnemyDeathsSoundEffects))]
}

func DeathFxForPlayer(playerIdx int) SoundEffect {
	if playerIdx == 0 {
		return FxDeathPlayer1
	} else {
		return FxDeathPlayer2
	}
}

func SoundEffectByFileName(fileName string) *SoundEffect {
	for _, se := range SoundEffects {
		if se == SoundEffect(fileName) {
			return &se
		}
	}
	return nil
}
