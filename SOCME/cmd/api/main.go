package main

import (
	"SOCME/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// API Handler para procesar la información de la empresa
	r.HandleFunc("/empresa", handler.EmpresaHandler).Methods("POST")
	r.HandleFunc("/descargar-pdf", handler.DescargarPDFHandler).Methods("GET")
	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Página de inicio")
	// })

	// Configuración para servir archivos estáticos desde el directorio /public
	staticFileDirectory := http.Dir("./public/")
	staticFileHandler := http.FileServer(staticFileDirectory)
	r.PathPrefix("/").Handler(http.StripPrefix("/", staticFileHandler))

	log.Println("Servidor iniciado en el puerto 8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
