package engine

import (
	"math"
	"retro-carnage/assets"
	"retro-carnage/engine/characters"
	"retro-carnage/engine/geometry"
	"retro-carnage/engine/graphics"
	"retro-carnage/logging"
)

var (
	backgroundOffsets map[string]geometry.Point
)

func init() {
	backgroundOffsets = make(map[string]geometry.Point)
	backgroundOffsets[geometry.Up.Name] = geometry.Point{X: 0, Y: -1500}
	backgroundOffsets[geometry.Left.Name] = geometry.Point{X: -1500, Y: 0}
	backgroundOffsets[geometry.Right.Name] = geometry.Point{X: 1500, Y: 0}
}

type LevelController struct {
	currentSegmentIdx           int
	distanceToScroll            float64
	distanceScrolled            float64
	enemies                     []characters.Enemy
	goal                        *geometry.Rectangle
	obstacles                   []geometry.Rectangle
	segments                    []assets.Segment
	segmentScrollLengthInPixels float64
	Backgrounds                 []graphics.SpriteWithOffset
}

func NewLevelController(mission *assets.Mission) *LevelController {
	var result = &LevelController{
		currentSegmentIdx:           0,
		distanceToScroll:            0,
		distanceScrolled:            0,
		enemies:                     make([]characters.Enemy, 0),
		goal:                        nil,
		obstacles:                   make([]geometry.Rectangle, 0),
		segments:                    mission.Segments,
		segmentScrollLengthInPixels: 0,
		Backgrounds:                 make([]graphics.SpriteWithOffset, 0),
	}
	result.loadSegment(&mission.Segments[result.currentSegmentIdx])
	return result
}

func (lc *LevelController) loadSegment(segment *assets.Segment) {
	lc.goal = segment.Goal
	lc.Backgrounds = make([]graphics.SpriteWithOffset, len(segment.Backgrounds))
	for idx, bgPath := range segment.Backgrounds {
		var sprite = assets.SpriteRepository.Get(bgPath)
		var offset = backgroundOffsets[segment.Direction]
		lc.Backgrounds[idx] = graphics.SpriteWithOffset{
			Offset: *offset.Multiply(float64(idx)),
			Source: bgPath,
			Sprite: sprite,
		}
	}

	lc.segmentScrollLengthInPixels = 1500 * float64(len(lc.Backgrounds)-1)
	lc.enemies = make([]characters.Enemy, 0)
	lc.obstacles = segment.Obstacles
	lc.distanceScrolled = 0
	lc.distanceToScroll = 0
}

func (lc *LevelController) progressToNextSegment() {
	if lc.currentSegmentIdx+1 < len(lc.segments) {
		lc.currentSegmentIdx++
		lc.loadSegment(&lc.segments[lc.currentSegmentIdx])
	}
}

/**
 * Returns a those character.Enemy instances that have been activated since the last scroll movement
 */
func (lc *LevelController) ActivatedEnemies() []characters.Enemy {
	var result = make([]characters.Enemy, 0)
	var newEnemySlice = make([]characters.Enemy, 0)
	for _, enemy := range lc.enemies {
		if lc.distanceScrolled >= enemy.ActivationDistance {
			result = append(result, enemy)
		} else {
			newEnemySlice = append(newEnemySlice, enemy)
		}
	}
	lc.enemies = newEnemySlice
	return result
}

func (lc *LevelController) UpdatePosition(elapsedTimeInMs int64, playerPositions []geometry.Rectangle) geometry.Point {
	// TODO: This currently ignores the position of the second player.
	// We should only scroll if we don't kick the other player out of the visible area

	// How far is the player behind the scroll barrier?
	var scrollDistanceByPlayerPosition = lc.distanceBehindScrollBarrier(playerPositions)

	// Has he been further behind the barrier before?
	lc.distanceToScroll = math.Max(scrollDistanceByPlayerPosition, lc.distanceToScroll)

	var numberOfPixelsToScrollLeftForThisSegment = lc.segmentScrollLengthInPixels - lc.distanceScrolled
	var availablePixelsToScroll = math.Min(lc.distanceToScroll, numberOfPixelsToScrollLeftForThisSegment)
	var scrollDistanceForTheElapsedTime = math.Floor(float64(elapsedTimeInMs) * ScrollMovementPerMs)
	availablePixelsToScroll = math.Min(availablePixelsToScroll, scrollDistanceForTheElapsedTime)

	return lc.scroll(availablePixelsToScroll)
}

func (lc *LevelController) scroll(pixels float64) geometry.Point {
	lc.distanceToScroll -= pixels
	lc.distanceScrolled += pixels

	var direction = lc.segments[lc.currentSegmentIdx].Direction
	if geometry.Up.Name == direction {
		return lc.scrollUp(pixels)
	}
	if geometry.Left.Name == direction {
		return lc.scrollLeft(pixels)
	}
	if geometry.Right.Name == direction {
		return lc.scrollRight(pixels)
	}

	// should not happen
	logging.Error.Fatalf("Level segment has unknown direction: %s", direction)
	return geometry.Point{X: 0, Y: 0}
}

func (lc *LevelController) scrollUp(pixels float64) geometry.Point {
	for _, background := range lc.Backgrounds {
		background.Offset.Y += pixels
	}
	if nil != lc.goal {
		lc.goal.Y += pixels
	}
	if 0 <= lc.Backgrounds[len(lc.Backgrounds)-1].Offset.Y {
		lc.Backgrounds[len(lc.Backgrounds)-1].Offset.Y = 0
		lc.progressToNextSegment()
	}
	return geometry.Point{X: 0, Y: -pixels}
}

func (lc *LevelController) scrollLeft(pixels float64) geometry.Point {
	for _, background := range lc.Backgrounds {
		background.Offset.X += pixels
	}
	if nil != lc.goal {
		lc.goal.X += pixels
	}
	if 0 <= lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X {
		lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X = 0
		lc.progressToNextSegment()
	}
	return geometry.Point{X: -pixels, Y: 0}
}

func (lc *LevelController) scrollRight(pixels float64) geometry.Point {
	for _, background := range lc.Backgrounds {
		background.Offset.X -= pixels
	}
	if nil != lc.goal {
		lc.goal.X -= pixels
	}
	if 0 >= lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X {
		lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X = 0
		lc.progressToNextSegment()
	}
	return geometry.Point{X: pixels, Y: 0}
}
