package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/GGboya/ggvim/pkg/game"
	"github.com/GGboya/ggvim/pkg/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

type Model struct {
	Player     *game.Avatar
	Ghost      *game.Ghost
	FirstInput bool
	Difficulty util.Difficulty
	lastInput  string
}

type tickMsg time.Time

func tick(d util.Difficulty) tea.Cmd {
	return tea.Tick(util.DifficultyMap[d], func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tick(m.Difficulty)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// 用户敲击了键盘，表示游戏开始，此时敌方才移动
		if !m.FirstInput {
			m.FirstInput = true
			return m, tick(m.Difficulty)
		}
		key := msg.String()
		// 特判 gg
		if m.lastInput == "g" && key == "g" {
			m.Player.GoToFirstNonBlank() // gg 跳转到第一行非空字符那里
			m.lastInput = ""
			return m, nil
		}

		if key == "g" {
			m.lastInput = "g"
			return m, nil
		} else {
			m.lastInput = ""
		}

		switch key {
		case "k":
			m.Player.MoveUp()

		case "j":
			m.Player.MoveDown()

		case "h":
			m.Player.MoveLeft()

		case "l":
			m.Player.MoveRight()

		case "0":
			m.Player.ParseToBeginning()

		case "^":
			m.Player.ParseToBeginningFor6()

		case "$":
			m.Player.ParseToEnd()

		case "e":
			m.Player.ParseWordEnd()

		case "E":
			m.Player.ParseWordEndForE()

		case "b":
			m.Player.ParseWordBackward()

		case "B":
			m.Player.ParseWordBackwardForB()

		case "w":
			m.Player.ParseWordForward()

		case "W":
			m.Player.ParseWordForwardForW()

		case "G":
			m.Player.GoToLastLineFirstNonBlank()

		case "ctrl+c":
			return m, tea.Quit
		}

	case tickMsg:
		//m.Ghost.Think()
		if m.FirstInput {
			m.Ghost.ThinkMore()
			return m, tick(m.Difficulty)
		}
	}

	return m, nil
}

func (m Model) View() string {
	var builder strings.Builder
	for _, line := range game.GlobMaze.Graph {
		for _, cell := range line {
			style := termenv.String(string(cell.Char))
			switch cell.Color {
			case util.WallColor:
				builder.WriteString(style.Bold().String())
			case util.WaterColor:
				builder.WriteString(style.Foreground(util.GetColor(util.Blue)).String())
			case util.Faint:
				builder.WriteString(style.Faint().String())
			case util.PlayerColor:
				builder.WriteString(style.Background(util.GetColor(util.BrightGreen)).String())
			default:
				builder.WriteString(style.Foreground(util.GetColor(cell.Color)).String())
			}
		}
		builder.WriteString("\n")
	}
	builder.WriteString(fmt.Sprintf("\nPoints: %d/%d", m.Player.Points, game.TotalPoints))

	switch game.WonGame {
	case util.Lost:
		return "You Lost\nPress Ctrl + C to quit game"

	case util.Win:
		return "You Win\nPress Ctrl + C to quit game"
	}

	return builder.String()
}
