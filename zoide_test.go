package main

import (
	"testing"
	//"fmt"
)

func TestIncPlayer(t *testing.T) {
	g := Game{
		numPlayers:  3,
		whichPlayer: 1,
	}
	for i := 0; i < 20; i++ {
		g.incPlayer()
	}

	if g.whichPlayer != 3 {
		t.Error("Wrong player turn. Need: 3, was: ", g.whichPlayer)
	}
}
