package app

import "github.com/vitalvas/gomarkov/libmarkov"

// Model godoc
type Model struct {
	chain *libmarkov.Chain
}

// NewModel godoc
func NewModel(order int) Model {
	return Model{
		chain: libmarkov.NewChain(order),
	}
}
