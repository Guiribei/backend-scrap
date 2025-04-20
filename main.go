package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Guiribei/backend-scrap/data"
)

func certHandler(w http.ResponseWriter, r *http.Request) {
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

	if !strings.Contains(req.Email, "@") {
		http.Error(w, "Email inválido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string {
		"status": "certidão gerada com sucesso",
	})
}

func main() {
	http.HandleFunc("/cert", certHandler)
}