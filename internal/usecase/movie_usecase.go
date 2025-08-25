package usecase

type (
	MovieUsecase struct {}
)

func (mu *MovieUsecase) GetMovieByName(name string) (*Movie, error) {
	return mu.movieRepo.FindByName(name)
}