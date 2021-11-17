package tdd_kentbeck

type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d Dollar) Times(number int) Dollar {
	return NewDollar(d.amount * number)
}

func (d Dollar) Amount() int {
	return d.amount
}
