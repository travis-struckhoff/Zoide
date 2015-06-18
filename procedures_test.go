package main

import (
	"testing"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	//"fmt"
)

func TestTestWinnerDL(t *testing.T) {
	b := board.MakeBoard(2)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(x,x,board.Black)
	}
	win := TestWinner(b, 2)
	if win[0] != board.Black {
		t.Error("Didn't see diag Black")
	}

	b = board.MakeBoard(2)
	for x := 0; x < 5; x++ {
		if x == 3 {
			continue
		}
		b.PlacePiece(x,x,board.Black)
	}
	b.PlacePiece(3,3,board.Yellow)
	//fmt.Println(b)
	win = TestWinner(b, 2)
	if win[0] != board.None {
		t.Error("Didn't block")
	}
}

func TestTestWinnerDR(t *testing.T) {
	b := board.MakeBoard(2)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(5-x,1+x,board.Black)
		b.PlacePiece(4-x,x,board.Yellow)
	}
	//fmt.Println(b)
	win := TestWinner(b, 2)
	if win[1] != board.Black && win[2] != board.Black && win[0] != board.Tie{
		t.Error("Didn't see diag Black")
	}
}

func TestTestWinnerR(t *testing.T) {
	b := board.MakeBoard(3)
	// Place a 5 in the row down right
	for x := 0; x < 5; x++ {
		b.PlacePiece(8-x,0,board.Black)
		b.PlacePiece(x,1+x,board.Red)
	}
	//fmt.Println(b)
	win := TestWinner(b, 2)
	if win[1] != board.Black && win[2] != board.Black && win[0] != board.Tie{
		t.Error("Didn't see diag Black")
	}
}