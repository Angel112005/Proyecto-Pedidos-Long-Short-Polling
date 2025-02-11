package applicationrepartidor

import domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"

type ViewRepartidorByIdUseCase struct {
	repository domainrepartidor.RepartidorRepository
}

func NewViewRepartidorByIdUseCase(repository domainrepartidor.RepartidorRepository) *ViewRepartidorByIdUseCase {
	return &ViewRepartidorByIdUseCase{repository: repository}
}

func (useCase *ViewRepartidorByIdUseCase) Execute(id int) (*domainrepartidor.Repartidor, error) {
	return useCase.repository.ViewById(id)
}
