package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var row, col int

	continuePlaying := true
	fmt.Println("")
	fmt.Println("=====================================================================================")
	fmt.Println("Let's Play Tic Tac Toe")
	fmt.Println("=====================================================================================")
	for continuePlaying {
		boardSize := getUserInput("Set the board size (min:3, max:10)", validateBoardSize).(int)
		numPlayers := getUserInput("Choose number of players. (1 or 2)", validateNumPlayers).(int)
		p1Piece := getUserInput("Player 1, select a piece: (X or O)", validatePiece).(string)
		p2Piece := (map[string]string{"X": "O", "O": "X"}[p1Piece])

		hasWinner := false
		winner := ""

		// get first to move
		firstToMove := (map[int]string{1: "X", 0: "O"}[rand.Intn(2)])
		board := InitializeBoard(boardSize, firstToMove)
		fmt.Println("=====================================================================================")
		fmt.Println("Starting game...")
		fmt.Println("To place a piece, enter your move as an input: row,column (e.g. 1,1 for the top left cell)")
		fmt.Printf("%s is first to move.\n", firstToMove)
		PrintBoard(board)
		fmt.Println("=====================================================================================")

		for !hasWinner && board.turnNumber < boardSize*boardSize {
			// get input from AI
			if numPlayers == 1 && board.turn == p2Piece {
				fmt.Printf("Computer's (%s) turn   \n", board.turn)
				row, col = getMoveFromAI(board)
				board.Add(col, row)
				// add delay
				time.Sleep(300 * time.Millisecond)

			} else {
				fmt.Printf("%s's turn. Move?   ", board.turn)
				userMove := getUserInput("", validateUserMove, board).([]int)
				row, col = userMove[0], userMove[1]
				// row, col = getMoveFromUser(board)
				board.Add(col, row)
			}

			PrintBoard(board)
			hasWinner, winner = CheckBoard(board.state)

			if hasWinner || board.turnNumber == boardSize*boardSize {
				if hasWinner {
					fmt.Printf("Game over. %s won\n", winner)
				} else if board.turnNumber == boardSize*boardSize {
					fmt.Println("Game over. It's a draw")
				}
				endGameResponse := getUserInput("Would you want to play again? restart/new/end", validateEndGameResponse).(string)

				if endGameResponse == "restart" {
					continuePlaying = true
					hasWinner = false
					winner = ""
					board = InitializeBoard(boardSize, (map[int]string{1: "X", 0: "O"}[rand.Intn(2)]))
					fmt.Println("Restarting game...")
					fmt.Println("===========================================================================")
					PrintBoard(board)
				} else if endGameResponse == "new" {
					continuePlaying = true
					// but do not reset the board => forces a new game
				} else if endGameResponse == "end" {
					continuePlaying = false
				}
			}

		} // end of board
	}
	fmt.Println("Thank you for playing Tic-Tac-Toe!")
}

// get user inputs
func getUserInput(msg string, validator func(string, ...interface{}) (interface{}, bool, string), params ...interface{}) interface{} {
	var isValid bool
	var res string
	var validatedInput interface{}
	for !isValid {
		fmt.Printf("%s  \n", msg)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fromStdIn := scanner.Text()
		validatedInput, isValid, res = validator(fromStdIn, params...)
		fmt.Printf("%s  \n", res)
	}
	return validatedInput
}

// validators
// param: input
// param: params (other params -- ordered in a slice of interfaces)
// returns interface, bool, string
func validateBoardSize(input string, params ...interface{}) (interface{}, bool, string) {
	boardSize := 0
	isValid := false
	res := ""
	_, err := fmt.Sscanf(input, "%d", &boardSize)
	if err != nil || !(boardSize >= 3 && boardSize <= 10) {
		res = "Input invalid. Please try again."
		boardSize = 0
	} else {
		isValid = true
		res = ""
	}
	var output interface{} = boardSize
	return output, isValid, res
}

func validatePiece(input string, params ...interface{}) (interface{}, bool, string) {
	p1Piece := ""
	isValid := false
	res := ""
	_, err := fmt.Sscanf(input, "%s", &p1Piece)
	if err != nil || !(p1Piece == "X" || p1Piece == "O") {
		res = "Input invalid. Please try again."
		p1Piece = ""
	} else {
		isValid = true
		res = ""
	}
	var output interface{} = p1Piece
	return output, isValid, res
}

func validateNumPlayers(input string, params ...interface{}) (interface{}, bool, string) {
	numPlayers := 0
	isValid := false
	res := ""
	_, err := fmt.Sscanf(input, "%d", &numPlayers)
	if err != nil || !(numPlayers == 1 || numPlayers == 2) {
		res = "Input invalid. Please try again."
		numPlayers = 0
	} else {
		isValid = true
		res = ""
	}
	var output interface{} = numPlayers
	return output, isValid, res
}

func validateEndGameResponse(input string, params ...interface{}) (interface{}, bool, string) {
	endGameResponse := ""
	isValid := false
	res := ""
	_, err := fmt.Sscanf(input, "%s", &endGameResponse)
	if err != nil || !(endGameResponse == "restart" || endGameResponse == "new" || endGameResponse == "end") {
		res = "Input invalid. Please try again."
		endGameResponse = ""
	} else {
		isValid = true
		res = ""
	}
	var output interface{} = endGameResponse
	return output, isValid, res
}

func isValidMove(x int, y int, boardSize int) bool {
	return y <= boardSize && x <= boardSize && y >= 1 && x >= 1
}

func validateUserMove(input string, params ...interface{}) (interface{}, bool, string) {
	row, col := 0, 0
	isValid := false
	var output interface{} = []int{0, 0}
	res := ""
	if len(params) == 0 {
		// this bug shouldn't be exposed since call on validateUserMove
		// needs board Board as part of the param (and as a first argument)
		return output, isValid, "[BUG IN CODE]"

	} else if board, ok := params[0].(Board); ok {
		_, err := fmt.Sscanf(input, "%d,%d", &row, &col)
		if err != nil || !(isValidMove(col, row, len(board.state))) {
			res = "Input invalid. Please try again."
		} else {
			isValid = true
			res = ""
		}
		var output interface{} = []int{row, col}
		return output, isValid, res
	} else {
		return output, isValid, "[BUG IN CODE]"
	}
}

// noob AI - picks a random move from the remaining list of all valid moves
func getMoveFromAI(board Board) (int, int) {
	var row, col int
	pick := rand.Intn(len(board.validMoves))
	ct := 0
	for k, _ := range board.validMoves {
		if ct == pick {
			// move_ := strings.Fields(move)
			fmt.Sscanf(k, "%d,%d", &row, &col)
		}
		ct += 1
	}
	return row, col
}

// game related functions

func InitializeBoard(boardSize int, startingPlayer string) Board {
	board := make([][]int, boardSize)
	for i, _ := range board {
		board[i] = make([]int, boardSize)
	}
	validMoves := make(map[string]int8)
	for row, _ := range board {
		for col, _ := range board {
			validMoves[fmt.Sprintf("%d,%d", row+1, col+1)] = 0
		}
	}
	game := Board{board, startingPlayer, 0, validMoves}
	return game
}

func checkWinner(sum int, boardSize int) string {
	switch true {
	case (sum == boardSize):
		return "X"
	case (sum == -boardSize):
		return "O"
	default:
		return ""
	}
}

func PrintBoard(board Board) {
	// X 1, O -1
	for _, xs := range board.state {
		for i, x := range xs {
			stringRep := (map[int]string{1: "X", -1: "O", 0: "_"})[x]
			if i < (len(xs) - 1) {
				fmt.Printf("%s ", stringRep)
			} else {
				fmt.Printf("%s\n", stringRep)
			}
		}
	}
}

func CheckBoard(board [][]int) (bool, string) {
	var winner string
	var sum int
	boardSize := len(board)
	// check rows
	for y := range board {
		sum = 0
		for x := range board[0] {
			sum += board[y][x]
		}
		winner = checkWinner(sum, boardSize)
		if winner != "" {
			return true, winner
		}
	}
	// check cols
	for x := range board[0] {
		sum = 0
		for y := range board {
			sum += board[y][x]
		}
		winner = checkWinner(sum, boardSize)
		if winner != "" {
			return true, winner
		}
	}
	// check diagonal from top left
	sum = 0
	for y := range board {
		sum += board[y][y]
	}
	winner = checkWinner(sum, boardSize)
	if winner != "" {
		return true, winner
	}
	// check diagonal from bottom left
	sum = 0
	for y := range board {
		sum += board[boardSize-y-1][y]
	}
	winner = checkWinner(sum, boardSize)
	if winner != "" {
		return true, winner
	}
	return false, winner
}

type Board struct {
	state      [][]int
	turn       string
	turnNumber int
	validMoves map[string]int8
}

func (board *Board) Add(x int, y int) (int, error) {
	y_zeroind := y - 1
	x_zeroind := x - 1
	if board.state[y_zeroind][x_zeroind] != 0 {
		return 1, errors.New("Spot already has a piece. Select a different spot")
	} else {
		board.state[y_zeroind][x_zeroind] = (map[string]int{"X": 1, "O": -1})[board.turn]
		board.turn = (map[string]string{"X": "O", "O": "X"})[board.turn]
		board.turnNumber += 1
		delete(board.validMoves, fmt.Sprintf("%d,%d", y, x))
		return 0, nil
	}
}

// func getPiece() string {
// 	p1Piece := ""
// 	for p1Piece == "" {
// 		fmt.Println("Player 1. Select a piece: (X or O)")
// 		_, err := fmt.Scanf("%s", &p1Piece)
// 		if err != nil || !(p1Piece == "X" || p1Piece == "O") {
// 			fmt.Println("Input invalid. Please try again.")
// 			p1Piece = ""
// 		}
// 	}
// 	return p1Piece
// }
// func getNumPlayers() int {
// 	numPlayers := 0
// 	for numPlayers == 0 {
// 		fmt.Println("Game mode. Input 1 for single player or 2 single player.")
// 		_, err := fmt.Scanf("%d", &numPlayers)
// 		if err != nil || !(numPlayers == 1 || numPlayers == 2) {
// 			fmt.Println("Input invalid.Input 1 for single player or 2 single player.")
// 			numPlayers = 0
// 		} else {
// 			fmt.Printf("'%s' selected \n", (map[int]string{1: "1P mode", 2: "Versus mode"}[numPlayers]))
// 		}
// 	}
// 	return numPlayers
// }
// func getEndGameResponse() string {
// 	endGameResponse := ""
// 	for endGameResponse == "" {
// 		fmt.Println("Would you want to play again? restart/new/end")
// 		_, err := fmt.Scanf("%s", &endGameResponse)
// 		if err != nil || !(endGameResponse == "restart" || endGameResponse == "new" || endGameResponse == "end") {
// 			fmt.Println("Please try again. Would you want to play again? restart/new/end")
// 			endGameResponse = ""
// 		}
// 	}
// 	return endGameResponse
// }
// func getMoveFromUser(board Board) (int, int) {
// 	var userInput string
// 	row, col := 0, 0
// 	_, err := fmt.Scanf("%s", &userInput)
// 	fmt.Println(userInput)
// 	for row == 0 || col == 0 {
// 		if err != nil {
// 			fmt.Println("Please try again. Move should be a valid string of format row,col")
// 			row, col = 0, 0
// 		} else {
// 			//check if valid input
// 			matched, _ := regexp.MatchString("[1-9]{1,2},[1-9]{1,2}", userInput)
// 			if matched {
// 				fmt.Sscanf(userInput, "%d,%d", &row, &col)
// 				if isValidMove(col, row, len(board.state)) {
// 					return row, col
// 				} else {
// 					fmt.Println("Please try again. Cell is beyond the board size")
// 					row, col = 0, 0
// 				}
// 			} else {
// 				fmt.Println("Please try again. Move should be a valid string of format row,col")
// 				row, col = 0, 0
// 			}
// 		}
// 	}
// 	return row, col
// }
