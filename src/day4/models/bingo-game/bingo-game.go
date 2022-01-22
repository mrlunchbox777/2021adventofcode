package bingogame

import (
	bb "github.com/mrlunchbox777/2021adventofcode/src/day4/models/bingo-board"
	wn "github.com/mrlunchbox777/2021adventofcode/src/day4/models/winning-numbers"
)

type BingoGame struct {
	bingoBoards   []bb.BingoBoard
	answers       wn.WinningNumbers
	winningBoards []bb.BingoBoard
	losingBoards  []bb.BingoBoard
	rounds        int
}

func (game BingoGame) BingoBoards() []bb.BingoBoard {
	return game.bingoBoards
}

func (game BingoGame) Answers() wn.WinningNumbers {
	return game.answers
}

func (game BingoGame) WinningBoards() []bb.BingoBoard {
	return game.winningBoards
}

func (game BingoGame) LosingBoards() []bb.BingoBoard {
	return game.losingBoards
}

func (game BingoGame) Rounds() int {
	return game.rounds
}

func CreateBingoGame(bingoBoards []bb.BingoBoard, winningNumbers wn.WinningNumbers) BingoGame {
	return BingoGame{bingoBoards: bingoBoards, answers: winningNumbers}
}
