package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicializar el repositorio
	repository := &InMemoryRepository{}
	// Inicializar el servicio
	service := &PersonService{repository: repository}
	// Inicializar el controlador
	controller := &Controller{service: service}
	// Configurar el router
	r := mux.NewRouter()
	r.Use(corsMiddleware)
	r.HandleFunc("sorteoamigos.vercel.app/people", controller.AddPerson).Methods(http.MethodPost)
	r.HandleFunc("sorteoamigos.vercel.app/grid", controller.GetGrid).Methods(http.MethodGet)
	// Iniciar el servidor
	http.ListenAndServe(":8080", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("CORS middleware executed")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}
