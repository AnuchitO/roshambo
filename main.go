package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/AnuchitPrasertsang/roshambo/display"
	"github.com/AnuchitPrasertsang/roshambo/play"
	"github.com/AnuchitPrasertsang/roshambo/score"
	"github.com/AnuchitPrasertsang/roshambo/utils"

	"github.com/CrowdSurge/banner"
)

// TODO: recgonize answer

func printDecided(arts string) {
	fmt.Println(arts)
	fmt.Println()
	fmt.Println()
}

func printResultArt(arts, result string) {
	fmt.Printf(utils.Color(result))
	banner.Print(display.Title(result))
	fmt.Println(arts)
}

func main() {
	clearScreen()
	clearScreen()
	showFirstPage()

	userScore := score.New()
	computerScore := score.New()

	p := play.New(userScore, computerScore)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userKey := scanner.Text()
		if userKey == "exit" {
			os.Exit(0)
		}
		fmt.Printf(utils.ResetColor)
		u := utils.MapUserKey(userKey)

		if u == "" {
			banner.Print("unknow choice")
			fmt.Println(display.Hint())
			continue
		}

		user, com, result := p.Play(u)
		banner.Print(display.Title(""))
		arts := display.Display(user, com)
		printDecided(arts)

		time.Sleep(time.Millisecond * 400)
		clearScreen()
		printResultArt(arts, result)
		fmt.Printf("Score => Me: %d , Computer: %d, Tie: %d\n", userScore.Win, computerScore.Win, userScore.Tie)

		fmt.Println(display.Hint())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
func clearScreen() {
	for i := 0; i < 20; i++ {
		fmt.Print("\x0c") // Clear screen
	}
}

func showFirstPage() {
	for _, line := range display.FirstPage() {
		banner.Print(line)
	}
}
