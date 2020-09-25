package types

//Money is type money
type Money int64

//Category of payments categories
type Category string

//Status payment status
type Status string

//Status categories
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment struct
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status Status
}
