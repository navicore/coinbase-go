package coinbase

import (
	"encoding/json"

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

func Id(id string) dynjson.DynNode {
	type idprop struct {
		Id string `json:"id"`
	}
	myid := idprop{id}
	b, _ := json.Marshal(myid)
	props := dynjson.NewFromBytes(b)
	return props
}
