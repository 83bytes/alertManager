package action

import "alertmanager/types"

var actionMap = make(ActionLut)

func GetActionMap() *ActionLut {
	return &actionMap
}

type ActionFunc func(types.Alert, types.Action, map[string]interface{}) error

type ActionLut map[string]ActionFunc

func (flut ActionLut) Add(fname string, f ActionFunc) {
	flut[fname] = f
}
