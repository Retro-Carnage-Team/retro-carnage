package geometry

import (
	"math"
	"sort"
)

type collisionCheckForCardinalDirection struct {
	border       *Line
	firstVector  *Line
	secondVector *Line
}

type collisionCheckForDiagonalDirection struct {
	firstBorder  *Line
	secondBorder *Line
	firstVector  *Line
	secondVector *Line
	thirdVector  *Line
}

type diagonalCollisionSet struct {
	distance *Point
	length   float32
}

type byLength []diagonalCollisionSet

func (a byLength) Len() int           { return len(a) }
func (a byLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLength) Less(i, j int) bool { return a[i].length < a[j].length }

// StopMovementOnCollision checks for a possible collision of rectangles movingRect and stillRect when movingRect gets
// moved into direction by the specified distance. Returns in case of a collision this method will return the movingRect
// moved for the max distance that will not cause a collision. It will return null of there is no collision of the two
// given Rectangles.

// movingRect: the rectangle that gets moved
// stillRect: another rectangle that doesn't move
// direction: direction of movement of movingRect
// distance: distance of movement of movingRect
func StopMovementOnCollision(movingRect *Rectangle, stillRect *Rectangle, direction Direction, distance *Point) *Rectangle {
	switch direction {
	case Up:
		return stopUpMovement(movingRect, stillRect, distance)
	case UpRight:
		return stopUpRightMovement(movingRect, stillRect, distance)
	case Right:
		return stopRightMovement(movingRect, stillRect, distance)
	case DownRight:
		return stopDownRightMovement(movingRect, stillRect, distance)
	case Down:
		return stopDownMovement(movingRect, stillRect, distance)
	case DownLeft:
		return stopDownLeftMovement(movingRect, stillRect, distance)
	case Left:
		return stopLeftMovement(movingRect, stillRect, distance)
	case UpLeft:
		return stopUpLeftMovement(movingRect, stillRect, distance)
	default:
		return nil
	}
}

func stopUpMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var collision = getCollisionForMovementUp(movingRect, stillRect, distance)
	if nil != collision {
		return &Rectangle{X: movingRect.X, Y: collision.Y, Width: movingRect.Width, Height: movingRect.Height}
	}
	if stillRect.Width < movingRect.Width && nil != getCollisionForMovementDown(stillRect, movingRect, &Point{X: 0, Y: -1 * distance.Y}) {
		return &Rectangle{X: movingRect.X, Y: stillRect.Y + stillRect.Height, Width: movingRect.Width, Height: movingRect.Height}
	}
	return nil
}

func stopDownMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var collision = getCollisionForMovementDown(movingRect, stillRect, distance)
	if nil != collision {
		return &Rectangle{X: movingRect.X, Y: collision.Y - movingRect.Height, Width: movingRect.Width, Height: movingRect.Height}
	}
	if stillRect.Width < movingRect.Width && nil != getCollisionForMovementUp(stillRect, movingRect, &Point{X: 0, Y: -1 * distance.Y}) {
		return &Rectangle{X: movingRect.X, Y: stillRect.Y - movingRect.Height, Width: movingRect.Width, Height: movingRect.Height}
	}
	return nil
}

func stopLeftMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var collision = getCollisionForMovementLeft(movingRect, stillRect, distance)
	if nil != collision {
		return &Rectangle{X: collision.X, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	if stillRect.Height < movingRect.Height && nil != getCollisionForMovementRight(stillRect, movingRect, &Point{X: -1 * distance.X, Y: 0}) {
		return &Rectangle{X: stillRect.X + stillRect.Width, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	return nil
}

func stopRightMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var collision = getCollisionForMovementRight(movingRect, stillRect, distance)
	if nil != collision {
		return &Rectangle{X: collision.X - movingRect.Width, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	if stillRect.Height < movingRect.Height && nil != getCollisionForMovementLeft(stillRect, movingRect, &Point{X: -1 * distance.X, Y: 0}) {
		return &Rectangle{X: stillRect.X - movingRect.Width, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	return nil
}

func stopUpRightMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var maxUpRightMovement = getMaxUpRightMovement(movingRect, stillRect, distance)
	if nil != maxUpRightMovement {
		return movingRect.Add(maxUpRightMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxDownLeftMovement = getMaxDownLeftMovement(stillRect, movingRect, &Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxDownLeftMovement {
			return &Rectangle{X: movingRect.X + -1*maxDownLeftMovement.X, Y: movingRect.Y + -1*maxDownLeftMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func stopDownRightMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var maxDownRightMovement = getMaxDownRightMovement(movingRect, stillRect, distance)
	if nil != maxDownRightMovement {
		return movingRect.Add(maxDownRightMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxUpLeftMovement = getMaxUpLeftMovement(stillRect, movingRect, &Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxUpLeftMovement {
			return &Rectangle{X: movingRect.X + -1*maxUpLeftMovement.X, Y: movingRect.Y + -1*maxUpLeftMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func stopUpLeftMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var maxUpLeftMovement = getMaxUpLeftMovement(movingRect, stillRect, distance)
	if nil != maxUpLeftMovement {
		return movingRect.Add(maxUpLeftMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxDownRightMovement = getMaxDownRightMovement(stillRect, movingRect, &Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxDownRightMovement {
			return &Rectangle{X: movingRect.X + -1*maxDownRightMovement.X, Y: movingRect.Y + -1*maxDownRightMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func stopDownLeftMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Rectangle {
	var maxDownLeftMovement = getMaxDownLeftMovement(movingRect, stillRect, distance)
	if nil != maxDownLeftMovement {
		return movingRect.Add(maxDownLeftMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxUpRightMovement = getMaxUpRightMovement(stillRect, movingRect, &Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxUpRightMovement {
			return &Rectangle{X: movingRect.X + -1*maxUpRightMovement.X, Y: movingRect.Y + -1*maxUpRightMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func checkCollisionOnCardinalDirection(provider *collisionCheckForCardinalDirection) *Point {
	var line = provider.border
	var collision = line.GetIntersection(provider.firstVector)
	if nil == collision {
		collision = line.GetIntersection(provider.secondVector)
	}
	return collision
}

func checkCollisionOnDiagonalDirection(provider *collisionCheckForDiagonalDirection) *Point {
	var collisions = make([]diagonalCollisionSet, 0)
	var borders = []*Line{provider.firstBorder, provider.secondBorder}
	var vectors = []*Line{provider.firstVector, provider.secondVector, provider.thirdVector}

	for _, v := range vectors {
		var collision *Point = nil
		for _, b := range borders {
			if nil == collision {
				collision = b.GetIntersection(v)
			}
		}
		if nil != collision {
			var a = math.Abs(float64(v.Start.X - collision.X))
			var b = math.Abs(float64(v.Start.Y - collision.Y))
			var length = math.Sqrt(a*a + b*b)

			var setItem = diagonalCollisionSet{distance: &Point{X: collision.X - v.Start.X, Y: collision.Y - v.Start.Y}, length: float32(length)}
			collisions = append(collisions, setItem)
		}
	}

	if len(collisions) > 0 {
		sort.Sort(byLength(collisions))
		return collisions[0].distance
	}
	return nil
}

func getCollisionForMovementUp(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.BottomBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y}, End: &Point{X: movingRect.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getCollisionForMovementDown(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.TopBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getCollisionForMovementLeft(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.RightBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getCollisionForMovementRight(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.LeftBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getMaxUpRightMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.LeftBorder()
	collisionCheck.secondBorder = stillRect.BottomBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.thirdVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}

func getMaxDownRightMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.LeftBorder()
	collisionCheck.secondBorder = stillRect.TopBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.thirdVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}

func getMaxUpLeftMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.RightBorder()
	collisionCheck.secondBorder = stillRect.BottomBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.thirdVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}

func getMaxDownLeftMovement(movingRect *Rectangle, stillRect *Rectangle, distance *Point) *Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.RightBorder()
	collisionCheck.secondBorder = stillRect.TopBorder()
	collisionCheck.firstVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &Line{Start: &Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.thirdVector = &Line{Start: &Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}
