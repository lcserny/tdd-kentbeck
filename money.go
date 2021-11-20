package tdd_kentbeck

type Money interface {
	Amount() int
	Times(multiplier int) Money
}

type Dollar struct {
	amount int
}

func NewDollar(amount int) Money {
	return Dollar{amount: amount}
}

func (d Dollar) Times(multiplier int) Money {
	return NewDollar(d.amount * multiplier)
}

func (d Dollar) Amount() int {
	return d.amount
}

type Franc struct {
	amount int
}

func NewFranc(amount int) Money {
	return Franc{amount: amount}
}

func (f Franc) Times(multiplier int) Money {
	return NewFranc(f.amount * multiplier)
}

func (f Franc) Amount() int {
	return f.amount
}
