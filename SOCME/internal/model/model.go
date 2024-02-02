package model

// EmpresaInfo representa la información básica de una empresa
type EmpresaInfo struct {
	NumEmpleados int     `json:"num_empleados"`
	Facturacion  float64 `json:"facturacion"`
	Actividad    string  `json:"actividad"`
}

// Presupuesto representa el presupuesto generado para una empresa
type Presupuesto struct {
	Categoria   string  `json:"categoria"`
	Monto       float64 `json:"monto"`
	Explicacion string  `json:"explicacion"`
}
