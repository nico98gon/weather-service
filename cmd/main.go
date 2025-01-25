package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Servicio de clima en funcionamiento!")
	})

	fmt.Println("Servicio de clima escuchando en el puerto 8083...")
	if err := http.ListenAndServe(":8083", nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
