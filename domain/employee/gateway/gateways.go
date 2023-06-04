package gateway

type Gateways struct {
	IDBInserter
	IDBDeleter
	IDBUpdater
}

func NewGateways(
	iDBInserter IDBInserter,
	iDBDeleter IDBDeleter,
	iDBUpdater IDBUpdater,
) *Gateways {
	return &Gateways{
		IDBInserter: iDBInserter,
		IDBDeleter:  iDBDeleter,
		IDBUpdater:  iDBUpdater,
	}
}
