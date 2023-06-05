package gateway

type Gateways struct {
	IDBProvider IDBProvider
	IDBInserter IDBInserter
}

func NewGateways(
	iDBProvider IDBProvider,
	iDBInserter IDBInserter,
) *Gateways {
	return &Gateways{
		IDBProvider: iDBProvider,
		IDBInserter: iDBInserter,
	}
}
