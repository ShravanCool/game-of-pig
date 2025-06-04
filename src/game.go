package src

import (
	"fmt"
	"strconv"
	"strings"
)

var winningScore = 100
var numberOfGames = 10

func printResults(holdScore1, holdScore2, wins int) {
	winPercentage := float64(wins) / float64(numberOfGames) * 100
	lossPercentage := (float64(100) - winPercentage)

	fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n",
		holdScore1,
		holdScore2,
		wins,
		numberOfGames,
		winPercentage,
		numberOfGames-wins,
		numberOfGames,
		lossPercentage,
	)
}

func printResultsAvg(holdScore, wins, gamesPlayed int) {
	winPercentage := float64(wins) / float64(gamesPlayed) * 100
	lossPercentage := (float64(100) - winPercentage)

	fmt.Printf("Result: Wins, losses staying at k=%d: %d/%d (%0.1f%%), %d/%d (%0.1f%%)\n",
		holdScore,
		wins,
		gamesPlayed,
		winPercentage,
		gamesPlayed-wins,
		gamesPlayed,
		lossPercentage,
	)

}

func extractStrategy(strategyRange string) (int, int, error) {
	rangeArray := strings.Split(strategyRange, "-")
	if len(rangeArray) > 2 {
		return 0, 0, fmt.Errorf("Expected the strategy in the format x-y eg. 1-100\n")
	}

	start, err := strconv.Atoi(rangeArray[0])
	if err != nil || start < 1 || start > 100 {
		return 0, 0, fmt.Errorf("Expected the strategy in the format x-y eg. 1-100\n")
	}

	end, err := strconv.Atoi(rangeArray[1])
	if err != nil || end < 1 || end > 100 {
		return 0, 0, fmt.Errorf("Expected the strategy in the format x-y eg. 1-100\n")
	}

	if start >= end {
		return 0, 0, fmt.Errorf("Invalid range provided")
	}

	return start, end, nil
}

func playStrategy(holdScore1, holdScore2 int) int {
	player1 := Player{Name: "player1", HoldScore: holdScore1}
	player2 := Player{Name: "player2", HoldScore: holdScore2}

	wins := make(map[string]int)

	for i := 0; i < numberOfGames; i++ {
		player1.Score = 0
		player2.Score = 0

		p := &player1
		for {
			p.Score += p.playTurn(winningScore)

			if p.Score >= winningScore {
				break
			}
			if p.Name == "player1" {
				p = &player2
			} else {
				p = &player1
			}
		}

		wins[p.Name]++
	}
	return wins[player1.Name]
}

func PlaySingleStrategyAgainstSingleStrategy(holdScore1, holdScore2 int) {
	wins := playStrategy(holdScore1, holdScore2)
	printResults(holdScore1, holdScore2, wins)
}

func PlaySingleStrategyAgainstMultipleStrategy(holdScore1 int, holdScore2 string) {
	p2start, p2end, err := extractStrategy(holdScore2)
	if err != nil {
		fmt.Println(err)
		return
	}

	for p2HoldScore := p2start; p2HoldScore <= p2end; p2HoldScore++ {
		if holdScore1 == p2HoldScore {
			continue
		}

		wins := playStrategy(holdScore1, p2HoldScore)
		printResults(holdScore1, p2HoldScore, wins)
	}

}

func PlayMultipleStrategyAgainstMultipleStrategy(holdScore1, holdScore2 string) {
	p1start, p1end, err := extractStrategy(holdScore1)
	if err != nil {
		fmt.Println(err)
		return
	}

	p2start, p2end, err := extractStrategy(holdScore2)
	if err != nil {
		fmt.Println(err)
		return
	}

	for p1Score := p1start; p1Score <= p1end; p1Score++ {
		totalWins := 0
		totalGamesPlayed := 0

		for p2Score := p2start; p2Score <= p2end; p2Score++ {
			if p1Score == p2Score {
				continue
			}

			wins := playStrategy(p1Score, p2Score)
			totalWins += wins
			totalGamesPlayed += numberOfGames

		}
		printResultsAvg(p1Score, totalWins, totalGamesPlayed)
	}
}
