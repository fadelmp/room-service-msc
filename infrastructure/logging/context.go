package logging

type Layer string

const (
	LayerRepository Layer = "repository"
	LayerUsecase    Layer = "usecase"
)

type Context struct {
	layer Layer
}
