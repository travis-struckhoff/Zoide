package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"subversion.ews.illinois.edu/svn/fa14-cs242/struckh2/Zoide/board"
)

// Get the coordinates from the user.
func getCoord(in *bufio.Reader) (int, int) {
	noQuit := true
	var coords []string
	for noQuit {
		input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.Trim(input, "\n")
		coords = strings.Split(input, ",")
		if len(coords) != 2 {
			fmt.Print("Re-enter:")
		} else {
			noQuit = false
		}
	}
	x64, _ := strconv.ParseInt(coords[0], 10, 32)
	y64, _ := strconv.ParseInt(coords[1], 10, 32)
	x, y := int(x64), int(y64)
	//fmt.Println(x,y, coords)
	return x, y
}

// Get the rotation from the user.
func getRotate(in *bufio.Reader) (int, int, bool) {
	var coords []string
	for {
		input, err := in.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.Trim(input, "\n")
		coords = strings.Split(input, ",")
		if len(coords) != 3 {
			fmt.Print("Re-enter:")
		} else {
			break
		}
	}
	x64, _ := strconv.ParseInt(coords[0], 10, 32)
	y64, _ := strconv.ParseInt(coords[1], 10, 32)
	rot, _ := strconv.ParseInt(coords[2], 10, 32)
	x, y := int(x64), int(y64)
	//fmt.Println(x,y, coords)
	return x, y, rot == 0
}

// Run the main game loop.
func bmain() {
	in := bufio.NewReader(os.Stdin)
	b := board.MakeBoard(2)

	player1 := true
	for {
		fmt.Println(b)

		tryAgain := true
		for tryAgain {
			if player1 {
				fmt.Print("Player 1 Coord: ")
			} else {
				fmt.Print("Player 2 Coord: ")
			}
			x, y := getCoord(in)

			if player1 {
				tryAgain = !b.PlacePiece(x, y, board.Red)
			} else {
				tryAgain = !b.PlacePiece(x, y, board.Blue)
			}
		}

		win := TestWinner(b, 2, false)
		if win[0] != board.None {
			fmt.Println(b)
			fmt.Println("Winner:", win)
			break
		}

		fmt.Println(b)

		tryAgain = true
		for tryAgain {
			if player1 {
				fmt.Print("Player 1 Rot: ")
			} else {
				fmt.Print("Player 2 Rot: ")
			}

			x, y, rot := getRotate(in)
			tryAgain = !b.RotatePanel(x, y, rot)
		}

		win = TestWinner(b, 2, true)
		if win[0] != board.None {
			fmt.Println(b)
			fmt.Println("Winner:", win)
			break
		}

		player1 = !player1
	}
}
