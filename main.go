package main

import (
	"fmt"
	"os"

	"github.com/GGboya/ggvim/pkg/game"
	"github.com/GGboya/ggvim/pkg/model"
	"github.com/GGboya/ggvim/pkg/util"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// 让玩家来控制游戏难度
	fmt.Println("请选择游戏难度")
	var d util.Difficulty
	fmt.Println("输入 0 easy, 1 mid, 2 hard")
	fmt.Scan(&d)
	if d < 0 || d > 2 {
		fmt.Println("违法输入，请输入 0, 1, 2")
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel(d))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel(d util.Difficulty) model.Model {
	game.InitGame()
	return model.Model{
		Player:     game.Player,
		Ghost:      game.Ghost1,
		FirstInput: false,
		Difficulty: d,
	}
}
