package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type User struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Age    uint8  `json:"age"`
	Gender string `json:"gender"`
}

var users = []User{
	{1, "James Doe", 34, "M"},
	{2, "Jane Doe", 32, "F"},
	{3, "Karen Doe", 32, "F"},
	{4, "Peter Doe", 21, "M"},
}

func getLeaveEndpoints(user User) map[string]string {
	var availablLeaves = map[string]string{
		"Sick":               fmt.Sprintf("/leaves/sick/%d", user.ID),
		"Sick with half pay": fmt.Sprintf("/leaves/sick-with-half-pay/%d", user.ID),
		"Compassionate":      fmt.Sprintf("/leaves/compassionate/%d", user.ID),
	}
	if user.Gender == "F" {
		availablLeaves["Maternity"] = fmt.Sprintf("/leaves/maternity/%d", user.ID)
	} else {
		availablLeaves["Paternity"] = fmt.Sprintf("/leaves/paternity/%d", user.ID)
	}
	return availablLeaves
}

func makeResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(data)
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	var response []map[string]any
	for _, user := range users {
		response = append(response, map[string]any{
			"user":  user,
			"links": getLeaveEndpoints(user),
		})
	}
	makeResponse(w, http.StatusOK, response)
}

func main() {
	var router = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "OPTIONS"},
		MaxAge:         300,
	}))

	router.Get("/users", getUsers)

	_ = http.ListenAndServe(":8080", router)
}
