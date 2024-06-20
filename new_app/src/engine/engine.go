package engine

type Engine struct {
	path string
}

func (e *Engine) Run() error {
	return nil
}

func New(path string) (*Engine, error) {
	eng := &Engine{
		path: path,
	}
	return eng, nil
}
