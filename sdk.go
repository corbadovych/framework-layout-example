package framework_layout_example

import "github.com/corbadovych/framework-layout-example/pkg/sdk"

type SDK interface {
	BuildClientInfo(remoteAddress string) (*sdk.ClientInfo, error)
}

func New() *Impl {
	return &Impl{}
}

type Impl struct {
}

func (i *Impl) BuildClientInfo(remoteAddress string) (*sdk.ClientInfo, error) {
	return sdk.NewClientInfo(remoteAddress), nil
}
