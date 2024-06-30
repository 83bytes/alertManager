package types

type Action struct {
	StepName   string `yaml:"step_name"`
	ActionName string `yaml:"action_name"`
	ActionArgs string `yaml:"action_args"`
}

func GetDefaultAction() Action {
	return Action{
		StepName:   "ACTION_STEP_1",
		ActionName: "NOOP_ACTION",
		ActionArgs: "ARG1,ARG2"}
}
