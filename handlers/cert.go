package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/Guiribei/backend-scrap/data"
	"github.com/Guiribei/backend-scrap/utils"
)

func CertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req data.CertRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	req.CNPJ = utils.SanitizeCNPJ(req.CNPJ)
	req.Email = utils.SanitizeEmail(req.Email)

	if len(req.CNPJ) != 14 {
		http.Error(w, "CNPJ deve ter exatamente 14 dígitos", http.StatusBadRequest)
		return
	}

	if !utils.IsValidEmail(req.Email) {
		http.Error(w, "Email inválido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string {
		"status": "certidão gerada com sucesso",
	})
}