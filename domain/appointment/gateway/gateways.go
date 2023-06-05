package gateway

type Gateways struct {
	IDBInserter
	IDBProvider
}

func NewGateways(
	iDBInserter IDBInserter,
	iDBProvider IDBProvider,
) *Gateways {
	return &Gateways{
		IDBInserter: iDBInserter,
		IDBProvider: iDBProvider,
	}
}
