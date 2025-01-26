package usecase

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
	toDBGateway ToDBGateway,
	toExternalGateway ToExternalGateway,
) ToUseCase {
	return &useCase{
		ToDomain:          toDomain,
		ToDBGateway:       toDBGateway,
		ToExternalGateway: toExternalGateway,
	}
}

type (
	// useCase ...
	useCase struct {
		ToDomain          ToDomain
		ToDBGateway       ToDBGateway
		ToExternalGateway ToExternalGateway
	}
)
