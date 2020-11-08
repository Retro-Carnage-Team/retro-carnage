package geometry

import (
	"math"
	"retro-carnage.net/util"
)

func calculateMovementDistance(elapsedTimeInMs int32, distancePerMs float32, maxDistance *float32) float32 {
	var distance = float32(elapsedTimeInMs) * distancePerMs
	if nil != maxDistance {
		return util.MathUtil{}.Min(*maxDistance, distance)
	}
	return distance
}

type movementCallbackFunc func(direction Direction, distance float32, diagonalDistance float32) float32

func calculateMovement(elapsedTimeInMs int32, direction Direction, distancePerMs float32, maxDistance *float32, fn movementCallbackFunc) float32 {
	var distanceExact = calculateMovementDistance(elapsedTimeInMs, distancePerMs, maxDistance)
	var diagonalDistance = math.Round(math.Sqrt(float64((distanceExact * distanceExact) / 2)))
	var distance = math.Round(float64(distanceExact))
	return fn(direction, float32(distance), float32(diagonalDistance))
}

func CalculateMovementX(elapsedTimeInMs int32, direction Direction, distancePerMs float32, maxDistance *float32) float32 {
	var fn movementCallbackFunc = func(direction Direction, distance float32, diagonalDistance float32) float32 {
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

func CalculateMovementY(elapsedTimeInMs int32, direction Direction, distancePerMs float32, maxDistance *float32) float32 {
	var fn movementCallbackFunc = func(direction Direction, distance float32, diagonalDistance float32) float32 {
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
