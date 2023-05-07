package lightning

type Invoice struct {
	Bolt11          string
	PayeePubkey     string
	PaymentHash     string
	Description     *string
	DescriptionHash *string
	AmountMsat      *int64
	Timestamp       int64
	Expiry          int64
	PaymentSecret   []byte
}

type NodeState struct {
	ID                string
	BlockHeight       uint32
	LocalBalanceMsat  uint64
	RemoteBalanceMsat uint64
	OnchainBalance    int64
}

type Payment struct {
	ID          string
	AmountMsat  int64
	FeeMsat     int64
	Pending     bool
	Settled     bool
	Timestamp   int64
	Description *string
}
