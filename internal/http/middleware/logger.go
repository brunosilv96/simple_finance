package http

import (
	"log"
	"net/http"
	"time"
)

// Logger é um middleware que registra detalhes da requisição
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Passa a requisição para o próximo handler (ou para o Mux)
		next.ServeHTTP(w, r)

		// Após a execução, logamos o resultado
		log.Printf(
			"%s %s | %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
