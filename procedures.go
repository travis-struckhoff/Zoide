package main

import (
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	//"fmt"
)

// Some directions for the where the 5 in a row is.
type Direction int
const (
	Down = iota
	DownLeft
	DownRight
	Right
)

// Look for the winner on the board.
// Return the winners or None for no winner.
func TestWinner(b *board.Board, num_players int) []board.Player {
	winners := make([]board.Player, num_players+1)
	winners[0] = board.None
	num_win := 0
	var currentPlayer board.Player
	dim := b.Size()*3
	players_tested := make(map[board.Player]bool) // set of players tested.

	for nP := 0; nP < num_players; nP++ {
		// Scan the board.
		for y := 0; y < dim; y++ {
			for x := 0; x < dim; x++ {
				piece := b.GetPiece(x,y)
				
				// Test the piece
				if piece != board.None && !players_tested[piece]{
					currentPlayer = piece
					players_tested[piece] = true

					win,_ := search5Row(b, x, y, currentPlayer)

					if win {
						// A winner.
						num_win++
						if winners[0] == board.None {
							winners[0] = currentPlayer
						} else if winners[0] != board.Tie { // Handle ties.
							winners[1] = winners[0]
							winners[0] = board.Tie
							winners[num_win] = currentPlayer
						} else {
							winners[num_win] = currentPlayer
						}
					}
				}
				
				
			}
		}
	}
	return winners
}

// Search around the piece for 5 in a row from the top down only.
// Return true if there exists a 5 in a row
func search5Row(b *board.Board, startX, startY int, p board.Player) (bool, Direction) {
	number := 0
	// Search column down
	for y := startY; y < b.Size()*3; y++ {
		if b.GetPiece(startX, y) == p {
			number++
			if number == 5 {
				return true, Down
			}
		} else {
			break
		}
	}

	// Search row right
	number = 0
	for x := startX; x < b.Size()*3; x++ {
		if b.GetPiece(x, startY) == p {
			number++
			if number == 5 {
				return true, Right
			}
		} else {
			break
		}
	}

	// Search down left
	number = 0
	dontQuit := true
	for i := 0; i < b.Size()*3 && dontQuit; i++ {
		if b.GetPiece(startX-i, startY+i) == p {
			number++
			if number == 5 {
				return true, DownLeft
			}
		} else {
			dontQuit = false
		}
	}

	// Search down right
	number = 0
	dontQuit = true
	for i := 0; i < b.Size()*3 && dontQuit; i++ {
		if b.GetPiece(startX+i, startY+i) == p {
			number++
			if number == 5 {
				return true, DownLeft
			}
		} else {
			dontQuit = false
		}
	}
	return false, Down
}