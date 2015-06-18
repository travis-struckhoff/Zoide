package main

/**
*	Sources: https://github.com/tcoopman/samegame
*		for how to build the game in qml and go.
*
*/

import (
	"fmt"
	"os"
	"time"
	"strconv"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	"gopkg.in/qml.v1"
)

// Time for twisting a panel.
var twist bool
// which player's turn.
var player1 bool


type Game struct {
	Model board.Board
	Board []qml.Object
	Marble *Marble
	parent qml.Object
	dialog qml.Object
	Score qml.Object
	started bool
	boardSize int
}

// Helper for getting the index in 1D array.
func (g *Game) index(col, row int) int {
	return col + (row * g.Model.Size()*3)
}

func (g *Game) StartNewGame(parent, dialog qml.Object) {
	for _, marble := range g.Board {
		if marble != nil {
			marble.Destroy()
		}
	}

	player1 = true
	twist = false

	g.parent = parent
	g.dialog = dialog
	g.Model = *board.MakeBoard(g.boardSize)

	score := 0
	g.parent.Set("score", score)
	g.Score.Set("text", "Score: " + strconv.Itoa(score))

	g.Marble.Size = parent.Int("blockSize")
	
	// Fill out the empty board.
	colSize := g.Model.Size()*3
	numMarbles := colSize*colSize
	g.Board = make([]qml.Object, numMarbles)
	for col := 0; col < colSize; col++ {
		for row := 0; row < colSize; row++ {
			g.Board[g.index(col, row)] = g.Marble.createMarble(col, row, parent)
		}
	}
	g.started = true
}

// Handle clicking on a marble location.
func (g *Game) HandleClick(xPos, yPos int) {
	if !g.started || twist {
		return
	}

	col := xPos / g.Marble.Size
	row := yPos / g.Marble.Size

	var player board.Player
	if player1 {
		player = board.Red
	} else {
		player = board.Blue
	}

	success := g.Model.PlacePiece(col, row, player)
	// Failed to place a piece
	if !success {
		return
	}

	g.Board[g.index(col, row)].Set("type", int(player))
	g.Board[g.index(col, row)].Set("exists", true)
	player1 = !player1
	twist = true
	checkForWin(g)
}

// Twist the panels with the buttons.
func (g *Game) HandleArrows(arrow qml.Object) {
	if !g.started || !twist {
		return
	}

	if arrow.Int("arrow") == 0 {
		// Top Left turns CounterClock
		g.Model.RotatePanel(0,0, true)
	} else if arrow.Int("arrow") == 1 {
		// Top Left turns Clock
		g.Model.RotatePanel(0,0, false)
	} else if arrow.Int("arrow") == 2 {
		// Top Right turns CounterClock
		g.Model.RotatePanel(1,0, true)
	} else if arrow.Int("arrow") == 3 {
		// Top Right turns Clock
		g.Model.RotatePanel(1,0, false)
	} else if arrow.Int("arrow") == 4 {
		// Bottom Left turns CounterClock
		g.Model.RotatePanel(0,1, true)
	} else if arrow.Int("arrow") == 5 {
		// Bottom Left turns Clock
		g.Model.RotatePanel(0,1, false)
	} else if arrow.Int("arrow") == 6 {
		// Bottom Right turns CounterClock
		g.Model.RotatePanel(1,1, true)
	} else if arrow.Int("arrow") == 7 {
		// Bottom Right turns Clock
		g.Model.RotatePanel(1,1, false)
	}
	g.RedrawBoard()
	twist = false

	checkForWin(g)
}

// Check if there is a win.
func checkForWin(g *Game) {
	win := TestWinner(&g.Model, 2)
	if win[0] != board.None {
		g.started = false
		player := "Tie"
		if win[0] == board.Red {
			player = "Red"
		} else if win[0] == board.Blue {
			player = "Blue"
		}

		g.dialog.Call("show", "Game over. "+ player +" player won! ")
		go func() {
			opened := time.Now()
			for time.Now().Sub(opened) < time.Second*10 {
			}
			g.dialog.Call("hide")
		}()
	}
}

// Updates the colors and view of the marbles.
func (g *Game) RedrawBoard() {
	colSize := g.Model.Size()*3
	for col := 0; col < colSize; col++ {
		for row := 0; row < colSize; row++ {
			// Marble color in view
			marble := g.Board[g.index(col, row)]
			mPlayer := marble.Int("type")
			// Marble in model
			player := int(g.Model.GetPiece(col, row))

			// Change the marble if its player changed.
			if player == board.None {
				marble.Set("exists", false)
			} else if player != mPlayer {
				marble.Set("type", player)
				marble.Set("exists", true)
			}
		}
	}
}

// Marble access struct.
type Marble struct {
	Component 	qml.Object
	Size 		int
}

// Create.
func (m *Marble) createMarble(col, row int, parent qml.Object) qml.Object {
	marble := m.Component.Create(nil)
	marble.Set("parent", parent)
	marble.Set("type", board.None)
	marble.Set("x", col*m.Size)
	marble.Set("y", row*m.Size)
	marble.Set("width", m.Size)
	marble.Set("height", m.Size)
	marble.Set("exists", false)
	return marble
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	engine := qml.NewEngine()

	// Load the main layout.
	mainComponent, err := engine.LoadFile("main.qml")
	if err != nil {
		return err
	}

	// Set the game and callback for the main layout to "game"
	game := Game {boardSize: 2}

	context := engine.Context()
	context.SetVar("game", &game)

	window := mainComponent.CreateWindow(nil)

	// Load the marble layout.
	marbleComponent, err := engine.LoadFile("Marble.qml")
	if err != nil {
		return err
	}

	// Add marbles to the game.
	marble := &Marble{Component: marbleComponent}
	game.Marble = marble
	game.Score = window.Root().ObjectByName("score")

	window.Show()
	window.Wait()
	return nil
}
