package geometry

// Positioned includes basically anything that has a position. BOOM!
type Positioned interface {
	Position() *Rectangle
}
