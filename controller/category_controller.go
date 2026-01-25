package controller

import (
	"bagus-category-api/controller/utils"
	"bagus-category-api/model"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetCategory(w http.ResponseWriter, r *http.Request) {
	utils.WriteResponse(w, http.StatusOK, "Succes Get All Data", model.ListCategory)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	var categories model.Category
	err := json.NewDecoder(r.Body).Decode(&categories)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Request", nil)
		return
	}

	categories.ID = len(model.ListCategory) + 1
	model.ListCategory = append(model.ListCategory, categories)

	utils.WriteResponse(w, http.StatusCreated, "Data Berhasil Ditambahkan", categories)
}

func GetCategoryById(w http.ResponseWriter, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	for _, category := range model.ListCategory {
		if category.ID == id {
			utils.WriteResponse(w, http.StatusOK, "Category successfully found", category)
			return
		}
	}
	utils.WriteResponse(w, http.StatusNotFound, "Category Not Found", nil)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	var updateCategory model.Category
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&updateCategory)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Request", nil)
		return
	}

	for i := range model.ListCategory {
		if model.ListCategory[i].ID == id {
			updateCategory.ID = id
			model.ListCategory[i] = updateCategory
			utils.WriteResponse(w, http.StatusOK, "Data successfully changed", nil)
			return
		}
	}

	utils.WriteResponse(w, http.StatusNotFound, "Category Not Found", nil)
}

func DeleteCategory(w http.ResponseWriter, idStr string) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	for i, p := range model.ListCategory {
		if p.ID == id {
			model.ListCategory = append(model.ListCategory[:i], model.ListCategory[i+1:]...)
			utils.WriteResponse(w, http.StatusOK, "Data deleted successfully", nil)
			return
		}
	}

	utils.WriteResponse(w, http.StatusNotFound, "Category Not Found", nil)
}
