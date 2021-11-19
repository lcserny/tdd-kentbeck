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

type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount: amount}
}

func (f Franc) Times(number int) Franc {
	return NewFranc(f.amount * number)
}

func (f Franc) Amount() int {
	return f.amount
}
