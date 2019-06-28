package main

import "fmt"
const exs = -1000000

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a int , b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// 1 means computer
// -1 means player
// 0 means blank


type GameBoard struct {
	board [3][3]int
	state int
}

func (game *GameBoard) printBoard() {
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			fmt.Printf("%d ", game.board[i][j])
		}
		fmt.Printf("\n")
	}
}

func (game *GameBoard) tie() bool {
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if game.board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func (game *GameBoard) calScore() ([3]int, [3]int, int, int) {
	var v, h [3]int
	var c1, c2 int
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			h[i] += game.board[i][j]
			v[i] += game.board[j][i]
		}
	}
	c1 = game.board[0][0] + game.board[1][1] + game.board[2][2]
	c2 = game.board[0][2] + game.board[1][1] + game.board[2][0]
	
	return v, h, c1, c2
}

func (game *GameBoard) isGameDone() int {
	v, h, c1, c2 := game.calScore()

	var flagWin, flagLost int
	flagWin = 10000
	flagLost = 10000

	for i:=0; i<3; i++ {
		if (v[i] == -2) || (h[i] == -2) {
			flagLost = -5
		}
		if (v[i] == 2) || (h[i] == 2) {
			flagWin = 5
		}
	}
	if (c1 == -2) || (c2 == -2) {
		flagLost = -5
	}
	if (c1 == 2) || (c2 == 2) {
		flagWin = 5
	}
	if flagWin < 100 || flagLost < 100 {
		return min(flagWin, flagLost)
	}

	return 0

}

func (game *GameBoard) isGameFinished() int {
	if game.tie() {
		fmt.Println("tie")
		return 0
	}
	v, h, c1, c2 := game.calScore()

	
	for i:=0; i<3; i++ {
		if (v[i] == -3) || (h[i] == -3) {
			return -1
		}
		if (v[i] == 3) || (h[i] == 3) {
			return 1
		}
	}
	if (c1 == -3) || (c2 == -3) {
		return -1
	}
	if (c1 == 3) || (c2 == 3) {
		return 1
	}

	return 0
}


func minimax(game GameBoard, depth int) int {
	var v int
	copy := game
	if game.tie() {
		return 0
	}
	if game.isGameDone() != 0 && depth > 0 {
		fmt.Println("Depth", depth)
		return game.isGameDone()
	}
	// simulate move
				
	// max
	v = exs


	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if copy.board[i][j] == 0 {
				copy.board[i][j] = copy.state
				copy.state = -copy.state

				fmt.Println("dd", depth)
				copy.printBoard()
				v = max(v, minimax(copy, depth+1))
				fmt.Println("v", v)
				
				copy.board[i][j] = 0
				copy.state = -copy.state
			}
		}
	}
	fmt.Println("minimax: ", v, " depth: ", depth)


	return v
}

func findMoveviaScore(game GameBoard, score int) (int, int) {
	v := exs
	for i:=0; i<3; i++ {
		for j:=0; j<3; j++ {
			if game.board[i][j] == 0 {
				game.board[i][j] = game.state
				game.state = -game.state

				v = max(v, minimax(game, 1))
				//fmt.Printf

				game.board[i][j] = 0
				game.state = -game.state

				if v == score {
					fmt.Println("Best Score: ", v)
					return i, j
				}
			}
		}
	}
	return 0, 0
}

func main() {
	game := new(GameBoard)
	game.board = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	game.printBoard()
	fmt.Println(game.isGameDone())
	game.state = -1
	for game.isGameFinished() == 0 && !game.tie() {
		if game.state == -1 {
			var c1, c2 int
			// Scan
			if _, err:= fmt.Scan(&c1, &c2); err != nil {
				panic(err)
			}

			//fmt.Printf("Entered %d%d", c1, c2)
			if game.board[c1][c2] == 0 {
				game.board[c1][c2] = -1
			}
		} else {
			// computer move
			x, y := findMoveviaScore(*game, minimax(*game, 0))
			fmt.Println("move taken: ", x, y)
			game.board[x][y] = game.state
		}
		game.state = -game.state

		game.printBoard()
		fmt.Println(game.isGameDone())
	}
	fmt.Println("Program Ended")

}

