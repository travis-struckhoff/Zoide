package main

import (
	"fmt"
	"testing"
	"net"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
)

func TestActionToString(t *testing.T) {
	action := LanAction{1,1,1,1,1,true}
	result := action.String()
	if result != "1,1,1,1,1,t" {
		t.Error("Not correct:", result)
	}

	action = LanAction{1,2,1,2,1,false}
	result = action.String()
	if result != "1,2,1,2,1,f" {
		t.Error("Not correct:", result)
	}
}

func TestStringToAction(t *testing.T) {
	action := "1,1,1,1,1,t "
	result := StringToLanAction(action)
	req := LanAction{1,1,1,1,1,true}
	if result != req {
		t.Error("Didn't convert str->action")
	}

	action = "1,2,1,2,1,f"
	result = StringToLanAction(action)
	req = LanAction{1,2,1,2,1,false}
	if result != req {
		t.Error("Didn't convert str->action")
	}
}

func TestMakeRequest(t *testing.T) {
	g := setupGame()
	act := LanAction{1,1,1,1,1,true}
	MakeRequest(g, act)
	if g.whichPlayer == 1 {
		t.Error("Failed move.")
	}

}

func TestMakeResult(t *testing.T) {
	g := setupGame()
	act := LanAction{1,1,1,1,1,true}
	MakeResult(g, act)
	if g.whichPlayer == 1 {
		t.Error("Failed move.")
	}
}

func TestDoAction(t *testing.T) {
	g := setupGame()
	act := LanAction{1,1,1,1,1,true}
	DoAction(g, act)
	if g.whichPlayer == 1 {
		t.Error("Failed move.")
	}
}

func setupGame() *Game {
	b := board.MakeBoard(2)

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
		fmt.Println(err)
		return nil
	}
	defer ln.Close()

	c, err := Connect("localhost:8080")
	if err != nil {
		return nil
	}
	g := Game{
		Model: *b,	
		started: true,
		numPlayers: 2,
		whichPlayer: 1,  
		twistTime: false, 
		aiPlaying: false,
		lanGame: true,
		server: false,
		hostsTurn: true,
		conn: c,
	}
	return &g
}