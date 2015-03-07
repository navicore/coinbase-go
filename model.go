package coinbase

import (
	"github.com/chrhlnd/dynjson"
)

type Model struct {
	Base
	props dynjson.DynNode
}

func (obj Model) id() string {
	id, _ := obj.props.Node("/id")
	return id.AsStr()
}
