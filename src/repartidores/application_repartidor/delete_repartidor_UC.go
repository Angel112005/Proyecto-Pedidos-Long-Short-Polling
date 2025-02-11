package applicationrepartidor

import domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"

type DeleteRepartidorUseCase struct {
	repository domainrepartidor.RepartidorRepository
}

func NewDeleteRepartidorUseCase(repository domainrepartidor.RepartidorRepository) *DeleteRepartidorUseCase {
	return &DeleteRepartidorUseCase{repository: repository}
}

func (useCase *DeleteRepartidorUseCase) Execute(id int) error {
	return useCase.repository.Delete(id)
}