package action

type Action struct {
	ActionName string `yaml:"action_name"`
	ActionArgs string `yaml:"action_args"`
}

func GetDefaultAction() Action {
	return Action{ActionName: "NOOP_ACTION", ActionArgs: "ARG1,ARG2"}
}

var actionMap = make(ActionLut)

func GetActionMap() *ActionLut {
	return &actionMap
}
