package Strategy

import "awesomeProject2/pkg/models/postgresql"

type BetMod interface {
	InsertStavki(string, float64, float64) error
	CheckStavki() (int, error)
	GetStavki() (string, error)
	DeleteStavki(int) error
	CheckPolzovatel() (int, error)
	GetPolzovatel() (float64, error)
	UpdatePolzovatel(float64) error
	InsertPolzovatel(float64) error
}

// создание пользовательского типа - это первый экземпляр интерфейса
type BetM struct{}

func (b BetM) InsertStavki(r string, s float64, p float64) error {
	return postgresql.InsertStavki(r, s, p)
}

func (b BetM) CheckStavki() (int, error) {
	return postgresql.CheckStavki()
}

func (b BetM) GetStavki() (string, error) {
	return postgresql.GetStavki()
}

func (b BetM) DeleteStavki(id int) error {
	return postgresql.DeleteStavki(id)
}

func (b BetM) CheckPolzovatel() (int, error) {
	return postgresql.CheckPolzovatel()
}

func (b BetM) GetPolzovatel() (float64, error) {
	return postgresql.GetPolzovatel()
}

func (b BetM) UpdatePolzovatel(p float64) error {
	return postgresql.UpdatePolzovatel(p)
}

func (b BetM) InsertPolzovatel(s float64) error {
	return postgresql.InsertPolzovatel(s)
}

// создание пользовательского типа - это второй экземпляр интерфейса
type CheckMock struct{}

var CheckStavkiM func() (int, error)
var CheckPolzovatelM func() (int, error)

var GetPolzovatelM func() (float64, error)

var GetStavkiM func() (string, error)

var InsertStavkiM func(r string, b float64, p float64) error

var DeleteStavkiM func(id int) error

var InsertPolzovatelM func(b float64) error

var UpdatePolzovatelM func(b float64) error

func (m CheckMock) CheckStavki() (int, error) {
	return CheckStavkiM()
}

func (m CheckMock) CheckPolzovatel() (int, error) {
	return CheckPolzovatelM()
}

func (m CheckMock) GetPolzovatel() (float64, error) {
	return GetPolzovatelM()
}

func (m CheckMock) GetStavki() (string, error) {
	return GetStavkiM()
}

func (m CheckMock) InsertStavki(r string, b float64, p float64) error {
	return InsertStavkiM(r, b, p)
}

func (m CheckMock) DeleteStavki(id int) error {
	return DeleteStavkiM(id)
}

func (m CheckMock) InsertPolzovatel(b float64) error {
	return InsertPolzovatelM(b)
}

func (m CheckMock) UpdatePolzovatel(b float64) error {
	return UpdatePolzovatelM(b)
}

type Context struct {
	Strategy BetMod
}

func (c *Context) Algorithm(a BetMod) {
	c.Strategy = a
}

func (c *Context) CheckStavki() (int, error) {
	return c.Strategy.CheckStavki()
}

func (c *Context) CheckPolzovatel() (int, error) {
	return c.Strategy.CheckPolzovatel()
}

func (c *Context) GetPolzovatel() (float64, error) {
	return c.Strategy.GetPolzovatel()
}

func (c *Context) GetStavki() (string, error) {
	return c.Strategy.GetStavki()
}

func (c *Context) InsertStavki(r string, b float64, p float64) error {
	return c.Strategy.InsertStavki(r, b, p)
}

func (c *Context) DeleteStavki(id int) error {
	return c.Strategy.DeleteStavki(id)
}

func (c *Context) InsertPolzovatel(b float64) error {
	return c.Strategy.InsertPolzovatel(b)
}

func (c *Context) UpdatePolzovatel(b float64) error {
	return c.Strategy.UpdatePolzovatel(b)
}
