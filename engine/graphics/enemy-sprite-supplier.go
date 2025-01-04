package graphics

type EnemySpriteSupplier interface {
	GetDurationOfEnemyDeathAnimation() int64
	Sprite(elapsedTimeInMs int64) *SpriteWithOffset
}
