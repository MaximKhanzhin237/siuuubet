package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

// Структура для записи информации о ставке в базу данных
type Bet struct {
	Balance float64 `json:"balance"`
	Result  string  `json:"result"`
	BetSum  float64 `json:"bet_sum"`
	Odds    float64 `json:"odds"`
}

// Структура для отправки имеющихся ставок на интерфейс сайта

// Структура для удаления ставки из базы данных
type Bet_del struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	BetSum  float64 `json:"bet_sum"`
}
