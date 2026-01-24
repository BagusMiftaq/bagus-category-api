package controller

import (
	"bagus-category-api/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ListCategory)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var categories model.Category
	err := json.NewDecoder(r.Body).Decode(&categories)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	categories.ID = len(model.ListCategory) + 1
	model.ListCategory = append(model.ListCategory, categories)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categories)
}

func GetCategoryById(w http.ResponseWriter, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for _, category := range model.ListCategory {
		if category.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	var updateCategory model.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	for i := range model.ListCategory {
		if model.ListCategory[i].ID == id {
			updateCategory.ID = id
			model.ListCategory[i] = updateCategory
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "success",
				"message": "Data berhasil diubah",
			})
			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, p := range model.ListCategory {
		if p.ID == id {
			model.ListCategory = append(model.ListCategory[:i], model.ListCategory[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"status":  "success",
				"message": "Data berhasil dihapus",
			})

			return
		}
	}

	http.Error(w, "Category Not Found", http.StatusNotFound)
}
