package handlers

import (
	"bagus-category-api/handlers/utils"
	"bagus-category-api/model"
	"bagus-category-api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (handler *CategoryHandler) HandleCategory(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetCategory(w, r)
	case http.MethodPost:
		handler.AddCategory(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *CategoryHandler) HandleCategoryById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handler.GetCategoryById(w, r)
	case http.MethodPut:
		handler.UpdateCategory(w, r)
	case http.MethodDelete:
		handler.DeleteCategory(w, r)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func (handler *CategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := handler.service.GetCategory()
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.WriteResponse(w, http.StatusOK, "Succes Get All Data", categories)
}

func (handler *CategoryHandler) AddCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Request", nil)
		return
	}

	err = handler.service.AddCategory(&category)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Request", nil)
		return
	}

	utils.WriteResponse(w, http.StatusCreated, "Data Berhasil Ditambahkan", category)
}

func (handler *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	category, err := handler.service.GetCategoryById(id)
	if err != nil {
		utils.WriteResponse(w, http.StatusNotFound, "Category Not Found", nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Category successfully found", category)
}

func (handler *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
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

	updateCategory.ID = id
	err = handler.service.UpdateCategory(&updateCategory)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Data successfully changed", nil)
}

func (handler *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteResponse(w, http.StatusBadRequest, "Invalid Category ID", nil)
		return
	}

	err = handler.service.DeleteCategory(id)
	if err != nil {
		utils.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, "Data deleted successfully", nil)
}
