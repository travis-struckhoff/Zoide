package main

import (
	"fmt"
	"net"
	//"io"
)

type LanAction struct {
	player		int
	placeX		int
	placeY		int
	turnX		int
	turnY		int
	direction	bool
}

// Convert Action to string.
func (a *LanAction) String() string {
	truth := 'f'
	if a.direction {
		truth = 't'
	}
	return fmt.Sprintf("%d,%d,%d,%d,%d,%c", a.player, a.placeX, a.placeY, a.turnX, a.turnY, truth)
}

// Convert string to action.
func StringToLanAction(s string) LanAction {
	a := LanAction{}
	truth := 't'
	fmt.Sscanf(s, "%d,%d,%d,%d,%d,%c", &a.player, &a.placeX, &a.placeY, &a.turnX, &a.turnY, &truth)
	if truth == 't' {
		a.direction = true
	} else {
		a.direction = false
	}

	return a
}

// Make list in future
var connections *net.Conn

func Connect(url string) (*net.Conn, error) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		// Post could not connect
		// Retry
		fmt.Println(err)
		fmt.Println("Could not connect to: ", url)
		return nil, err
	}
	return &conn, nil
}

// Create the server to run the game.
func CreateServer(g *Game) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	// Close the connection later when done.
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			fmt.Println(err)
			continue
		}
		go handleConnection(conn, g)
	}
}

// Handle the connnections from each client.
func handleConnection(conn net.Conn, g *Game) {
	// Parse command from connection //actions
	// Issue command to game
	// Issue result to connection.
	// Shut down the connection.
	connections = &conn
	defer conn.Close()
	for {

	}
}

// A client requests to make a move.
// See if it works and return if they succedded
func MakeRequest(g *Game, action LanAction) {
	// Send request
	fmt.Fprintf(*g.conn, "%v", &action)

	// Handle response
	response := true
	if response {
		DoAction(g, action)
	}
}

// The server user is making a move so tell the clients.
func MakeResult(g *Game, action LanAction) {
	if DoAction(g, action) {

		// Issue result
		fmt.Fprintf(*g.conn, "%v", &action)
	}
}

// Compute the action to make and take it in the game.
func DoAction(g *Game, action LanAction) bool {
	return true
}