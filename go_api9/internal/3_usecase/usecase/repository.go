package usecase

type (
	// UseCase ...
	UseCase struct {
		ToDomain  ToDomain
		ToService ToService
	}

	// ToService ...
	ToService interface {
		Dummy() error
	}

	// ToDomain ...
	ToDomain interface {
		Dummy() error
	}
)
