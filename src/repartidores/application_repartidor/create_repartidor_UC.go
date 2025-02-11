package applicationrepartidor

import domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"

type CreateRepartidorUseCase struct {
	repository domainrepartidor.RepartidorRepository
}

func NewCreateRepartidorUseCase(repository domainrepartidor.RepartidorRepository) *CreateRepartidorUseCase {
	return &CreateRepartidorUseCase{repository: repository}
}

func (useCase *CreateRepartidorUseCase) Execute(nombre string, disponible bool) error {
	repartidor := domainrepartidor.NewRepartidor(nombre, disponible)
	return useCase.repository.Create(repartidor)
}