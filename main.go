package main

import (
	"log"
	"net/http"

	"gplaydb/db"
	"gplaydb/internal/handler"
	"gplaydb/internal/repositories"
	"gplaydb/internal/services"
)

func main() {
	db.Connect()
	defer db.Close()

	repo := repositories.NewUserRepository(db.DB)
	service := services.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	http.HandleFunc("PUT /users/{id}", handler.UpdateUser)
	http.HandleFunc("GET /users/{id}/products", handler.UserWithProducts)
	http.HandleFunc("DELETE /users/{id}", handler.DeleteUserById)
	http.HandleFunc("POST /users", handler.InsertUser)
	http.HandleFunc("GET /users/{id}", handler.GetUserById)
	http.HandleFunc("GET /users", handler.GetAllUsers)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
