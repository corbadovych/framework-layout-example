package sdk

type ClientInfo struct {
	RemoteAddress string
}

func NewClientInfo(remoteAddress string) *ClientInfo {
	return &ClientInfo{
		RemoteAddress: remoteAddress,
	}
}
