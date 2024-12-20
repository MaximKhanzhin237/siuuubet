package main

import (
	"awesomeProject2/cmd/web/Builder"
	"awesomeProject2/pkg/models"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"sync"
)

// реализация паттерна Singleton
var (
	res  Builder.ListOfBets
	once sync.Once
)

func GetInstance() {
	once.Do(func() {
		res = Builder.ListOfBets{}
	})
}

var flag int
var mut sync.RWMutex

func init() {
	GetInstance()
}

// обработчик для пути '/'
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	mut.RLock()
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		mut.RUnlock()
		return
	}

	//если запрос с методом GET
	if r.Method == "GET" {
		ts, err := template.ParseFiles("./ui/html/home.page.html")

		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			mut.RUnlock()
			return
		}
		res, err = Get() //вызов метода уровня Service для извлечения актуальных данных о ставках и балансе пользователя
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			mut.RUnlock()
			return
		}
		res.Check = flag

		//отправка данных о ставках и балансе на frontend
		err = ts.Execute(w, res)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
			mut.RUnlock()
			return
		}

		//если запрос с методом POST(создание новой ставки)
	} else if r.Method == "POST" {

		var bet models.Bet
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&bet)
		if err != nil {
			mut.RUnlock()
			http.Error(w, "Internal Server Error", 500)
			return
		}
		flag, err = Count(bet.Balance, bet.Odds, bet.BetSum, bet.Result) //вызов метода уровня service для вычисления баланса
		// и потенциальной суммы выигрыша ставки
		if err != nil {
			mut.RUnlock()
			http.Error(w, "Internal Server Error", 500)
			return
		}

		//если запрос с методом DELETE(продажа ставки)
	} else if r.Method == "DELETE" {
		var bet models.Bet_del
		decoder := json.NewDecoder(r.Body)
		er := decoder.Decode(&bet)
		if er != nil {
			mut.RUnlock()
			http.Error(w, "Internal Server Error", 500)
			return
		}
		flag, er = Decrease(bet.ID, bet.BetSum, bet.Balance) //вызов метода уровня service для вычисления суммы баланса после удаления ставки
		if er != nil {
			mut.RUnlock()
			http.Error(w, "Internal Server Error", 500)
			return
		}

	} else {
		mut.RUnlock()
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	mut.RUnlock()
}
