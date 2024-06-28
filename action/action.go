package action

import "alertmanager/utils"

type Action struct {
	ActionName string `yaml:"action_name"`
	ActionArgs string `yaml:"action_args"`
}

func GetDefaultAction() Action {
	return Action{ActionName: "NOOP_ACTION", ActionArgs: "ARG1,ARG2"}
}

var actionMap = make(utils.FunctionLut)

func GetActionMap() *utils.FunctionLut {
	return &actionMap
}
