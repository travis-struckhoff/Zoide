package board

import "fmt"

// Enumerate some player colors
type Player int
const (
	None = iota // No player
	Red
	Blue
	Yellow
	Green
	Black
	White
	Tie 		// Tie players
)

func (p Player) GetPlayerString() string {
	types := [...]string{
		"No one",
		"Red",
		"Blue",
		"Yellow",
		"Green",
		"Black",
		"White",
		"Tie",
	}
	return types[int(p)]
}

// Class for holding the 3x3 squares.
type panel struct {
	square [3][3]Player
}

// Rotate the panel clockwise.
func (p *panel) RotateClock() {
	newSqaure := [3][3]Player{}
	for i := 0; i < 3; i++ {
		newSqaure[0][i] = p.square[2-i][0]
		newSqaure[1][i] = p.square[2-i][1]
		newSqaure[2][i] = p.square[2-i][2]
	}
	p.square = newSqaure
}

// Rotate the panel counter-clockwise.
func (p *panel) RotateCClock() {
	newSqaure := [3][3]Player{}
	for i := 0; i < 3; i++ {
		newSqaure[0][i] = p.square[i][2]
		newSqaure[1][i] = p.square[i][1]
		newSqaure[2][i] = p.square[i][0]
	}
	p.square = newSqaure
}

// Class for holding the default board of 4 panels.
// 0,0 top left corner; 1,1 bottom right, (y,x) access pattern.
type Board struct {
	board [][]panel 
	size int
}

// Construct the board. Size is panel dimensions, ie. 2x2 panels.
func MakeBoard(size int) *Board {
	b := new(Board)
	b.board = make([][]panel, size)
	for i := 0; i < size; i++ {
		b.board[i] = make([]panel, size)
	}
	b.size = size
	return b
}

// Get the dimension of the board.
func (b *Board) Size() int {
	return b.size
}

// Place a piece at the requested location on the board.
// Return false if x,y is invalid.
func (b *Board) PlacePiece(x, y int, player Player) bool {
	ptr := b.accessPiece(x,y)
	if ptr == nil || *ptr != None {
		return false
	}

	*ptr = player
	
	return true
}

// Get a pointer to the piece at x,y.
func (b *Board) accessPiece(x,y int) *Player {
	if !InRange(0, b.size*3-1, x) || !InRange(0, b.size*3-1, y) {
		return nil
	}

	col := x / 3
	row := y / 3
	inX := x % 3
	inY := y % 3
	//fmt.Printf("Col:%v, row:%v, inX:%v, inY:%v\nX:%v, Y:%v\n", col, row, inX, inY, x, y)
	return &b.board[row][col].square[inY][inX]
}

// Get read access to piece on board.
func (b *Board) GetPiece(x,y int) Player {
	ptr := b.accessPiece(x,y)
	if ptr == nil {
		return None
	}
	return *ptr
}

// Rotate the panel at x,y clockwise if dir=false, counter-clock if true.
func (b *Board) RotatePanel(x, y int, direction bool) bool {
	if !InRange(0, b.size-1,x) || !InRange(0, b.size-1,y) {
		return false
	}

	// For now use inputs of 0 and 1, and scale with boardsize.
	if b.size > 2 {
		x = x*(b.size-1)
		y = y*(b.size-1)
	}

	if direction {
		b.board[y][x].RotateCClock()
	} else {
		b.board[y][x].RotateClock()
	}
	return true
}

// Return true if number is in range inclusive.
func InRange(start, end, num int) bool {
	return num <= end && num >= start
}

// Print out the board.
func (b *Board) String() string {
	// Put helper numbers.
	output := "|-0-1-2---3-4-5-|\n"
	num := 0

	for set := 0; set < b.size; set++ {
		for row := 0; row < 3; row++ {
			for pan := 0; pan < b.size; pan++ {
				if pan == 0 {
					// Put helper numbers.
					output += fmt.Sprintf("%v ", num)
					num++
				}
				for col := 0; col < 3; col++ {
					output += fmt.Sprintf("%v ", 
						b.board[set][pan].square[row][col])
				}
				output += "| "
			}
			output += "\n"
		}
			output += "|---------------|\n"
	}
	return output
}

// Tells AI if the board is finished with moves.
func (b *Board) IsLeaf() bool {
	for x := 0; x < b.size*3; x++ {
		for y := 0; y <b.size*3; y++ {
			// Blank spots left to be filled.
			if b.GetPiece(x, y) == None {
				return false
			}
		}
	}
	return true
}

// Return a copy of the board.
func (b *Board) Copy() *Board {
	newBoard := MakeBoard(b.size)

	for x := 0; x < b.size*3; x++ {
		for y := 0; y <b.size*3; y++ {
			newBoard.PlacePiece(x, y, b.GetPiece(x,y))
		}
	}
	return newBoard
}