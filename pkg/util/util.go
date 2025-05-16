package util

import (
	"math/rand"
	"sort"
	"time"
	"unicode"

	"github.com/muesli/termenv"
)

type Difficulty int

const (
	Easy Difficulty = iota
	Medium
	Hard
)

const (
	Faint = iota
	White
	Red
	Green
	BrightGreen
	Blue
	Yellow
	Magenta
	Cyan
	WallColor
	WaterColor  = Blue
	PlayerColor = BrightGreen
	Ghost       = Red
	Char
	Nothing

	WallPortrait   = '#'
	WaterPortrait  = '~'
	PlayerPortrait = 'P'
	GhostPortrait  = 'G'

	Win    = 1
	Lost   = -1
	Gaming = 0
)

var DifficultyMap = map[Difficulty]time.Duration{
	Easy:   time.Second,
	Medium: 500 * time.Millisecond,
	Hard:   250 * time.Millisecond,
}

func init() {
	rand.Seed(time.Now().Unix())
}

func GetColor(color int) termenv.ANSIColor {
	switch color {
	case Red:
		return termenv.ANSIRed

	case Green:
		return termenv.ANSIGreen

	case BrightGreen:
		return termenv.ANSIBrightGreen

	case Blue:
		return termenv.ANSIBlue

	case Yellow:
		return termenv.ANSIYellow

	case Cyan:
		return termenv.ANSICyan

	case Magenta:
		return termenv.ANSIMagenta
	}

	return termenv.ANSIBlack
}

func MinFloat(nums ...float64) float64 {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums[0]
}

func IsAlphanumeric(chr rune) bool {
	if unicode.IsLetter(chr) || unicode.IsDigit(chr) {
		return true
	}
	return false
}
