package coinbase

import (
	"github.com/chrhlnd/dynjson"
)

type Base struct {
	props dynjson.DynNode
}

func (obj Base) id() string {
	id, _ := obj.props.Node("/id")
	return id.AsStr()
}
