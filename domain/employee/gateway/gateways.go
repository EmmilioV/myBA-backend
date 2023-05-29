package gateway

type Gateways struct {
	IDBInserter
	IDBDeleter
}

func NewGateways(
	iDBInserter IDBInserter,
	iDBDeleter IDBDeleter,
) *Gateways {
	return &Gateways{
		iDBInserter,
		iDBDeleter,
	}
}
