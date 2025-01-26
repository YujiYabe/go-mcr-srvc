package usecase

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
	toGatewayDB ToGatewayDB,
	toGatewayExternal ToGatewayExternal,
) ToUseCase {
	return &useCase{
		ToDomain:          toDomain,
		ToGatewayDB:       toGatewayDB,
		ToGatewayExternal: toGatewayExternal,
	}
}

type (
	// useCase ...
	useCase struct {
		ToDomain          ToDomain
		ToGatewayDB       ToGatewayDB
		ToGatewayExternal ToGatewayExternal
	}
)
