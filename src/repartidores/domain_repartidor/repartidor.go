package domainrepartidor

type Repartidor struct {
	ID         int    `json:"id"`
	Nombre     string `json:"nombre"`
	Disponible bool   `json:"disponible"`
}

func NewRepartidor(nombre string, disponible bool) *Repartidor {
	return &Repartidor{
		Nombre:     nombre,
		Disponible: disponible,
	}
}