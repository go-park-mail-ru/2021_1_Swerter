package index

import (
	"net/http"

	"../../utils"
)

func index(w http.ResponseWriter, r *http.Request) {
	utils.SetupCORS(&w)
	w.Write([]byte("INDEX"))
}
