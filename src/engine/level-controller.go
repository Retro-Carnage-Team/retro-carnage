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
	backgroundOffsets[geometry.Up.Name] = geometry.Point{X: 0, Y: -ScreenSize}
	backgroundOffsets[geometry.Left.Name] = geometry.Point{X: -ScreenSize, Y: 0}
	backgroundOffsets[geometry.Right.Name] = geometry.Point{X: ScreenSize, Y: 0}
}

type LevelController struct {
	currentSegmentIdx           int
	distanceToScroll            float64
	distanceScrolled            float64
	enemies                     []assets.Enemy
	goal                        *geometry.Rectangle
	obstacles                   []assets.Obstacle
	segments                    []assets.Segment
	segmentScrollLengthInPixels float64
	Backgrounds                 []graphics.SpriteWithOffset
}

// NewLevelController creates a new LevelController.
func NewLevelController(segments []assets.Segment) *LevelController {
	var result = &LevelController{
		currentSegmentIdx:           0,
		distanceToScroll:            0,
		distanceScrolled:            0,
		enemies:                     make([]assets.Enemy, 0),
		goal:                        nil,
		obstacles:                   make([]assets.Obstacle, 0),
		segments:                    segments,
		segmentScrollLengthInPixels: 0,
		Backgrounds:                 make([]graphics.SpriteWithOffset, 0),
	}
	result.loadSegment(&segments[result.currentSegmentIdx])
	return result
}

func (lc *LevelController) loadSegment(segment *assets.Segment) {
	lc.goal = segment.Goal
	lc.Backgrounds = make([]graphics.SpriteWithOffset, len(segment.Backgrounds))
	for idx, bgPath := range segment.Backgrounds {
		var offset = backgroundOffsets[segment.Direction]
		var sprite = assets.SpriteRepository.Get(bgPath)
		if nil == sprite {
			logging.Warning.Printf("Failed to load level background sprite: %s", bgPath)
		}
		lc.Backgrounds[idx] = graphics.SpriteWithOffset{
			Offset: *offset.Multiply(float64(idx)),
			Source: bgPath,
			Sprite: sprite,
		}
	}

	lc.segmentScrollLengthInPixels = 1500 * float64(len(lc.Backgrounds)-1)
	lc.enemies = segment.Enemies
	lc.obstacles = segment.Obstacles
	lc.distanceScrolled = 0
	lc.distanceToScroll = 0
}

// ProgressToNextSegment proceeds to the next level segment when the Player(s) finished the current segment.
func (lc *LevelController) ProgressToNextSegment() {
	if lc.currentSegmentIdx+1 < len(lc.segments) {
		lc.currentSegmentIdx++
		lc.loadSegment(&lc.segments[lc.currentSegmentIdx])
	}
}

// ActivatedEnemies returns Enemy instances that have been activated since the last scroll movement
func (lc *LevelController) ActivatedEnemies() []characters.ActiveEnemy {
	var result = make([]characters.ActiveEnemy, 0)
	var newEnemySlice = make([]assets.Enemy, 0)
	for _, enemy := range lc.enemies {
		if lc.distanceScrolled >= enemy.Position.BottomBorder().Start.Y-ScreenSize {
			var activatedEnemy = lc.activateEnemy(&enemy)
			result = append(result, activatedEnemy)
		} else {
			newEnemySlice = append(newEnemySlice, enemy)
		}
	}
	lc.enemies = newEnemySlice
	return result
}

func (lc *LevelController) UpdatePosition(elapsedTimeInMs int64, playerPositions []*geometry.Rectangle) geometry.Point {
	// How far is the player behind the scroll barrier?
	var scrollDistanceByPlayerPosition = lc.distanceBehindScrollBarrier(playerPositions)

	// Has he been further behind the barrier before?
	lc.distanceToScroll = math.Max(scrollDistanceByPlayerPosition, lc.distanceToScroll)

	var numberOfPixelsToScrollLeftForThisSegment = lc.segmentScrollLengthInPixels - lc.distanceScrolled
	var availablePixelsToScroll = math.Min(lc.distanceToScroll, numberOfPixelsToScrollLeftForThisSegment)
	var scrollDistanceForTheElapsedTime = math.Floor(float64(elapsedTimeInMs) * ScrollMovementPerMs)
	availablePixelsToScroll = math.Min(availablePixelsToScroll, scrollDistanceForTheElapsedTime)
	availablePixelsToScroll = math.Min(availablePixelsToScroll, lc.distanceFromScreenExit(playerPositions))

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
	for idx := range lc.Backgrounds {
		lc.Backgrounds[idx].Offset.Y += pixels
	}
	if nil != lc.goal {
		lc.goal.Y += pixels
	}
	if 0 <= lc.Backgrounds[len(lc.Backgrounds)-1].Offset.Y {
		lc.Backgrounds[len(lc.Backgrounds)-1].Offset.Y = 0
		lc.ProgressToNextSegment()
	}
	return geometry.Point{X: 0, Y: -pixels}
}

func (lc *LevelController) scrollLeft(pixels float64) geometry.Point {
	for idx := range lc.Backgrounds {
		lc.Backgrounds[idx].Offset.X += pixels
	}
	if nil != lc.goal {
		lc.goal.X += pixels
	}
	if 0 <= lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X {
		lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X = 0
		lc.ProgressToNextSegment()
	}
	return geometry.Point{X: -pixels, Y: 0}
}

func (lc *LevelController) scrollRight(pixels float64) geometry.Point {
	for idx := range lc.Backgrounds {
		lc.Backgrounds[idx].Offset.X -= pixels
	}
	if nil != lc.goal {
		lc.goal.X -= pixels
	}
	if 0 >= lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X {
		lc.Backgrounds[len(lc.Backgrounds)-1].Offset.X = 0
		lc.ProgressToNextSegment()
	}
	return geometry.Point{X: pixels, Y: 0}
}

// VisibleBackgrounds returns the graphics.SpriteWithOffset for all backgrounds in the visible screen rect.
func (lc *LevelController) VisibleBackgrounds() []graphics.SpriteWithOffset {
	var result = make([]graphics.SpriteWithOffset, 0)
	var negativeScreenSize = float64(ScreenSize * -1)
	for _, background := range lc.Backgrounds {
		var x = background.Offset.X
		var y = background.Offset.Y
		if (negativeScreenSize < x) && (ScreenSize > x) && (negativeScreenSize < y) && (ScreenSize > y) {
			result = append(result, background)
		}
	}
	return result
}

func (lc *LevelController) distanceBehindScrollBarrier(playerPositions []*geometry.Rectangle) float64 {
	var direction = lc.segments[lc.currentSegmentIdx].Direction
	if geometry.Up.Name == direction {
		var topMostPosition = float64(ScreenSize)
		for _, pos := range playerPositions {
			topMostPosition = math.Min(topMostPosition, pos.Y)
		}
		return ScrollBarrierUp - topMostPosition
	}
	if geometry.Left.Name == direction {
		var leftMostPosition = float64(ScreenSize)
		for _, pos := range playerPositions {
			leftMostPosition = math.Min(leftMostPosition, pos.X)
		}
		return ScrollBarrierLeft - leftMostPosition
	}
	if geometry.Right.Name == direction {
		var rightMostPosition float64 = 0
		for _, pos := range playerPositions {
			rightMostPosition = math.Max(rightMostPosition, pos.X+pos.Width)
		}
		return rightMostPosition - ScrollBarrierRight
	}

	// should not happen
	logging.Error.Fatalf("Level segment has unknown direction: %s", direction)
	return 0
}

// distanceFromScreenExit returns the minimum distance of a player from the screen side opposite the scrolling
// direction. This is used to limit the scrolling so that no player gets pushed off screen.
func (lc *LevelController) distanceFromScreenExit(playerPositions []*geometry.Rectangle) float64 {
	var direction = lc.segments[lc.currentSegmentIdx].Direction
	if geometry.Up.Name == direction {
		var result = float64(ScreenSize)
		for _, pos := range playerPositions {
			result = math.Min(ScreenSize-MinPlayerDistanceToBorder-PlayerHitRectHeight-pos.Y, result)
		}
		return result
	}
	if geometry.Left.Name == direction {
		var result = float64(ScreenSize)
		for _, pos := range playerPositions {
			result = math.Min(ScreenSize-MinPlayerDistanceToBorder-PlayerHitRectWidth-pos.X, result)
		}
		return result
	}
	if geometry.Right.Name == direction {
		var result = float64(ScreenSize)
		for _, pos := range playerPositions {
			result = math.Min(pos.X-MinPlayerDistanceToBorder, result)
		}
		return result
	}

	// should not happen
	logging.Error.Fatalf("Level segment has unknown direction: %s", direction)
	return 0
}

// GoalReached returns true when a Player reached the goal rect of the current mission.
func (lc *LevelController) GoalReached(playerPositions []*geometry.Rectangle) bool {
	if nil != lc.goal {
		for _, playerPosition := range playerPositions {
			if nil != playerPosition.Intersection(lc.goal) {
				return true
			}
		}
	}
	return false
}

// ObstaclesOnScreen returns all assets.Obstacle that within the visible screen rect.
func (lc *LevelController) ObstaclesOnScreen() []assets.Obstacle {
	var direction = lc.segments[lc.currentSegmentIdx].Direction
	var scrollAdjustment = geometry.Point{X: 0, Y: 0}
	switch direction {
	case geometry.Up.Name:
		scrollAdjustment = geometry.Point{X: 0, Y: lc.distanceScrolled}
	case geometry.Left.Name:
		scrollAdjustment = geometry.Point{X: lc.distanceScrolled, Y: 0}
	case geometry.Right.Name:
		scrollAdjustment = geometry.Point{X: -1 * lc.distanceScrolled, Y: 0}
	}

	var result = make([]assets.Obstacle, 0)
	var screenRect = screenRect()
	for _, obstacle := range lc.obstacles {
		obstacle.Rectangle.Add(&scrollAdjustment)
		if nil != obstacle.Rectangle.Intersection(screenRect) {
			result = append(result, obstacle)
		}
	}
	return result
}

func (lc *LevelController) activateEnemy(e *assets.Enemy) characters.ActiveEnemy {
	var direction = geometry.GetDirectionByName(e.Direction)
	var result = characters.ActiveEnemy{
		Actions:                 e.Actions,
		Dying:                   false,
		DyingAnimationCountDown: 0,
		Movements:               lc.convertEnemyMovements(e.Movements),
		Skin:                    characters.EnemySkin(e.Skin),
		Type:                    characters.EnemyType(e.Type),
		ViewingDirection:        direction,
	}
	result.SetPosition(e.Position.Clone())

	if int(characters.Person) == e.Type {
		if nil == direction {
			logging.Error.Fatalf("no such direction: %s", e.Direction)
		}
		if !characters.IsEnemySkin(e.Skin) {
			logging.Error.Fatalf("no such enemy skin: %s", e.Skin)
		}
		result.SpriteSupplier = characters.NewEnemyPersonSpriteSupplier(*result.ViewingDirection)
	}

	if int(characters.Landmine) == e.Type {
		result.SpriteSupplier = &characters.EnemyLandmineSpriteSupplier{}
	}

	return result
}

func (lc *LevelController) convertEnemyMovements(movements []assets.EnemyMovement) []*characters.EnemyMovement {
	var result = make([]*characters.EnemyMovement, 0)
	for _, movement := range movements {
		var converted = characters.NewEnemyMovement(&movement)
		result = append(result, &converted)
	}
	return result
}
