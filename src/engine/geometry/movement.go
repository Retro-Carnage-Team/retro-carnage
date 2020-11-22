package geometry

import (
	"math"
	"retro-carnage/util"
)

func calculateMovementDistance(elapsedTimeInMs int64, distancePerMs float64, maxDistance *float64) float64 {
	var distance = float64(elapsedTimeInMs) * distancePerMs
	if nil != maxDistance {
		return util.Min(*maxDistance, distance)
	}
	return distance
}

type movementCallbackFunc func(direction Direction, distance float64, diagonalDistance float64) float64

func calculateMovement(elapsedTimeInMs int64, direction Direction, distancePerMs float64, maxDistance *float64, fn movementCallbackFunc) float64 {
	var distanceExact = calculateMovementDistance(elapsedTimeInMs, distancePerMs, maxDistance)
	var diagonalDistance = math.Round(math.Sqrt((distanceExact * distanceExact) / 2))
	var distance = math.Round(distanceExact)
	return fn(direction, distance, diagonalDistance)
}

func CalculateMovementX(elapsedTimeInMs int64, direction Direction, distancePerMs float64, maxDistance *float64) float64 {
	var fn movementCallbackFunc = func(direction Direction, distance float64, diagonalDistance float64) float64 {
		switch direction {
		case Up:
			return 0.0
		case UpRight:
			return diagonalDistance
		case Right:
			return distance
		case DownRight:
			return diagonalDistance
		case Down:
			return 0
		case DownLeft:
			return diagonalDistance * -1
		case Left:
			return distance * -1
		case UpLeft:
			return diagonalDistance * -1
		default:
			return 0
		}
	}
	return calculateMovement(elapsedTimeInMs, direction, distancePerMs, maxDistance, fn)
}

func CalculateMovementY(elapsedTimeInMs int64, direction Direction, distancePerMs float64, maxDistance *float64) float64 {
	var fn movementCallbackFunc = func(direction Direction, distance float64, diagonalDistance float64) float64 {
		switch direction {
		case Up:
			return distance * -1
		case UpRight:
			return diagonalDistance * -1
		case Right:
			return 0
		case DownRight:
			return diagonalDistance
		case Down:
			return distance
		case DownLeft:
			return diagonalDistance
		case Left:
			return 0
		case UpLeft:
			return diagonalDistance * -1
		default:
			return 0
		}
	}
	return calculateMovement(elapsedTimeInMs, direction, distancePerMs, maxDistance, fn)
}
