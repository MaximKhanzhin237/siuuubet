package Builder

type Builder interface {
	MakeBalance(float64)
	MakeBets(string)
}

type Director struct {
	Builder Builder
}

func (d *Director) Construct() {
	d.Builder.MakeBalance(1000)
	d.Builder.MakeBets("")
}

type ConcreteBuilder struct {
	Product *ListOfBets
}

// MakeHeader builds a header of document..
func (b *ConcreteBuilder) MakeBalance(num float64) {
	b.Product.Balance = num
}

// MakeBody builds a body of document.
func (b *ConcreteBuilder) MakeBets(str string) {
	b.Product.Bets = str
}

type ListOfBets struct {
	Check   int
	Balance float64
	Bets    string
}
