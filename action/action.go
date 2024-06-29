package action

import "alertmanager/types"

var actionMap = make(ActionLut)

func GetActionMap() *ActionLut {
	return &actionMap
}

type ActionLut map[string]func(types.Action, map[string]interface{}) error

func (flut ActionLut) Add(fname string, f func(types.Action, map[string]interface{}) error) {
	flut[fname] = f
}
