package engine

import (
	"math"
	"retro-carnage.net/engine/geometry"
	"sort"
)

type collisionCheckForCardinalDirection struct {
	border       *geometry.Line
	firstVector  *geometry.Line
	secondVector *geometry.Line
}

type collisionCheckForDiagonalDirection struct {
	firstBorder  *geometry.Line
	secondBorder *geometry.Line
	firstVector  *geometry.Line
	secondVector *geometry.Line
	thirdVector  *geometry.Line
}

type diagonalCollisionSet struct {
	distance *geometry.Point
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
func StopMovementOnCollision(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, direction Direction, distance *geometry.Point) *geometry.Rectangle {
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

func stopUpMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var collision = getCollisionForMovementUp(movingRect, stillRect, distance)
	if nil != collision {
		return &geometry.Rectangle{X: movingRect.X, Y: collision.Y, Width: movingRect.Width, Height: movingRect.Height}
	}
	if stillRect.Width < movingRect.Width && nil != getCollisionForMovementDown(stillRect, movingRect, &geometry.Point{X: 0, Y: -1 * distance.Y}) {
		return &geometry.Rectangle{X: movingRect.X, Y: stillRect.Y + stillRect.Height, Width: movingRect.Width, Height: movingRect.Height}
	}
	return nil
}

func stopDownMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var collision = getCollisionForMovementDown(movingRect, stillRect, distance)
	if nil != collision {
		return &geometry.Rectangle{X: movingRect.X, Y: collision.Y - movingRect.Height, Width: movingRect.Width, Height: movingRect.Height}
	}
	if stillRect.Width < movingRect.Width && nil != getCollisionForMovementUp(stillRect, movingRect, &geometry.Point{X: 0, Y: -1 * distance.Y}) {
		return &geometry.Rectangle{X: movingRect.X, Y: stillRect.Y - movingRect.Height, Width: movingRect.Width, Height: movingRect.Height}
	}
	return nil
}

func stopLeftMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var collision = getCollisionForMovementLeft(movingRect, stillRect, distance)
	if nil != collision {
		return &geometry.Rectangle{X: collision.X, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	if stillRect.Height < movingRect.Height && nil != getCollisionForMovementRight(stillRect, movingRect, &geometry.Point{X: -1 * distance.X, Y: 0}) {
		return &geometry.Rectangle{X: stillRect.X + stillRect.Width, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	return nil
}

func stopRightMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var collision = getCollisionForMovementRight(movingRect, stillRect, distance)
	if nil != collision {
		return &geometry.Rectangle{X: collision.X - movingRect.Width, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	if stillRect.Height < movingRect.Height && nil != getCollisionForMovementLeft(stillRect, movingRect, &geometry.Point{X: -1 * distance.X, Y: 0}) {
		return &geometry.Rectangle{X: stillRect.X - movingRect.Width, Y: movingRect.Y, Width: movingRect.Width, Height: movingRect.Height}
	}

	return nil
}

func stopUpRightMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var maxUpRightMovement = getMaxUpRightMovement(movingRect, stillRect, distance)
	if nil != maxUpRightMovement {
		return movingRect.Add(maxUpRightMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxDownLeftMovement = getMaxDownLeftMovement(stillRect, movingRect, &geometry.Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxDownLeftMovement {
			return &geometry.Rectangle{X: movingRect.X + -1*maxDownLeftMovement.X, Y: movingRect.Y + -1*maxDownLeftMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func stopDownRightMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var maxDownRightMovement = getMaxDownRightMovement(movingRect, stillRect, distance)
	if nil != maxDownRightMovement {
		return movingRect.Add(maxDownRightMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxUpLeftMovement = getMaxUpLeftMovement(stillRect, movingRect, &geometry.Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxUpLeftMovement {
			return &geometry.Rectangle{X: movingRect.X + -1*maxUpLeftMovement.X, Y: movingRect.Y + -1*maxUpLeftMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func stopUpLeftMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var maxUpLeftMovement = getMaxUpLeftMovement(movingRect, stillRect, distance)
	if nil != maxUpLeftMovement {
		return movingRect.Add(maxUpLeftMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxDownRightMovement = getMaxDownRightMovement(stillRect, movingRect, &geometry.Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxDownRightMovement {
			return &geometry.Rectangle{X: movingRect.X + -1*maxDownRightMovement.X, Y: movingRect.Y + -1*maxDownRightMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func stopDownLeftMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Rectangle {
	var maxDownLeftMovement = getMaxDownLeftMovement(movingRect, stillRect, distance)
	if nil != maxDownLeftMovement {
		return movingRect.Add(maxDownLeftMovement)
	}

	if stillRect.Width < movingRect.Width || stillRect.Height < movingRect.Height {
		var maxUpRightMovement = getMaxUpRightMovement(stillRect, movingRect, &geometry.Point{X: -1 * distance.X, Y: -1 * distance.Y})
		if nil != maxUpRightMovement {
			return &geometry.Rectangle{X: movingRect.X + -1*maxUpRightMovement.X, Y: movingRect.Y + -1*maxUpRightMovement.Y, Width: movingRect.Width, Height: movingRect.Height}
		}
	}

	return nil
}

func checkCollisionOnCardinalDirection(provider *collisionCheckForCardinalDirection) *geometry.Point {
	var line = provider.border
	var collision = line.GetIntersection(provider.firstVector)
	if nil == collision {
		collision = line.GetIntersection(provider.secondVector)
	}
	return collision
}

func checkCollisionOnDiagonalDirection(provider *collisionCheckForDiagonalDirection) *geometry.Point {
	var collisions = make([]diagonalCollisionSet, 0)
	var borders = []*geometry.Line{provider.firstBorder, provider.secondBorder}
	var vectors = []*geometry.Line{provider.firstVector, provider.secondVector, provider.thirdVector}

	for _, v := range vectors {
		var collision *geometry.Point = nil
		for _, b := range borders {
			if nil == collision {
				collision = b.GetIntersection(v)
			}
		}
		if nil != collision {
			var a = math.Abs(float64(v.Start.X - collision.X))
			var b = math.Abs(float64(v.Start.Y - collision.Y))
			var length = math.Sqrt(a*a + b*b)

			var setItem = diagonalCollisionSet{distance: &geometry.Point{X: collision.X - v.Start.X, Y: collision.Y - v.Start.Y}, length: float32(length)}
			collisions = append(collisions, setItem)
		}
	}

	if len(collisions) > 0 {
		sort.Sort(byLength(collisions))
		return collisions[0].distance
	}
	return nil
}

func getCollisionForMovementUp(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.GetBottomBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getCollisionForMovementDown(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.GetTopBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getCollisionForMovementLeft(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.GetRightBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getCollisionForMovementRight(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForCardinalDirection
	collisionCheck.border = stillRect.GetLeftBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height}}
	return checkCollisionOnCardinalDirection(&collisionCheck)
}

func getMaxUpRightMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.GetLeftBorder()
	collisionCheck.secondBorder = stillRect.GetBottomBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.thirdVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}

func getMaxDownRightMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.GetLeftBorder()
	collisionCheck.secondBorder = stillRect.GetTopBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.thirdVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}

func getMaxUpLeftMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.GetRightBorder()
	collisionCheck.secondBorder = stillRect.GetBottomBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.thirdVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}

func getMaxDownLeftMovement(movingRect *geometry.Rectangle, stillRect *geometry.Rectangle, distance *geometry.Point) *geometry.Point {
	var collisionCheck collisionCheckForDiagonalDirection
	collisionCheck.firstBorder = stillRect.GetRightBorder()
	collisionCheck.secondBorder = stillRect.GetTopBorder()
	collisionCheck.firstVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + distance.Y}}
	collisionCheck.secondVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	collisionCheck.thirdVector = &geometry.Line{Start: &geometry.Point{X: movingRect.X + movingRect.Width, Y: movingRect.Y + movingRect.Height}, End: &geometry.Point{X: movingRect.X + movingRect.Width + distance.X, Y: movingRect.Y + movingRect.Height + distance.Y}}
	return checkCollisionOnDiagonalDirection(&collisionCheck)
}
