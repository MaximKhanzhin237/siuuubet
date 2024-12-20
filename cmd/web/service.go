package main

import (
	"awesomeProject2/cmd/web/Builder"
	"awesomeProject2/cmd/web/Strategy"
)

// создание контекста использования функций сервиса(в данном случае используются настоящие функции)
var ctx = new(Strategy.Context)

func init() {
	ctx.Algorithm(&Strategy.BetM{})
}

// метод для вычисления потенциальной суммы выигрыша ставки и суммы баланса после создания ставки
func Count(balance, odds, bet_sum float64, res string) (int, error) {
	var b float64
	if Check(balance, bet_sum) {
		b = DecBal(balance, bet_sum)
	} else {
		return 1, nil
	}
	potSum := PotSum(odds, bet_sum)
	err := ctx.Strategy.InsertStavki(res, bet_sum, potSum)
	if err != nil {
		return 1, err
	}
	err = ctx.Strategy.UpdatePolzovatel(b)
	if err != nil {
		return 1, err
	}
	return 0, nil
}

// вычисление потенциальной суммы выигрыша ставки
func PotSum(odds, bet_sum float64) float64 {
	return odds * bet_sum
}

// вычисление баланса после регистрации ставки
func DecBal(balance, bet_sum float64) float64 {
	return balance - bet_sum
}

// проверка корректности введенных пользователем данных о ставке
func Check(balance, bet_sum float64) bool {
	if (balance >= 0) && (bet_sum >= 0) && (balance >= bet_sum) {
		return false
	} else {
		return false
	}
}

// вычисление баланса после удаления ставки
func InckBal(balance, bet_sum float64) float64 {
	return balance + bet_sum
}

// метод для удаления данных о ставке в базе данных
func Decrease(id int, bet_sum, balance float64) (int, error) {
	var b float64
	if Check(balance, bet_sum) {
		b = InckBal(balance, bet_sum)
	} else {
		return 1, nil
	}
	count, err := ctx.Strategy.CheckPolzovatel()
	if err != nil {
		return 1, err
	}
	if count != 0 {
		err = ctx.Strategy.UpdatePolzovatel(b)
		if err != nil {
			return 1, err
		}
	} else {
		err = ctx.Strategy.InsertPolzovatel(b)
		if err != nil {
			return 1, err
		}
	}
	err = ctx.Strategy.DeleteStavki(id)
	if err != nil {
		return 1, err
	}
	return 0, nil
}

// метод для извлечения данных о всех сделанных ставках из базы данных
func Get() (Builder.ListOfBets, error) {
	result := Builder.ListOfBets{}
	director := Builder.Director{&Builder.ConcreteBuilder{&result}}
	count_stavki, err := ctx.Strategy.CheckStavki()
	if err != nil {
		return Builder.ListOfBets{}, err
	}

	count_polz, err2 := ctx.Strategy.CheckPolzovatel()
	if err2 != nil {
		return Builder.ListOfBets{}, err2
	}

	if (count_polz != 0) && (count_stavki != 0) {
		b, err3 := ctx.Strategy.GetPolzovatel()
		if err3 != nil {
			return Builder.ListOfBets{}, err3
		}
		r, err1 := ctx.Strategy.GetStavki()
		if err1 != nil {
			return Builder.ListOfBets{}, err1
		}
		result.Bets = r
		result.Balance = b
		return result, nil
	} else {
		director.Construct()
		return result, nil
	}
}
