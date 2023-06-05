package gateway

type Gateways struct {
	IDBProvider IDBProvider
	IDBInserter IDBInserter
	IDBUpdater  IDBUpdater
}

func NewGateways(
	iDBProvider IDBProvider,
	iDBInserter IDBInserter,
	iDBUpdater IDBUpdater,
) *Gateways {
	return &Gateways{
		IDBProvider: iDBProvider,
		IDBInserter: iDBInserter,
		IDBUpdater:  iDBUpdater,
	}
}
