package buy_your_weapons

import (
	"fmt"
	"retro-carnage/ui/common"
)

const (
	timeAfterLastChar = 500
	timeBetweenChars  = 120
)

type controller struct {
	millisecondsPassed   int64
	playerIdx            int
	screenChangeRequired common.ScreenChangeCallback
	text                 string
	textLength           int
}

func newController(playerIndex int) *controller {
	var result = controller{
		playerIdx: playerIndex,
	}
	return &result
}

func (c *controller) setScreenChangeCallback(callback common.ScreenChangeCallback) {
	c.screenChangeRequired = callback
}

func (c *controller) update(elapsedTimeInMs int64) {
	c.millisecondsPassed += elapsedTimeInMs
	if c.textLength < 25 {
		if c.millisecondsPassed >= timeBetweenChars {
			c.textLength++
			c.text = c.getFullText()[:c.textLength]
			c.millisecondsPassed = 0
		}
	} else if c.millisecondsPassed >= timeAfterLastChar {
		if c.playerIdx == 0 {
			c.screenChangeRequired(common.ShopP1)
		} else {
			c.screenChangeRequired(common.ShopP2)
		}
	}
}

func (c *controller) getFullText() string {
	return fmt.Sprintf("BUY YOUR WEAPONS PLAYER %d", c.playerIdx+1)
}
