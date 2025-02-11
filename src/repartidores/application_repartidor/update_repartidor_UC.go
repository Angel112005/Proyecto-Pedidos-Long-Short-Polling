package applicationrepartidor

import domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"

type UpdateRepartidorUseCase struct {
	repository domainrepartidor.RepartidorRepository
}

func NewUpdateRepartidorUseCase(repository domainrepartidor.RepartidorRepository) *UpdateRepartidorUseCase {
	return &UpdateRepartidorUseCase{repository: repository}
}

func (useCase *UpdateRepartidorUseCase) Execute(repartidor *domainrepartidor.Repartidor) error {
	return useCase.repository.Update(repartidor)
}