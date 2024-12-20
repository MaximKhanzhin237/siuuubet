package postgresql

import (
	"database/sql"
	"sync"
)

var mut sync.Mutex

// метод для открытия базы данных
func CreateBD() *sql.DB {
	connStr := "user=postgres password=Superman2024$$ dbname=siuuubet sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

// метод для проверки наличия данных о ставках в базе данных
func CheckStavki() (int, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt1 := "select count(*) from Stavki"
	result1 := db.QueryRow(stmt1) //проверка: есть ли запись о балансе в таблице Polzovatel
	var n int
	er := result1.Scan(&n)
	if er != nil {
		mut.Unlock()
		return 0, er
	}
	mut.Unlock()
	return n, nil
}

// вставка информации о ставке в базу данных
func InsertStavki(Result string, BetSum float64, PotentialSum float64) error {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "insert into Stavki(result, bet_sum, potential_sum) values($1, $2, $3)"
	_, err := db.Exec(stmt, Result, BetSum, PotentialSum)
	if err != nil {
		mut.Unlock()
		return err
	}
	mut.Unlock()
	return nil
}

// проверка наличия информации о балансе пользователя
func CheckPolzovatel() (int, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt1 := "select count(*) from Polzovatel"
	result1 := db.QueryRow(stmt1) //проверка: есть ли запись о балансе в таблице Polzovatel
	var n int
	er := result1.Scan(&n)
	if er != nil {
		mut.Unlock()
		return 0, er
	}
	mut.Unlock()
	return n, nil
}

// обновление баланса пользователя в базе данных
func UpdatePolzovatel(balance float64) error {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "update Polzovatel set balance=$1" //если в Polzovatel есть запись о балансе, то значение баланса обновляется
	_, err := db.Exec(stmt, balance)
	if err != nil {
		mut.Unlock()
		return err
	}
	mut.Unlock()
	return nil
}

// вставка баланса пользователя в базу данных
func InsertPolzovatel(balance float64) error {
	db := CreateBD()
	defer db.Close()
	stmt2 := "insert into Polzovatel(balance) values($1)" //если в Polzovatel нет записи о балансе, то значение баланса вставляется
	_, err := db.Exec(stmt2, balance)
	if err != nil {
		mut.Unlock()
		return err
	}
	return nil
}

// извлечение информации о балансе пользователя из базы данных
func GetPolzovatel() (float64, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stm := "select * from Polzovatel"
	var num1 float64
	e := db.QueryRow(stm).Scan(&num1)
	if e != nil {
		mut.Unlock()
		return 0, e
	}
	mut.Unlock()
	return num1, nil
}

func GetStavki() (string, error) {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stm := "select * from Stavki" //если в Stavki есть данные о ставках, то эти данные записываются в выходную структуру ListOfBets
	rows, err := db.Query(stm)
	if err != nil {
		mut.Unlock()
		return "", err
	}
	r := ""
	for rows.Next() {
		var s string
		var s1 string
		var s2 string
		var s3 string
		if err := rows.Scan(&s, &s1, &s2, &s3); err != nil {
			mut.Unlock()
			return "", err
		}
		s = s + " " + s1 + " " + s2 + " " + s3
		r += s + "\n"
	}
	mut.Unlock()
	return r, nil
}

func DeleteStavki(id int) error {
	mut.Lock()
	db := CreateBD()
	defer db.Close()
	stmt := "delete from Stavki where id=$1" //удаление ставки из Stavki по id
	db.QueryRow(stmt, id)
	mut.Unlock()
	return nil
}
