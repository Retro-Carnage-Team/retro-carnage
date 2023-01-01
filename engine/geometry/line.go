package geometry

type Line struct {
	Start *Point
	End   *Point
}

func (l *Line) GetIntersection(o *Line) *Point {
	var a = ((o.End.X-o.Start.X)*(l.Start.Y-o.Start.Y) - (o.End.Y-o.Start.Y)*(l.Start.X-o.Start.X)) /
		((o.End.Y-o.Start.Y)*(l.End.X-l.Start.X) - (o.End.X-o.Start.X)*(l.End.Y-l.Start.Y))
	var b = ((l.End.X-l.Start.X)*(l.Start.Y-o.Start.Y) - (l.End.Y-l.Start.Y)*(l.Start.X-o.Start.X)) /
		((o.End.Y-o.Start.Y)*(l.End.X-l.Start.X) - (o.End.X-o.Start.X)*(l.End.Y-l.Start.Y))
	if a >= 0 && a <= 1 && b >= 0 && b <= 1 {
		var intersectionX = l.Start.X + a*(l.End.X-l.Start.X)
		var intersectionY = l.Start.Y + a*(l.End.Y-l.Start.Y)
		return &Point{X: intersectionX, Y: intersectionY}
	}
	return nil
}

func (l *Line) Equals(o *Line) bool {
	if nil == o {
		return false
	}

	var thisStart = l.Start.String()
	var otherStart = o.Start.String()
	var thisEnd = l.End.String()
	var otherEnd = o.End.String()

	return (thisStart == otherStart && thisEnd == otherEnd) || (thisEnd == otherStart && thisStart == otherEnd)
}
