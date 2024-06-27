package utils

// A lookup table for name -> function
// Functions are pretty generic.
// Accepts a string of args (todo: make this more generic)
// there are no timeouts yet (todo)
// each function is supposed to handle its own arg structure and how to parse it.
// todo
// there has to be a better way to do this with interfaces and gorotines and channels

type FunctionLut map[string]func(string) (string, error)

func (flut FunctionLut) Add(fname string, f func(string) (string, error)) FunctionLut {
	flut[fname] = f

	return flut
}
