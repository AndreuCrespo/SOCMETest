package handler

import (
	"SOCME/internal/service"
	"net/http"
	"os"
)

func DescargarPDFHandler(w http.ResponseWriter, r *http.Request) {
	nombreArchivo, err := service.GenerarPDF(service.UltimaExplicacionChatGPT) // Usar la última explicación para el PDF
	if err != nil {
		http.Error(w, "Error al generar el PDF", http.StatusInternalServerError)
		return
	}

	defer os.Remove(nombreArchivo) // Limpieza del archivo PDF

	w.Header().Set("Content-Disposition", "attachment; filename=respuesta_chatgpt.pdf")
	w.Header().Set("Content-Type", "application/pdf")
	http.ServeFile(w, r, nombreArchivo)
}
