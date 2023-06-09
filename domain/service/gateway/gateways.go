package gateway

type Gateways struct {
	IDBProvider  IDBProvider
	IDBInserter  IDBInserter
	IDBUpdater   IDBUpdater
	IMQPublisher IMQPublisher
}

func NewGateways(
	iDBProvider IDBProvider,
	iDBInserter IDBInserter,
	iDBUpdater IDBUpdater,
	iMQPublisher IMQPublisher,
) *Gateways {
	return &Gateways{
		IDBProvider:  iDBProvider,
		IDBInserter:  iDBInserter,
		IDBUpdater:   iDBUpdater,
		IMQPublisher: iMQPublisher,
	}
}
