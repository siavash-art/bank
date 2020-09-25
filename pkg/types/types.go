package types

//Money is type money
type Money int64

//Category of payments categories
type Category string

//Payment struct
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
