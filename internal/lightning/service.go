package lightning

type LightningInterface interface {
	NodeInfo() (*NodeState, error)
	ReceivePayment(amountSats int64, description string) (*Invoice, <-chan Payment, error)
	SendPayment(bolt11 string, amountSats *int64) (<-chan Payment, error)
}

