package main

import (
	"fmt"
	"math"
	//"os"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	//"time"
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
	for _, action := range g.AvailActions(&g.Model, player) {
		nextState := g.ResultState(&g.Model, action, player)
		utility := ab_maxmin_value(g, nextState, player, math.MinInt32, math.MaxInt32, max_depth, false)

		if utility > best_utility {
			best_utility = utility
			best_action = action
		}
		// fmt.Println("Util: ", utility)
		// fmt.Println(nextState)
	}

	fmt.Println("Best Util: ", best_utility)

	return best_action
}

// Get the next utility for yourself or enemy.
func ab_maxmin_value(g *Game, state *board.Board, player board.Player, alpha, beta, max_depth int, getMax bool) int {
	// No futher actions, return utility for this state.
	if state.IsLeaf() {
		return g.Utility(state, player)
	}

	// Have to quit early, guess what the utility would be.
	if max_depth == 1 {
		return g.Evaluation(state, player)
	}

	utility := math.MaxInt32
	if getMax {
		utility = math.MinInt32
	}

	// Get the action that maximises or minimizes utility.
	for _, action := range g.AvailActions(state, player) {
		result := ab_maxmin_value(g, g.ResultState(state, action, player), player, alpha, beta, max_depth-1, !getMax)

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
