package main

import (
	"awesomeProject2/cmd/web/Builder"
	"awesomeProject2/cmd/web/Strategy"
	_ "github.com/lib/pq"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

// создание контекста использования функций сервиса(в данном случае используются моки функций сервиса)
func init() {
	ctx.Algorithm(&Strategy.CheckMock{})
}

func TestPotSum(t *testing.T) {
	t.Log("Тестирование PotSum(считает потенциальную сумму выигрыша) вход: odds: 2.1, bet_sum: 4")
	var odds = 2.1
	var bet_sum float64 = 4
	res1 := PotSum(odds, bet_sum)
	if 8.4 == res1 {
		t.Logf("\t%s\tТест PotSum вернул верные данные: 2.1*4=8.4", success)
	} else {
		t.Fatalf("\t%s\tТест PotSum провалился: 8.4 не равно %f", failed, res1)
	}
}

func TestDecBal(t *testing.T) {
	t.Log("Тестирование DecBal(считает разницу баланса и суммы ставки) вход: balance: 100, bet_sum: 15")
	var balance float64 = 100
	var bet_sum float64 = 15
	res1 := DecBal(balance, bet_sum)
	if float64(85) == res1 {
		t.Logf("\t%s\tТест DecBal вернул верные данные: 100-15=85", success)
	} else {
		t.Fatalf("\t%s\tТест DecBal провалился: 85 не равно %f", failed, res1)
	}
}

func TestCheck(t *testing.T) {
	//корректные данные
	t.Log("Тестирование Check(проверяет корректность значений баланса и коэффициентов) вход: balance: 100, bet_sum: 15")
	var balance float64 = 100
	var bet_sum float64 = 15
	res1 := Check(balance, bet_sum)
	if true == res1 {
		t.Logf("\t%s\tТест Check вернул верные данные: true", success)
	} else {
		t.Fatalf("\t%s\tТест Check провалился: true не равно false", failed)
	}

	//некорректные данные
	bet_sum = 200
	res2 := Check(balance, bet_sum)
	if false == res2 {
		t.Logf("\t%s\tТест Check вернул верные данные: false", success)
	} else {
		t.Fatalf("\t%s\tТест Check провалился: false не равно true", failed)
	}
}

func TestInckBal(t *testing.T) {
	t.Log("Тестирование InckBal(вычисляет сумму баланса и ставки) вход: balance: 100, bet_sum: 15")
	var balance float64 = 100
	var bet_sum float64 = 15
	res1 := InckBal(balance, bet_sum)
	if float64(115) == res1 {
		t.Logf("\t%s\tТест InckBal вернул верные данные: 100+15=115", success)
	} else {
		t.Fatalf("\t%s\tТест InckBal провалился: 115 не равно %f", failed, res1)
	}
}

func TestGet(t *testing.T) {
	t.Log("Тестирование Get(возвращает в виде структуры информацию о ставках) ставки из базы данных: \"74 Odds (Arsenal win): 2.1 (Odds: 2.1) 100 210\\n75 Odds (Liverpool win): 2.8 (Odds: 2.8) 100 280\\n\"")

	//инициализация второго экземпляра интерфейса betMod
	Strategy.CheckStavkiM = func() (int, error) {
		return 1, nil
	}
	Strategy.CheckPolzovatelM = func() (int, error) {
		return 1, nil
	}
	Strategy.GetPolzovatelM = func() (float64, error) {
		return 1000, nil
	}
	Strategy.GetStavkiM = func() (string, error) {
		return "74 Odds (Arsenal win): 2.1 (Odds: 2.1) 100 210\n75 Odds (Liverpool win): 2.8 (Odds: 2.8) 100 280\n", nil
	}
	bet := Builder.ListOfBets{Bets: "74 Odds (Arsenal win): 2.1 (Odds: 2.1) 100 210\n75 Odds (Liverpool win): 2.8 (Odds: 2.8) 100 280\n",
		Balance: 1000, Check: 0}
	bets, err := Get()
	if err != nil {
		t.Fatalf("\t%s\tТест Get провалился из-за ошибки: %s", failed, err.Error())
	}
	if bet == bets {
		t.Logf("\t%s\tТест Get вернул верные данные", success)
	} else {
		t.Fatalf("\t%s\tТест Get провалился", failed)
	}
}
