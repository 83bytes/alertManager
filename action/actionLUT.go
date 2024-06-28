package action

type ActionLut map[string]func(Action, map[string]interface{}) error

func (flut ActionLut) Add(fname string, f func(Action, map[string]interface{}) error) {
	flut[fname] = f
}
