package main

/**
*	Sources: https://github.com/tcoopman/samegame
*		for how to build the game in qml and go.
*
 */

import (
	"fmt"
	"gopkg.in/qml.v1"
	"os"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
	"time"
	"net"
)

type Game struct {
	Model       board.Board
	Board       []qml.Object
	Marble      *Marble
	parent      qml.Object
	dialog      qml.Object
	Score       qml.Object
	started     bool
	numPlayers  int
	whichPlayer int  // This player's turn.
	twistTime   bool // Time for twisting panel.
	numAI       int
	aiPlaying   bool

	// For multi over lan
	lanGame 	bool
	server 		bool
	hostsTurn 	bool
	conn 		*net.Conn
}

// Helper to decide who's turn. Keeps player within 1 and numPlayers inclusive.
func (g *Game) incPlayer() {
	if g.whichPlayer < g.numPlayers {
		g.whichPlayer++
	} else {
		g.whichPlayer = 1
	}
}

// Helper for getting the index in 1D array.
func (g *Game) index(col, row int) int {
	return col + (row * g.Model.Size() * 3)
}

// Starts the game,
// lan - is it over lan
// host - am I the host
func (g *Game) StartNewGame(parent, dialog qml.Object, numPlayers, boardSize int, lan, host bool, url string) {
	for _, marble := range g.Board {
		if marble != nil {
			marble.Destroy()
		}
	}

	g.whichPlayer = 1
	g.twistTime = false
	g.numPlayers = numPlayers
	g.numAI = 0
	g.aiPlaying = false
	if numPlayers == 1 {
		g.numPlayers = 2
		g.numAI = 1
	}

	g.parent = parent
	g.dialog = dialog
	g.Model = *board.MakeBoard(boardSize)

	g.Marble.Size = parent.Int("blockSize")

	// Fill out the empty board.
	colSize := g.Model.Size() * 3
	numMarbles := colSize * colSize
	g.Board = make([]qml.Object, numMarbles)
	for col := 0; col < colSize; col++ {
		for row := 0; row < colSize; row++ {
			g.Board[g.index(col, row)] = g.Marble.createMarble(col, row, parent)
		}
	}
	g.started = true

	// Defaults
	g.lanGame = false
	g.server = false
	g.hostsTurn = true

	if lan {
		g.lanGame = true
		if host {
			// Start server
			go CreateServer(g)
			g.server = true

		} else {
			//runningServer = false
			conn, err := Connect(url)
			// Fix for weird multi assingment fluke ^
			g.conn = conn
			if err != nil {
				os.Exit(1)
			}
		}
	}
}

// Handle clicking on a marble location.
func (g *Game) HandleClick(xPos, yPos int) {
	if !g.started || g.twistTime || g.aiPlaying || !g.hostsTurn {
		return
	}

	col := xPos / g.Marble.Size
	row := yPos / g.Marble.Size

	// If lanGame do server stuff
	if g.lanGame {
		if !g.server {
			// Issue the request to the server and wait for a response
			action := LanAction{int(g.whichPlayer), col, row, -1, -1, false}
			MakeRequest(g, action)
			return
		}

		// Server stuff
		action := LanAction{int(g.whichPlayer), col, row, -1, -1, false}
		MakeResult(g, action)
		return
	}

	g.placePieceInGame(col, row)
}

func (g *Game) placePieceInGame(col, row int) bool {
	success := g.Model.PlacePiece(col, row, board.Player(g.whichPlayer))
	// Failed to place a piece
	if !success {
		return false
	}

	g.Board[g.index(col, row)].Set("type", g.whichPlayer)
	g.Board[g.index(col, row)].Set("exists", true)
	g.twistTime = true
	g.checkForWin()
	return true
}

// Twist the panels with the buttons.
func (g *Game) HandleArrows(arrow qml.Object) {
	if !g.started || !g.twistTime || g.aiPlaying || !g.hostsTurn {
		return
	}
	x := 0
	y := 0
	direction := false
	offset := 1

	switch arrow.Int("arrow") {
	case 0: // Top Left turns CounterClock
		x, y, direction = 0, 0, true
	case 1: // Top Left turns Clock
		x, y, direction = 0, 0, false
	case 2: // Top Right turns CounterClock
		x, y, direction = offset, 0, true
	case 3: // Top Right turns Clock
		x, y, direction = offset, 0, false
	case 4: // Bottom Left turns CounterClock
		x, y, direction = 0, offset, true
	case 5: // Bottom Left turns Clock
		x, y, direction = 0, offset, false
	case 6: // Bottom Right turns CounterClock
		x, y, direction = offset, offset, true
	case 7: // Bottom Right turns Clock
		x, y, direction = offset, offset, false
	}

	// If lanGame do server stuff
	if g.lanGame {
		if !g.server {
			// Issue the request to the server and wait for a response
			action := LanAction{int(g.whichPlayer), -1, -1, x, y, direction}
			MakeRequest(g, action)
			return
		}

		// Server stuff
		action := LanAction{int(g.whichPlayer), -1, -1, x, y, direction}
		MakeResult(g, action)
		return
	}

	g.rotateGamePanel(x, y, direction)		

	go g.AI() // Run the ai in another thread.
}

func (g *Game) rotateGamePanel(x, y int, direction bool) bool {
	success := g.Model.RotatePanel(x, y, direction)
	if !success {
		return false
	}

	g.RedrawBoard()
	g.twistTime = false

	g.checkForWin()
	g.incPlayer()
	return true
}

// Check if there is a win.
func (g *Game) checkForWin() {
	win := TestWinner(&g.Model, g.numPlayers, !g.twistTime)
	if win[0] != board.None {
		fmt.Println(win)
		g.started = false
		var winString string
		if win[0] != board.Tie {
			winString = "Game over. " + win[0].GetPlayerString() + " player won!"
		} else {
			winString = "Game over. Tie Players: "
			for _, player := range win[1:] {
				if player != board.None {
					winString += player.GetPlayerString() + ", "
				}
			}
			winString += "won!"
		}
		g.dialog.Call("show", winString)
		go func() {
			time.Sleep(time.Second * 4)
			g.dialog.Call("hide")
		}()

		if win[0] != board.Tie {
			g.Score.Set("text", win[0].GetPlayerString()+" player won last.")
		} else {
			g.Score.Set("text", "Last game was a tie.")
		}
	}
}

// Updates the colors and view of the marbles.
func (g *Game) RedrawBoard() {
	colSize := g.Model.Size() * 3
	for col := 0; col < colSize; col++ {
		for row := 0; row < colSize; row++ {
			// Marble color in view
			marble := g.Board[g.index(col, row)]
			// Marble in model
			player := int(g.Model.GetPiece(col, row))

			// Change the marble if its player changed.
			if player == board.None {
				marble.Set("exists", false)
			} else {
				marble.Set("type", player)
				marble.Set("exists", true)
			}
		}
	}
}

// AI handler.
func (g *Game) AI() {
	if g.started && g.numAI == 1 && g.whichPlayer == 2 {
		g.aiPlaying = true
		g.dialog.Call("show", "AI's Turn...")

		start := time.Now()
		action := Alpha_beta_action(g, 2)
		fmt.Println("Time: ", time.Since(start))
		fmt.Println("Action: ", action)

		g.placePieceInGame(action.placeX, action.placeY)

		// Don't turn if already won.
		if g.started {
			g.rotateGamePanel(action.turnX, action.turnY, action.direction)
		}
		g.dialog.Call("hide")
		g.aiPlaying = false
	}
}

// Marble access struct.
type Marble struct {
	Component qml.Object
	Size      int
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
	var game Game

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
