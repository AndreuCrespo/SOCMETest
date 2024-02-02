function enviarDatos() {
    const numEmpleados = parseInt(document.getElementById('numEmpleados').value, 10);
    const facturacion = parseInt(document.getElementById('facturacion').value, 10);
    const actividad = document.getElementById('actividad').value;

    fetch('/empresa', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            num_empleados: numEmpleados,
            facturacion: facturacion,
            actividad: actividad,
        }),
    })
    .then(response => response.json())
    .then(data => {
        const formattedResponse = JSON.stringify(data, null, 2).replace(/\n/g, '<br>').replace(/ /g, '&nbsp;');
        document.getElementById('responseOutput').innerHTML = formattedResponse; // Usar innerHTML para interpretar el HTML
    })
    .catch((error) => {
        console.error('Error:', error);
    });
}

document.getElementById('numEmpleados').oninput = function() {
    document.getElementById('numEmpleadosValue').textContent = this.value;
}
document.getElementById('facturacion').oninput = function() {
    document.getElementById('facturacionValue').textContent = this.value;
}
