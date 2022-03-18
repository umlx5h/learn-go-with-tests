package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (c *CLI) PlayPoker() {
	scanner := bufio.NewScanner(c.in)
	scanner.Scan()

	c.playerStore.RecordWin(extractWinner(scanner.Text()))
}

func extractWinner(userInput string) string {
	return strings.TrimSuffix(userInput, " wins")
}
