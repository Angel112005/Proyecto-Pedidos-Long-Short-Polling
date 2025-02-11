

// ✅ Short Polling para obtener los pedidos y contar pendientes
async function obtenerPedidos() {
    try {
        let response = await fetch("http://localhost:3000/pedido");
        if (response.ok) {
            let pedidos = await response.json();
            mostrarPedidos(pedidos.pedidos);
            contarPedidosPendientes(); // Llama a la función para actualizar el contador
        }
    } catch (error) {
        console.error("❌ Error obteniendo pedidos:", error);
    }
    setTimeout(obtenerPedidos, 5000); // Short Polling cada 5 segundos
}

//  Mostrar pedidos en la interfaz
function mostrarPedidos(pedidos) {
    const pedidosDiv = document.getElementById("pedidos");
    pedidosDiv.innerHTML = "";
    pedidos.forEach(pedido => {
        let p = document.createElement("div");
        p.className = "pedido-card";
        p.textContent = `Pedido #${pedido.id} - Estado: ${pedido.estado}`;
        pedidosDiv.appendChild(p);
    });
}

// ✅ Contador de pedidos pendientes
async function contarPedidosPendientes() {
    try {
        let response = await fetch("http://localhost:3000/pedido/pendientes");
        if (response.ok) {
            let data = await response.json();
            document.getElementById("contador-pendientes").textContent = `Pedidos pendientes: ${data.pedidos_pendientes}`;
        }
    } catch (error) {
        console.error("❌ Error contando pedidos pendientes:", error);
    }
}

// ✅ Long Polling para obtener repartidores disponibles
async function obtenerRepartidores() {
    try {
        let response = await fetch("http://localhost:3000/repartidores");
        if (response.ok) {
            let repartidores = await response.json();
            mostrarRepartidores(repartidores.repartidores);
        } else if (response.status === 204) {
            console.log("⏳ No hay cambios en los repartidores.");
        }
    } catch (error) {
        console.error("❌ Error obteniendo repartidores:", error);
    }
}

// ✅ Mostrar repartidores disponibles en la interfaz 
function mostrarRepartidores(repartidores) {
    const repartidoresDiv = document.getElementById("repartidores");
    repartidoresDiv.innerHTML = "";
    repartidores.forEach(rep => {
        let p = document.createElement("div");
        p.className = "repartidor-card";
        p.textContent = `🚴 Repartidor: ${rep.nombre} - Disponible: ${rep.disponible ? "✅ Sí" : "❌ No"}`;
        repartidoresDiv.appendChild(p);
    });
}

// ✅ Función para crear un nuevo pedido (simulación)
async function crearPedido() {
    try {
        let response = await fetch("http://localhost:3000/pedido", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ cliente_id: 1, estado: "pendiente" })
        });
        if (response.ok) {
            console.log("✅ Pedido creado con éxito");
            contarPedidosPendientes(); // Actualizar contador tras creación
        }
    } catch (error) {
        console.error("❌ Error creando pedido:", error);
    }
}

// ✅ Iniciar Short Polling y Long Polling
obtenerPedidos();
setInterval(obtenerRepartidores, 10000); // Long Polling cada 10 segundos
setInterval(contarPedidosPendientes, 5000); // Actualizar contador de pedidos cada 5 segundos
