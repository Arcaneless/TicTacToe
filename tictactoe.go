package main

import "fmt"
const exs = -1000000

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
	board [9]int
	state int
}

func (game *GameBoard) printBoard() {
	for i:=0; i<9; i++ {
		fmt.Printf("%d ", game.board[i])
		if (i+1)%3 == 0 {
			fmt.Printf("\n")
		}
	}
}

func (game *GameBoard) tie() bool {
	for i:=0; i<9; i++ {
		if game.board[i] == 0 {
			return false
		}
	}
	return true
}

func (game *GameBoard) calScore() (int) {
	wins := [8][3]int{{0,1,2},{3,4,5},{6,7,8},{0,3,6},{1,4,7},{2,5,8},{0,4,8},{2,4,6}}
	for i:=0; i<8; i++ {
		if (game.board[wins[i][0]] != 0 &&
		   (game.board[wins[i][0]] == game.board[wins[i][1]]) &&
		   (game.board[wins[i][0]] == game.board[wins[i][2]]) {
				return game.board[wins[i][0]]
		   }
	}
	return 0
}

func (game *GameBoard) isGameDone() int {
	if game.tie() {
		fmt.Println("tie")
		return 10
	}

	if game.calScore() != 0 {
		return game.calScore()
	}

	return 0
}


func minimax(game GameBoard) int {
	var v int
	if game.tie() {
		return 0
	}
	if game.calScore() != 0 {
		return game.calScore()
	}
	// simulate move
	// max
	v = exs


	for i:=0; i<9; i++ {
		if game.board[i][j] == 0 {
			// Init the trial
			game.board[i][j] = copy.state
			game.state = -copy.state

			copy.printBoard()
			fmt.Println("try")
			v = max(v, minimax(game))

			// Reset the trial
			game.board[i][j] = 0
			game.state = -game.state
		}
	}
	fmt.Println("minimax: ", v)


	return v
}

func findMoveviaScore(game GameBoard, score int) (int, int) {
	v := exs
	for i:=0; i<9; i++ {
		if game.board[i][j] == 0 {
			// Init the trial
			game.board[i][j] = copy.state
			game.state = -copy.state

			copy.printBoard()
			fmt.Println("try")
			v = max(v, minimax(game))

			// Reset the trial
			game.board[i][j] = 0
			game.state = -game.state

			if v == score {
				fmt.Println("Best Score: ", v)
				return i, j
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

