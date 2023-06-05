package gateway

type Gateways struct {
	IDBInserter
	IDBProvider
	IDBDeleter
	IDBUpdater
}

func NewGateways(
	iDBInserter IDBInserter,
	iDBProvider IDBProvider,
	iDBDeleter IDBDeleter,
	iDBUpdater IDBUpdater,
) *Gateways {
	return &Gateways{
		IDBInserter: iDBInserter,
		IDBProvider: iDBProvider,
		IDBDeleter:  iDBDeleter,
		IDBUpdater:  iDBUpdater,
	}
}
