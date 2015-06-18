package main

import (
	"math/rand"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	"time"
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
func TestWinner(b *board.Board, num_players int, didTwist bool) []board.Player {
	winners := make([]board.Player, num_players+1)
	winners[0] = board.None
	num_win := 0
	var currentPlayer board.Player
	dim := b.Size() * 3
	players_tested := make(map[board.Player]bool) // set of players tested.
	keepPlaying := false                          // If no more moves, stop playing

	for nP := 0; nP < num_players; nP++ {
		// Scan the board.
		for y := 0; y < dim; y++ {
			for x := 0; x < dim; x++ {
				piece := b.GetPiece(x, y)

				if piece == board.None {
					keepPlaying = true
				}

				// Test the piece
				if piece != board.None && !players_tested[piece] {
					currentPlayer = piece

					win, _ := SearchInARow(b, x, y, currentPlayer, 5)

					if win {
						players_tested[piece] = true

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
	if !keepPlaying && didTwist {
		return tieEndGame(b, num_players)
	}
	return winners
}

func tieEndGame(b *board.Board, num_players int) []board.Player {
	winners := make([]board.Player, num_players+1)
	winners[0] = board.Tie
	for p := 1; p < num_players+1; p++ {
		winners[p] = board.Player(p)
	}
	return winners
}

// Search around the piece for 'lineSize' in a row from the top down only.
// Return true if there exists a 'lineSize' in a row
func SearchInARow(b *board.Board, startX, startY int, p board.Player, lineSize int) (bool, Direction) {
	number := 0
	// Search column down
	for y := startY; y < b.Size()*3; y++ {
		if b.GetPiece(startX, y) == p {
			number++
			if number == lineSize {
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
			if number == lineSize {
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
			if number == lineSize {
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
			if number == lineSize {
				return true, DownLeft
			}
		} else {
			dontQuit = false
		}
	}
	return false, Down
}

// Make a guess about the utility for the action.
func ComputeEval(b *board.Board, player board.Player) int {
	eval := 0
	for x := 0; x < b.Size()*3; x++ {
		for y := 0; y < b.Size()*3; y++ {
			boardPlayer := b.GetPiece(x, y)
			if boardPlayer != board.None {
				for size := 2; size < 5; size++ {
					// Check the player for a 'size' in the row.
					if truth, _ := SearchInARow(b, x, y, boardPlayer, size); truth {
						if boardPlayer == player {
							eval += size
						} else {
							eval -= size * size * size
						}
					}
				}
			}
		}
	}
	return eval
}

func Permute(actions []Action) {
	rand.Seed(time.Now().Unix())
	size := len(actions)
	for i := 0; i < size; i++ {
		newPos := rand.Intn(size - 1)
		// Swap
		actions[i], actions[newPos] = actions[newPos], actions[i]
	}
}
