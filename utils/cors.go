package utils

import "net/http"

func SetupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}
