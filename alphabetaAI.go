package main

import (
	"fmt"
	"math"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
)

// Some weights for utility.
const (
	WIN  = 1000
	LOST = -1000
	TIE  = 0
	// Guesses:
	GWIN   = 500
	GLOST  = -500
	ALMOST = 50
	CLOSE  = 1
)

type Action struct {
	placeX    int
	placeY    int
	turnX     int
	turnY     int
	direction bool
}

// Get the action that gives the best outcome.
func Alpha_beta_action(g *Game, max_depth int) Action {
	best_action := Action{}
	best_utility := math.MinInt32
	player := board.Player(g.whichPlayer)

	// For every available action, choose the action that gives the most utility.
	actions := AvailActions(&g.Model, player)

	for _, action := range actions {
		nextState := ResultState(&g.Model, action, player)
		utility := ab_maxmin_value(g, nextState, player, math.MinInt32, math.MaxInt32, max_depth, false)

		if utility > best_utility {
			best_utility = utility
			best_action = action
		}
	}

	fmt.Println("Best Util: ", best_utility)

	return best_action
}

// Get the next utility for yourself or enemy.
func ab_maxmin_value(g *Game, state *board.Board, player board.Player, alpha, beta, max_depth int, getMax bool) int {
	// No futher actions, return utility for this state.
	if state.IsLeaf() {
		return Utility(g, state, player)
	}

	// Have to quit early, guess what the utility would be.
	if max_depth == 1 {
		return Evaluation(state, player)
	}

	utility := math.MaxInt32
	if getMax {
		utility = math.MinInt32
	}

	// Get the action that maximises or minimizes utility.
	for _, action := range AvailActions(state, player) {
		nextState := ResultState(state, action, player)
		nextPlayer := nextPlayer(g, player)
		result := ab_maxmin_value(g, nextState, nextPlayer, alpha, beta, max_depth-1, !getMax)

		if getMax {
			utility = max(utility, result)
			// This utility is better than anything in the other branches
			if utility >= beta {
				return utility
			}

			alpha = max(alpha, utility)
		} else {
			utility = min(utility, result)
			if utility <= alpha {
				return utility
			}
			beta = min(beta, utility)
		}
	}

	return utility
}

// Returns the next player to play.
func nextPlayer(g *Game, cur board.Player) board.Player {
	intCur := int(cur)
	if intCur < g.numPlayers {
		return board.Player(intCur + 1)
	}
	return 1
}

// Max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Returns an array of available actions.
func AvailActions(b *board.Board, player board.Player) []Action {
	actions := make([]Action, 0, (b.Size()*3)*(b.Size()*3))
	for x := 0; x < b.Size()*3; x++ {
		for y := 0; y < b.Size()*3; y++ {
			for turnX := 0; turnX < 2; turnX++ {
				for turnY := 0; turnY < 2; turnY++ {
					// Append placing a marble here if available.
					if b.GetPiece(x, y) == board.None {
						cur_action := Action{
							placeX:    x,
							placeY:    y,
							turnX:     turnX,
							turnY:     turnY,
							direction: true,
						}
						actions = append(actions, cur_action)
						cur_action.direction = false
						actions = append(actions, cur_action)
					}
				}
			}
		}
	}
	Permute(actions)
	return actions
}

// Return the next state from taking the action.
func ResultState(b *board.Board, action Action, player board.Player) *board.Board {
	// Take the action on a copy of the state.
	newState := b.Copy()
	newState.PlacePiece(action.placeX, action.placeY, player)
	newState.RotatePanel(action.turnX, action.turnY, action.direction)

	return newState
}

// Computes the utility for a engame arrangement.
func Utility(g *Game, b *board.Board, player board.Player) int {
	winners := TestWinner(b, g.numPlayers, true)

	if winners[0] == player {
		return WIN
	} else if winners[0] == board.Tie {
		for _, p := range winners[1:] {
			if p == player {
				return TIE
			}
		}
	}
	return LOST
}

// Guess how much the state is worth. Heuristics...
func Evaluation(b *board.Board, player board.Player) int {
	eval := ComputeEval(b, player)

	return eval
}
