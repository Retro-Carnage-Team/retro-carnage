package geometry

import (
	"math"
)

// CalculateMovementVector returns a Point that holds the movement specified by duration, speed (distancePerMs) and
// direction.
func CalculateMovementVector(duration int64, direction Direction, distancePerMs float64) Point {
	var distance = calculateDistance(duration, distancePerMs, nil)
	return Point{
		X: calculateDistanceX(distance, direction, 0),
		Y: calculateDistanceY(distance, direction, 0),
	}
}

// Move modified a given position so that it reflects a movement specified by the duration and direction of the movement.
// The direction can be modified by a deviation in degrees - e.g. for shotgun shots that scatter on screen.
// Speed is given as distancePerMs. You can specify a maximum distance - e.g. for bullets which have a restricted range.
// Returns the distance of movement.
func Move(position *Rectangle, duration int64, direction Direction, deviationInRadians float64, distancePerMs float64, maxDistance *float64) float64 {
	var distance = calculateDistance(duration, distancePerMs, maxDistance)
	position.X += calculateDistanceX(distance, direction, deviationInRadians)
	position.Y += calculateDistanceY(distance, direction, deviationInRadians)
	return math.Round(distance)
}

func calculateDistance(elapsedTimeInMs int64, distancePerMs float64, maxDistance *float64) float64 {
	var distance = float64(elapsedTimeInMs) * distancePerMs
	if nil != maxDistance {
		distance = math.Min(*maxDistance, distance)
	}
	return distance
}

func calculateDistanceX(distance float64, direction Direction, deviationInRadians float64) float64 {
	var angle = direction.ToAngle() + deviationInRadians
	return math.Cos(angle) * distance
}

func calculateDistanceY(distance float64, direction Direction, deviationInRadians float64) float64 {
	var angle = direction.ToAngle() + deviationInRadians
	return math.Sin(angle) * distance
}
