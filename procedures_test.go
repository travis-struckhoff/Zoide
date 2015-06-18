package main

import (
	//"fmt"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	"testing"
)

func TestTestWinnerDL(t *testing.T) {
	b := board.MakeBoard(2)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(x, x, board.Black)
	}
	win := TestWinner(b, 2, true)
	if win[0] != board.Black {
		t.Error("Didn't see diag Black")
	}

	b = board.MakeBoard(2)
	for x := 0; x < 5; x++ {
		if x == 3 {
			continue
		}
		b.PlacePiece(x, x, board.Black)
	}

	b.PlacePiece(3, 3, board.Yellow)
	win = TestWinner(b, 2, true)
	if win[0] != board.None {
		t.Error("Didn't block")
	}
}

func TestTestWinnerDR(t *testing.T) {
	b := board.MakeBoard(2)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(5-x, 1+x, board.Black)
		b.PlacePiece(4-x, x, board.Yellow)
	}

	win := TestWinner(b, 2, true)
	if win[1] != board.Black && win[2] != board.Black && win[0] != board.Tie {
		t.Error("Didn't see diag Black")
	}
}

func TestTestWinnerR(t *testing.T) {
	b := board.MakeBoard(3)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(8-x, 0, board.Black)
		b.PlacePiece(x, 1+x, board.Red)
	}

	win := TestWinner(b, 2, true)
	if win[1] != board.Black && win[2] != board.Black && win[0] != board.Tie {
		t.Error("Didn't see diag Black")
	}
}

func TestTestWinner3x3(t *testing.T) {
	b := board.MakeBoard(3)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(x, 7, board.Black)
	}

	win := TestWinner(b, 2, true)
	if win[0] != board.Black {
		t.Error("Didn't see Black")
	}
}

func TestTestWinnerDiagLeft(t *testing.T) {
	b := board.MakeBoard(2)
	b.PlacePiece(3, 0, board.Black)
	for x := 0; x < 5; x++ {
		b.PlacePiece(x, 5-x, board.Black)
	}

	win := TestWinner(b, 2, true)
	if win[0] != board.Black {
		t.Error("Didn't see Black")
	}
}

func TestTestWinnerTie(t *testing.T) {
	b := board.MakeBoard(3)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(x, 6, board.Black)
		b.PlacePiece(x, 7, board.Red)
		b.PlacePiece(x, 8, board.Blue)
	}
	win := TestWinner(b, 3, true)
	if win[0] != board.Tie || win[1] != board.Black || win[2] != board.Red || win[3] != board.Blue {
		t.Error("Didn't see 3 tie players in correct order.")
	}
}

func TestWeirdGreen(t *testing.T) {
	b := board.MakeBoard(3)
	setupGreenBoard(b)

	win := TestWinner(b, 2, true)
	if win[0] != board.Tie && win[0] != board.None {
		t.Error("None one should win.")
	}
}

func TestPermute(t *testing.T) {
	b := board.MakeBoard(2)
	actions := AvailActions(b, board.Red)
	// Can't test deterministicaly.
	// fmt.Println(actions)
	Permute(actions)
	// fmt.Println("new:")
	// fmt.Println(actions)
}

func setupGreenBoard(b *board.Board) {
	b.PlacePiece(1, 0, board.Red)
	b.PlacePiece(7, 0, board.Red)
	b.PlacePiece(8, 1, board.Red)
	b.PlacePiece(0, 2, board.Red)
	b.PlacePiece(5, 2, board.Red)
	b.PlacePiece(3, 3, board.Red)
	b.PlacePiece(2, 4, board.Red)
	b.PlacePiece(8, 4, board.Red)
	b.PlacePiece(0, 7, board.Red)
	b.PlacePiece(1, 7, board.Red)
	b.PlacePiece(1, 8, board.Red)
	b.PlacePiece(5, 8, board.Red)

	b.PlacePiece(5, 0, board.Green)
	b.PlacePiece(2, 1, board.Green)
	b.PlacePiece(4, 1, board.Green)
	b.PlacePiece(5, 1, board.Green)
	b.PlacePiece(2, 2, board.Green)
	b.PlacePiece(3, 2, board.Green)
	b.PlacePiece(6, 3, board.Green)
	b.PlacePiece(4, 4, board.Green)
	b.PlacePiece(7, 4, board.Green)
	b.PlacePiece(6, 5, board.Green)
	b.PlacePiece(3, 6, board.Green)
	b.PlacePiece(2, 7, board.Green)

	b.PlacePiece(0, 1, board.Blue)
	b.PlacePiece(4, 2, board.Blue)
	b.PlacePiece(6, 2, board.Blue)
	b.PlacePiece(0, 3, board.Blue)
	b.PlacePiece(7, 3, board.Blue)
	b.PlacePiece(8, 3, board.Blue)
	b.PlacePiece(2, 5, board.Blue)
	b.PlacePiece(8, 5, board.Blue)
	b.PlacePiece(0, 6, board.Blue)
	b.PlacePiece(8, 6, board.Blue)
	b.PlacePiece(5, 7, board.Blue)
	b.PlacePiece(2, 8, board.Blue)

	b.PlacePiece(3, 0, board.Yellow)
	b.PlacePiece(6, 0, board.Yellow)
	b.PlacePiece(1, 1, board.Yellow)
	b.PlacePiece(7, 1, board.Yellow)
	b.PlacePiece(1, 3, board.Yellow)
	b.PlacePiece(4, 5, board.Yellow)
	b.PlacePiece(5, 5, board.Yellow)
	b.PlacePiece(7, 5, board.Yellow)
	b.PlacePiece(2, 6, board.Yellow)
	b.PlacePiece(3, 7, board.Yellow)
	b.PlacePiece(7, 7, board.Yellow)
	b.PlacePiece(6, 8, board.Yellow)
}
