package api

import (
	"log/slog"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
)

const logoRoute = "/ctf1.png"
const logoFilePath = "./static/ctf1.png"

// Hosts some information that doesn't change while app is running.
// Some logo or fonts, for example
func static(r chi.Router, logger *slog.Logger) {
	logoBytes, readLogoErr := os.ReadFile(logoFilePath)

	r.HandleFunc(logoRoute, func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if readLogoErr != nil {
				logger.Error("Failed to read image",
					slog.String("error", readLogoErr.Error()),
					slog.String("image_path", logoFilePath))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Write(logoBytes)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
