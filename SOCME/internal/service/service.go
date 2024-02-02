package service

import (
	"SOCME/internal/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GenerarExplicacionConOpenAI(ctx context.Context, info model.EmpresaInfo) (string, error) {
	// Ajustando la estructura de la solicitud para el endpoint de chat
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": "Explica los siguientes datos de la empresa."},
			{"role": "user", "content": fmt.Sprintf("Número de empleados: %d, Facturación: %.2f, Actividad: %s.", info.NumEmpleados, info.Facturacion, info.Actividad)},
		},
	})
	if err != nil {
		log.Printf("Error al serializar los datos de entrada para OpenAI: %v", err)
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(requestBody))
	if err != nil {
		log.Printf("Error al preparar la solicitud HTTP para OpenAI: %v", err)
		return "", err
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error al realizar la solicitud a OpenAI: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Verificación del código de estado HTTP
	if resp.StatusCode >= 400 {
		responseBody, _ := io.ReadAll(resp.Body)
		log.Printf("Respuesta inesperada de OpenAI, código de estado: %d, estado: %s, cuerpo: %s", resp.StatusCode, resp.Status, string(responseBody))
		return "", fmt.Errorf("error %d: %s", resp.StatusCode, resp.Status)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Printf("Error al deserializar la respuesta de OpenAI: %v", err)
		return "", err
	}

	// Procesamiento de la respuesta
	if choices, found := response["choices"].([]interface{}); found && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if messages, exists := choice["message"].(map[string]interface{}); exists {
				if text, exists := messages["content"].(string); exists {
					return text, nil
				}
			}
		}
	}

	log.Println("Nota: OpenAI no devolvió una explicación.")
	return "No se pudo generar una explicación.", nil
}
