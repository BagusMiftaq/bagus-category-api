package main

import (
	"bagus-category-api/controller"
	"bagus-category-api/controller/utils"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	//GET localhost:8080/categories/{id}
	//PUT localhost:8080/categories/{id}
	//DELETE localhost:8080/categories/{id}
	http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/categories/")

		switch r.Method {
		case "GET":
			controller.GetCategoryById(w, idStr)
		case "PUT":
			controller.UpdateCategory(w, r, idStr)
		case "DELETE":
			controller.DeleteCategory(w, idStr)
		default:
			return
		}
	})

	//GET localhost:8080/categories
	//POST localhost:8080/categories
	http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controller.GetCategory(w, r)
		} else {
			controller.AddCategory(w, r)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.WriteResponse(w, http.StatusOK, "Server is Running, please use the path /categories and /categories{id}", nil)
		w.Header().Set("Content-Type", "application/json")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server Failed")
	}
}
