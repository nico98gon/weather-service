package main

import (
	"fmt"
	"net/http"
	"nilus-challenge-backend/internal/infrastructure"
)

func main() {
	infrastructure.StartHTTPServer()

	fmt.Println("Servicio de clima escuchando en el puerto 8083...")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
