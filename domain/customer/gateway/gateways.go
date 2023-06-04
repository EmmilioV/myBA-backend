package gateway

type Gateways struct {
	IDBInserter
}

func NewGateways(
	iDBInserter IDBInserter,
) *Gateways {
	return &Gateways{
		IDBInserter: iDBInserter,
	}
}
