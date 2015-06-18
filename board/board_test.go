package board

import (
	"testing"
	//"fmt"
)

func TestPanelRotate(t *testing.T) {
	panel := new(panel)
	panel.square[0][0] = 1
	panel.square[1][1] = 9
	panel.square[1][0] = 2

	panel.RotateClock()
	panel.RotateCClock()
	if panel.square[0][0] != 1 {
		t.Error("Rotations don't work...")
	}
	panel.RotateCClock()
	if panel.square[2][0] != 1 {
		t.Error(panel, "CounterClock...")
	}
}

func TestPlacePiece(t *testing.T) {
	board := MakeBoard(3)
	board.PlacePiece(0,1, Black)
	board.PlacePiece(0, 6, Red)
	if board.GetPiece(0,1) != Black {
		t.Error("3, 0,1, Black")
	}
	if board.GetPiece(0,6) != Red {
		t.Error("3, 0,6, Red")
	}

	board = MakeBoard(2)
	board.PlacePiece(5, 5, Red)
	if board.GetPiece(5,5) != Red {
		t.Error("2, 5,5, Red")
	}
}

func TestRotatePanel(t *testing.T) {
	board := MakeBoard(2)
	board.PlacePiece(0,1,Black)
	board.RotatePanel(0,0,false)
	if board.GetPiece(1,0) != Black {
		t.Error("2, RC:Not Black")
	}

	board = MakeBoard(3)
	board.PlacePiece(6,6, Yellow)
	board.RotatePanel(2,2, false)
	if board.GetPiece(8,6) != Yellow {
		t.Error("3, RotateC")
	}
}
