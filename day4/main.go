package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type BingoGame struct {
	Instructions []int
	Boards       [][5][5]interface{}
}

func getUnmarkedTotal(game *BingoGame, winnerBoardIdx int) int {
	var unmarkedTotal int

	for _, row := range game.Boards[winnerBoardIdx] {
		for _, item := range row {
			if item != "x" {
				// log.Println("adding", item, "to total")
				unmarkedTotal = unmarkedTotal + item.(int)
				// log.Println("total", unmarkedTotal)
			}
		}
	}

	return unmarkedTotal
}

func checkForWinner(game *BingoGame) (bool, int) {

	var rowCheck, colCheck bool

	for boardIdx, board := range game.Boards {

		rowCheck = true
		for _, row := range board {
			// Check all items in the row, see if there is a row win.
			for _, item := range row {
				if item != "x" {
					rowCheck = false
				}
			}

			if rowCheck {
				// log.Println("Row win found!", game.Boards)
				return true, boardIdx
			}
		}

		colCheck = true
		for itemIdx := 0; itemIdx < 5; itemIdx++ {
			for _, row := range board {

				if row[itemIdx] != "x" {
					colCheck = false
				}
			}

			if colCheck {
				// log.Println("Column win found!", game.Boards[boardIdx])
				return true, boardIdx
			}
		}
	}

	return false, 0
}

func playGame(game *BingoGame) (int, int) {
	var winningNumber, unmarkedTotal int

	for _, instruction := range game.Instructions {
		for boardIdx, board := range game.Boards {
			for rowIdx, row := range board {
				for itemIdx, item := range row {
					if instruction == item {

						// Mark this as a hit.
						game.Boards[boardIdx][rowIdx][itemIdx] = "x"

						isWinner, winnerBoardIdx := checkForWinner(game)

						// Check for a winner relative to last placed x.
						if isWinner {
							winningNumber = instruction
							unmarkedTotal = getUnmarkedTotal(game, winnerBoardIdx)
							return winningNumber, unmarkedTotal
						}

					}
				}
			}
		}
	}

	return winningNumber, unmarkedTotal
}

func part1(game *BingoGame) {
	winningNumber, unmarkedTotal := playGame(game)

	log.Println("Winning Number:", winningNumber)
	log.Println("Unmarked Total:", unmarkedTotal)
	log.Println("Part2:", unmarkedTotal*winningNumber)

	return
}

func part2() {

}

func resetBoard(board [5][5]interface{}) [5][5]interface{} {
	var i, j int

	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			board[i][j] = 0
		}
	}

	return board
}

func processInput() (*BingoGame, error) {
	game := &BingoGame{}
	filePath := "day4/input"
	absPath, _ := filepath.Abs("../" + filePath)
	f, err := os.Open(absPath)

	if err != nil {
		log.Panic("Error reading input!: ", err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Read first line to instructions
	scanner.Scan()
	instructionStrings := strings.Split(scanner.Text(), ",")

	for _, instruction := range instructionStrings {
		bingoNumber, _ := strconv.Atoi(instruction)
		game.Instructions = append(game.Instructions, bingoNumber)
	}

	// Keep track of our index.
	lineNumber := 0

	// Temporary board for use in creating boards.
	var board [5][5]interface{}

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			// log.Println("Resetting board.")
			board = resetBoard(board)
			// log.Println(board)
			continue
		} else {
			rowIdx := lineNumber % 5
			// log.Println("Populating row", rowIdx)
			// Not a command, but groups of boards.
			// log.Println("Line", lineNumber, "\t\tRow Index", rowIdx, "\t\tText", text)

			items := strings.Fields(scanner.Text())
			for colIdx, bingoItem := range items {

				bingoItemNumber, _ := strconv.Atoi(bingoItem)
				board[rowIdx][colIdx] = bingoItemNumber
			}
		}

		// log.Println(lineNumber, board)

		if (lineNumber+1)%5 == 0 {
			game.Boards = append(game.Boards, board)
			// log.Println(game.Boards)
		}

		lineNumber++
	}

	return game, nil
}

func main() {

	game, _ := processInput()

	part1(game)
}
