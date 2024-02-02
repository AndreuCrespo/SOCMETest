package handler

import (
	"SOCME/internal/model"
	"SOCME/internal/service"
	"encoding/json"
	"net/http"
)

func EmpresaHandler(w http.ResponseWriter, r *http.Request) {
	var info model.EmpresaInfo
	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "Error al decodificar solicitud", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	explicacion, err := service.GenerarExplicacionConOpenAI(ctx, info)
	if err != nil {
		http.Error(w, "Error al generar explicación", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"explicacion": explicacion})
}
