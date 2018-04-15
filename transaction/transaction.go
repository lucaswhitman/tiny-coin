package transaction

type Tranaction struct {
	To     string `json:"to"`
	From   string `json:"from"`
	Amount int    `json:"amount"`
}
