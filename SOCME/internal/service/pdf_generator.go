package service

import (
	"os"
	"path/filepath"

	"github.com/jung-kurt/gofpdf"
)

func GenerarPDF(texto string) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)
	_, lineHt := pdf.GetFontSize()
	pdf.MultiCell(190, lineHt, texto, "", "L", false)

	// Genera un nombre de archivo único para cada PDF
	nombreArchivo := filepath.Join(os.TempDir(), "Respuesta_ChatGPT.pdf")
	err := pdf.OutputFileAndClose(nombreArchivo)
	if err != nil {
		return "", err
	}
	return nombreArchivo, nil
}
