package app

type App struct {
	container *Container
}

func NewApp() *App {
	return &App{
		container: NewContainer(),
	}
}

func (a *App) Run() {
	runGraphQLServer(a.container.GetResolver())
}
