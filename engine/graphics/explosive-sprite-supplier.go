package graphics

// ExplosiveSpriteSupplier is an interface common to all explosives.
type ExplosiveSpriteSupplier interface {
	Sprite(int64) *SpriteWithOffset
}
