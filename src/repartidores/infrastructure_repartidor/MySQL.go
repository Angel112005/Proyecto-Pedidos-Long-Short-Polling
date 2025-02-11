package infrastructurerepartidor

import (
	"log"
	"lab-test.com/module/src/core"
	domainrepartidor "lab-test.com/module/src/repartidores/domain_repartidor"
)

type MySQLRepartidor struct {
	conn *core.Conn_MySQL
}

func NewMySQLRepartidor() *MySQLRepartidor {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("❌ Error al configurar la conexión a MySQL: %v", conn.Err)
	}
	return &MySQLRepartidor{conn: conn}
}

// ✅ Crear un repartidor en la base de datos
func (mysql *MySQLRepartidor) Create(repartidor *domainrepartidor.Repartidor) error {
	query := "INSERT INTO repartidores (nombre, disponible) VALUES (?, ?)"
	result, err := mysql.conn.DB.Exec(query, repartidor.Nombre, repartidor.Disponible)
	if err != nil {
		log.Println("❌ Error insertando repartidor:", err)
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("❌ Error obteniendo ID del nuevo repartidor:", err)
		return err
	}
	log.Println("✅ Repartidor creado con ID:", lastID)
	return nil
}

// ✅ Obtener todos los repartidores
func (mysql *MySQLRepartidor) ViewAll() ([]*domainrepartidor.Repartidor, error) {
	query := "SELECT id, nombre, disponible FROM repartidores"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Println("❌ Error consultando repartidores:", err)
		return nil, err
	}
	defer rows.Close()

	var repartidores []*domainrepartidor.Repartidor
	for rows.Next() {
		var repartidor domainrepartidor.Repartidor
		if err := rows.Scan(&repartidor.ID, &repartidor.Nombre, &repartidor.Disponible); err != nil {
			log.Println("❌ Error leyendo fila:", err)
			continue
		}
		repartidores = append(repartidores, &repartidor)
	}
	log.Println("📌 Repartidores obtenidos con éxito")
	return repartidores, nil
}

// ✅ Obtener un repartidor por ID
func (mysql *MySQLRepartidor) ViewById(id int) (*domainrepartidor.Repartidor, error) {
	query := "SELECT id, nombre, disponible FROM repartidores WHERE id = ?"
	var repartidor domainrepartidor.Repartidor
	row := mysql.conn.DB.QueryRow(query, id)
	if err := row.Scan(&repartidor.ID, &repartidor.Nombre, &repartidor.Disponible); err != nil {
		log.Println("❌ Error obteniendo repartidor por ID:", err)
		return nil, err
	}
	log.Println("✅ Repartidor obtenido con éxito:", repartidor)
	return &repartidor, nil
}

// ✅ Actualizar un repartidor
func (mysql *MySQLRepartidor) Update(repartidor *domainrepartidor.Repartidor) error {
	query := "UPDATE repartidores SET nombre = ?, disponible = ? WHERE id = ?"
	result, err := mysql.conn.DB.Exec(query, repartidor.Nombre, repartidor.Disponible, repartidor.ID)
	if err != nil {
		log.Println("❌ Error actualizando repartidor:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("❌ Error obteniendo número de filas afectadas:", err)
		return err
	}
	log.Printf("✅ Repartidor actualizado con éxito (%d filas afectadas)\n", rowsAffected)
	return nil
}

// ✅ Eliminar un repartidor
func (mysql *MySQLRepartidor) Delete(id int) error {
	query := "DELETE FROM repartidores WHERE id = ?"
	result, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Println("❌ Error eliminando repartidor:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("❌ Error obteniendo número de filas afectadas:", err)
		return err
	}
	log.Printf("✅ Repartidor eliminado con éxito (%d filas afectadas)\n", rowsAffected)
	return nil
}
