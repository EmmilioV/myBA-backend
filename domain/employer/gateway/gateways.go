package gateway

type Gateways struct {
	IDBProvider IDBProvider
}

func NewGateways(
	iDBProvider IDBProvider,
) *Gateways {
	return &Gateways{
		iDBProvider,
	}
}
