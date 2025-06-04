package main

import (
	"fmt"
	"game-of-pig/src"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
)

var command = &cobra.Command{
	Use: "pig",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Help()
			return
		}

		p1, p2 := args[0], args[1]
		if strings.Contains(p1, "-") && strings.Contains(p2, "-") {
			src.PlayMultipleStrategyAgainstMultipleStrategy(p1, p2)
			return
		}

		holdScore1, err := strconv.Atoi(p1)
		if err != nil || holdScore1 < 1 || holdScore1 > 100 {
			fmt.Printf("Expected the strategy in the format x-y eg. 1-100")
		}

		if strings.Contains(p2, "-") {
			src.PlaySingleStrategyAgainstMultipleStrategy(holdScore1, p2)
			return
		}

		holdScore2, err := strconv.Atoi(p1)
		if err != nil || holdScore2 < 1 || holdScore2 > 100 {
			fmt.Printf("Expected the strategy in the format x-y eg. 1-100")
		}

		src.PlaySingleStrategyAgainstSingleStrategy(holdScore1, holdScore2)
	},
}

func execute() {
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	command.SetUsageFunc(nil)
	command.SetUsageTemplate(
		`pig - A command line tool to simulate a game of pig. It is a two player game played with a six-sided die.

		Usage:
			pig [strategy] [strategy]
	
		Args:
			strategy	The number between 1 to 100

		Description:
			This CLI tool accepts two numbers between 1 and 100 as arguments. These are strategies for the two players, and performs a simulation using them, and show the result.

		Example usage:
			$ pig 10 15
			Result: Holding at 10 vs Holding at 15: wins 3/10 (30.0%), losses: 7/10 (70.0%)`,
	)
}

func main() {
	execute()
}
