package applicationrepartidor

import domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"

type ViewAllRepartidoresUseCase struct {
	repository domainrepartidor.RepartidorRepository
}

func NewViewAllRepartidoresUseCase(repository domainrepartidor.RepartidorRepository) *ViewAllRepartidoresUseCase {
	return &ViewAllRepartidoresUseCase{repository: repository}
}

func (useCase *ViewAllRepartidoresUseCase) Execute() ([]*domainrepartidor.Repartidor, error) {
	return useCase.repository.ViewAll()
}