package currency

type Country struct {
	Name string
}

type Currency struct {
	Country
	Name string
}

type Money struct {
	Currency
	Amount int64
}
