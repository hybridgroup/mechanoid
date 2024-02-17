package engine

type Engine struct {
	Interpreter interpreter
	FileStore   filestore
	Bridge      bridge
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Init() error {
	if err := e.Interpreter.Init(); err != nil {
		return err
	}

	if err := e.FileStore.Init(); err != nil {
		return err
	}

	for _, d := range e.Bridge.Devices() {
		if err := d.Init(); err != nil {
			return err
		}
	}
	return nil
}
