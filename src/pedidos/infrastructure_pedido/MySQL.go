package infrastructurepedido

import (
	"log"
	"lab-test.com/module/src/core"
	domainpedido "lab-test.com/module/src/pedidos/domain_pedido"
)

type MySQLPedido struct {
	conn *core.Conn_MySQL
}

func NewMySQLPedido() *MySQLPedido {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("❌ Error al configurar la conexión a MySQL: %v", conn.Err)
	}
	return &MySQLPedido{conn: conn}
}

func (mysql *MySQLPedido) Create(pedido *domainpedido.Pedido) error {
	query := "INSERT INTO pedidos (cliente_id, estado) VALUES (?, ?)"
	result, err := mysql.conn.DB.Exec(query, pedido.ClienteID, pedido.Estado)
	if err != nil {
		log.Println("❌ Error insertando pedido:", err)
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("❌ Error obteniendo ID del nuevo pedido:", err)
		return err
	}
	log.Println("✅ Pedido creado con ID:", lastID)
	return nil
}

// ✅ Obtener todos los pedidos
func (mysql *MySQLPedido) ViewAll() ([]*domainpedido.Pedido, error) {
	query := "SELECT id, cliente_id, estado FROM pedidos"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Println("❌ Error consultando pedidos:", err)
		return nil, err
	}
	defer rows.Close()

	var pedidos []*domainpedido.Pedido
	for rows.Next() {
		var pedido domainpedido.Pedido
		if err := rows.Scan(&pedido.ID, &pedido.ClienteID, &pedido.Estado); err != nil {
			log.Println("❌ Error leyendo fila:", err)
			continue
		}
		pedidos = append(pedidos, &pedido)
	}
	log.Println("📌 Pedidos obtenidos con éxito")
	return pedidos, nil
}

// ✅ Obtener un pedido por ID
func (mysql *MySQLPedido) ViewById(id int) (*domainpedido.Pedido, error) {
	query := "SELECT id, cliente_id, estado FROM pedidos WHERE id = ?"
	var pedido domainpedido.Pedido
	row := mysql.conn.DB.QueryRow(query, id)
	if err := row.Scan(&pedido.ID, &pedido.ClienteID, &pedido.Estado); err != nil {
		log.Println("❌ Error obteniendo pedido por ID:", err)
		return nil, err
	}
	log.Println("✅ Pedido obtenido con éxito:", pedido)
	return &pedido, nil
}

// ✅ Actualizar un pedido
func (mysql *MySQLPedido) Update(pedido *domainpedido.Pedido) error {
	query := "UPDATE pedidos SET estado = ? WHERE id = ?"
	result, err := mysql.conn.DB.Exec(query, pedido.Estado, pedido.ID)
	if err != nil {
		log.Println("❌ Error actualizando pedido:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("❌ Error obteniendo número de filas afectadas:", err)
		return err
	}
	log.Printf("✅ Pedido actualizado con éxito (%d filas afectadas)\n", rowsAffected)
	return nil
}

// ✅ Eliminar un pedido
func (mysql *MySQLPedido) Delete(id int) error {
	query := "DELETE FROM pedidos WHERE id = ?"
	result, err := mysql.conn.DB.Exec(query, id)
	if err != nil {
		log.Println("❌ Error eliminando pedido:", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("❌ Error obteniendo número de filas afectadas:", err)
		return err
	}
	log.Printf("✅ Pedido eliminado con éxito (%d filas afectadas)\n", rowsAffected)
	return nil
}

// ✅ Función para contar pedidos con estado "pendiente"
func (mysql *MySQLPedido) ContarPendientes() (int, error) {
    query := "SELECT COUNT(*) FROM pedidos WHERE estado = 'pendiente'"
    var count int
    err := mysql.conn.DB.QueryRow(query).Scan(&count)
    if err != nil {
        return 0, err
    }
    return count, nil
}
