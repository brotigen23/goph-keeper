package app

type App struct{}

func New() *App {
	// Services and JWT storage
	return &App{}
}

func (a App) Run() {
}
