package poker_test

import (
	"strings"
	"testing"

	poker "github.com/26huitailang/go_tdd/command-line"
)

func TestCLI(t *testing.T) {
	t.Run("record Chris win from user input", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		in := strings.NewReader("Chris wins\n")
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Chris")
	})
	t.Run("record Cleo win from user input", func(t *testing.T) {
		playerStore := &poker.StubPlayerStore{}
		in := strings.NewReader("Cleo wins\n")
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})
}
