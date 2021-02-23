package engine

import (
	"github.com/stretchr/testify/assert"
	"retro-carnage/assets"
	"retro-carnage/engine/geometry"
	"testing"
)

var (
	posPlayerOne = geometry.Rectangle{X: 500, Y: 1200, Width: 90, Height: 20}
	posPlayerTwo = geometry.Rectangle{X: 1000, Y: 900, Width: 90, Height: 20}
	posPlayerTop = geometry.Rectangle{X: 50, Y: 200, Width: 90, Height: 200}
)

func TestShouldCalculateOffsetsForBackgroundsForDirectionUp(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())
	var backgrounds = controller.Backgrounds
	assert.Equal(t, 5, len(backgrounds))
	assert.InDelta(t, 0, backgrounds[0].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[0].Offset.Y, 0.0001)

	assert.InDelta(t, 0, backgrounds[1].Offset.X, 0.0001)
	assert.InDelta(t, -1500, backgrounds[1].Offset.Y, 0.0001)

	assert.InDelta(t, 0, backgrounds[2].Offset.X, 0.0001)
	assert.InDelta(t, -3000, backgrounds[2].Offset.Y, 0.0001)
}

func TestShouldCalculateOffsetsForBackgroundsForDirectionLeft(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())
	controller.ProgressToNextSegment()
	var backgrounds = controller.Backgrounds
	assert.Equal(t, 4, len(backgrounds))
	assert.InDelta(t, 0, backgrounds[0].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[0].Offset.Y, 0.0001)

	assert.InDelta(t, -1500, backgrounds[1].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[1].Offset.Y, 0.0001)

	assert.InDelta(t, -3000, backgrounds[2].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[2].Offset.Y, 0.0001)
}

func TestShouldCalculateOffsetsForBackgroundsForDirectionRight(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())
	controller.ProgressToNextSegment()
	controller.ProgressToNextSegment()
	var backgrounds = controller.Backgrounds
	assert.Equal(t, 3, len(backgrounds))
	assert.InDelta(t, 0, backgrounds[0].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[0].Offset.Y, 0.0001)

	assert.InDelta(t, 1500, backgrounds[1].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[1].Offset.Y, 0.0001)

	assert.InDelta(t, 3000, backgrounds[2].Offset.X, 0.0001)
	assert.InDelta(t, 0, backgrounds[2].Offset.Y, 0.0001)
}

func TestShouldCalculateHowFarThePlayerIsBehindTheScrollBarrierForDirectionUp(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())

	var result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerOne, posPlayerTwo})
	assert.InDelta(t, 100, result, 0.0001)

	result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerTwo, posPlayerOne})
	assert.InDelta(t, 100, result, 0.0001)

	result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerOne})
	assert.InDelta(t, -200, result, 0.0001)
}

func TestShouldCalculateHowFarThePlayerIsBehindTheScrollBarrierForDirectionLeft(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())
	controller.ProgressToNextSegment()

	var result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerOne, posPlayerTwo})
	assert.InDelta(t, 500, result, 0.0001)

	result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerTwo, posPlayerOne})
	assert.InDelta(t, 500, result, 0.0001)

	result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerTwo})
	assert.InDelta(t, 0, result, 0.0001)
}

func TestShouldCalculateHowFarThePlayerIsBehindTheScrollBarrierForDirectionRight(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())
	controller.ProgressToNextSegment()
	controller.ProgressToNextSegment()

	var result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerOne, posPlayerTwo})
	assert.InDelta(t, 590, result, 0.0001)

	result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerTwo, posPlayerOne})
	assert.InDelta(t, 590, result, 0.0001)

	result = controller.distanceBehindScrollBarrier([]geometry.Rectangle{posPlayerOne})
	assert.InDelta(t, 90, result, 0.0001)
}

func TestShouldScrollUpWhenDirectionIsUpAndPlayerIsAtTheTop(t *testing.T) {
	var controller = NewLevelController(buildTestSegments())

	var offset = controller.UpdatePosition(50, []geometry.Rectangle{posPlayerTop})
	assert.InDelta(t, 0, offset.X, 0.0001)
	assert.InDelta(t, -15, offset.Y, 0.0001)
	assert.InDelta(t, 0, controller.Backgrounds[0].Offset.X, 0.0001)
	assert.InDelta(t, 15, controller.Backgrounds[0].Offset.Y, 0.0001)
	assert.InDelta(t, 0, controller.Backgrounds[1].Offset.X, 0.0001)
	assert.InDelta(t, -1485, controller.Backgrounds[1].Offset.Y, 0.0001)

	offset = controller.UpdatePosition(90, []geometry.Rectangle{posPlayerTop})
	assert.InDelta(t, 0, offset.X, 0.0001)
	assert.InDelta(t, -27, offset.Y, 0.0001)
	assert.InDelta(t, 0, controller.Backgrounds[0].Offset.X, 0.0001)
	assert.InDelta(t, 42, controller.Backgrounds[0].Offset.Y, 0.0001)
	assert.InDelta(t, 0, controller.Backgrounds[1].Offset.X, 0.0001)
	assert.InDelta(t, -1458, controller.Backgrounds[1].Offset.Y, 0.0001)
}

/*
test("Should scroll left when direction is left and player is at the left", () => {
  const controller = new LevelController(SEGMENTS);
  controller.ProgressToNextSegment();
  const posPlayerOne = new Rectangle(200, 500, 90, 200);

  let offset = controller.updatePosition(50, [posPlayerOne]);
  expect(offset).toEqual({ x: -15, y: 0 });
  expect(controller.backgrounds[0].offsetX).toBe(15);
  expect(controller.backgrounds[0].offsetY).toBe(0);
  expect(controller.backgrounds[1].offsetX).toBe(-1485);
  expect(controller.backgrounds[1].offsetY).toBe(0);

  offset = controller.updatePosition(90, [posPlayerOne]);
  expect(offset).toEqual({ x: -27, y: 0 });
  expect(controller.backgrounds[0].offsetX).toBe(42);
  expect(controller.backgrounds[0].offsetY).toBe(0);
  expect(controller.backgrounds[1].offsetX).toBe(-1458);
  expect(controller.backgrounds[1].offsetY).toBe(0);
});

test("Should scroll right when direction is right and player is at the right", () => {
  const controller = new LevelController(SEGMENTS);
  controller.ProgressToNextSegment();
  controller.ProgressToNextSegment();
  const posPlayerOne = new Rectangle(1300, 500, 90, 200);

  let offset = controller.updatePosition(50, [posPlayerOne]);
  expect(offset).toEqual({ x: 15, y: 0 });
  expect(controller.backgrounds[0].offsetX).toBe(-15);
  expect(controller.backgrounds[0].offsetY).toBe(0);
  expect(controller.backgrounds[1].offsetX).toBe(1485);
  expect(controller.backgrounds[1].offsetY).toBe(0);

  offset = controller.updatePosition(90, [posPlayerOne]);
  expect(offset).toEqual({ x: 27, y: 0 });
  expect(controller.backgrounds[0].offsetX).toBe(-42);
  expect(controller.backgrounds[0].offsetY).toBe(0);
  expect(controller.backgrounds[1].offsetX).toBe(1458);
  expect(controller.backgrounds[1].offsetY).toBe(0);
});

test("Should activate enemies when game scrolled far enough", () => {
  const controller = new LevelController(SEGMENTS);
  const posPlayerOne = new Rectangle(500, 200, 90, 200);

  let offset = controller.updatePosition(2000, [posPlayerOne]);
  let enemies = controller.getActivatedEnemies();
  expect(offset).toEqual({ x: 0, y: -600 });
  expect(enemies.length).toBe(1);

  offset = controller.updatePosition(50, [posPlayerOne]);
  expect(offset).toEqual({ x: 0, y: -15 });
  enemies = controller.getActivatedEnemies();
  expect(enemies.length).toBe(0);
});

test("Should return obstacles when they scroll into the visible area", () => {
  const controller = new LevelController(SEGMENTS);
  let obstacles = controller.getObstaclesOnScreen();
  expect(obstacles.length).toBe(0);

  const posPlayerOne = new Rectangle(500, 200, 90, 200);
  let offset = controller.updatePosition(2000, [posPlayerOne]);
  expect(offset).toEqual({ x: 0, y: -600 });

  obstacles = controller.getObstaclesOnScreen();
  expect(obstacles.length).toBe(1);
  expect(obstacles[0].x).toBe(400);
  expect(obstacles[0].y).toBe(300);
  expect(obstacles[0].width).toBe(200);
  expect(obstacles[0].height).toBe(200);
});

*/

func buildTestSegments() []assets.Segment {
	var result = make([]assets.Segment, 0)
	var segment = assets.Segment{
		Backgrounds: []string{"bg-dummy-1.jpg", "bg-dummy-2.jpg", "bg-dummy-3.jpg", "bg-dummy-1.jpg", "bg-dummy-2.jpg"},
		Direction:   geometry.Up.Name,
		/*
			enemies: [
			      new Enemy(
			        550,
			        [
			          {
			            duration: 4375,
			            offsetXPerMs: 0,
			            offsetYPerMs: 0.4,
			            timeElapsed: 0,
			          },
			        ],
			        new Rectangle(300, -200, 90, 200),
			        EnemySkins.GREY_ONESIE_WITH_RIFLE,
			        Directions.Down,
			        EnemyType.Person
			      ),
			    ],
		*/
		Goal:      nil,
		Obstacles: []geometry.Rectangle{{X: 400, Y: -300, Width: 200, Height: 200}},
	}
	result = append(result, segment)

	segment = assets.Segment{
		Backgrounds: []string{"bg-dummy-2.jpg", "bg-dummy-3.jpg", "bg-dummy-1.jpg", "bg-dummy-2.jpg"},
		Direction:   geometry.Left.Name,
		// Enemies: [],
		Goal:      nil,
		Obstacles: []geometry.Rectangle{},
	}
	result = append(result, segment)

	segment = assets.Segment{
		Backgrounds: []string{"bg-dummy-2.jpg", "bg-dummy-3.jpg", "bg-dummy-1.jpg"},
		Direction:   geometry.Right.Name,
		// Enemies: [],
		Goal:      &geometry.Rectangle{X: 42, Y: 42, Width: 200, Height: 200},
		Obstacles: []geometry.Rectangle{},
	}
	result = append(result, segment)

	return result
}
