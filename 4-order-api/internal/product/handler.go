package product

import (
	"4-order-api/pkg/request"
	"4-order-api/pkg/responce"
	"fmt"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func RegisterProductHandler(router *http.ServeMux, repo *ProductRepository) {
	h := &ProductHandler{
		ProductRepository: repo,
	}
	router.HandleFunc("POST /product", h.CreateHandler())
	router.HandleFunc("GET /product", h.GetAllHandler())
	router.HandleFunc("GET /product/{id}", h.GetByIdHandler())
	router.HandleFunc("DELETE /product/{id}", h.DeleteByIdHandler())
	router.HandleFunc("PUT /product/{id}", h.UpdateByIdHandler())
}

func (h *ProductHandler) CreateHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, err := request.HandleBody[Product](&w, r)
		if err != nil {
			responce.JsonErrorResponce(w, err.Error(), http.StatusBadRequest)
			return
		}
		p, err = h.ProductRepository.Create(p)
		if err != nil {
			responce.JsonErrorResponce(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responce.JsonOkResponce(w, p, http.StatusCreated)
	}
}

func (h *ProductHandler) GetAllHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h.ProductRepository.GetAll()
		if err != nil {
			responce.JsonErrorResponce(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responce.JsonOkResponce(w, res, http.StatusOK)
	}
}

func (h *ProductHandler) GetByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := getId(r)
		if !ok {
			responce.JsonErrorResponce(w, "Wrong product ID", http.StatusBadRequest)
			return
		}
		p := h.getProductById(w, id)
		if p == nil {
			return
		}
		responce.JsonOkResponce(w, p, http.StatusOK)
	}
}

func (h *ProductHandler) DeleteByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := getId(r)
		if !ok {
			responce.JsonErrorResponce(w, "Wrong product ID", http.StatusBadRequest)
			return
		}
		err := h.ProductRepository.DeleteById(id)
		if err != nil {
			responce.JsonErrorResponce(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responce.JsonOkResponce(w, nil, http.StatusOK)
	}
}

func (h *ProductHandler) UpdateByIdHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, ok := getId(r)
		if !ok {
			responce.JsonErrorResponce(w, "Wrong product ID", http.StatusBadRequest)
			return
		}
		newProd, err := request.HandleBody[ProductToUpdate](&w, r)
		if err != nil {
			responce.JsonErrorResponce(w, err.Error(), http.StatusBadRequest)
			return
		}
		p := h.getProductById(w, id)
		if p == nil {
			return
		}
		updateProduct(p, newProd)
		err = h.ProductRepository.Update(id, p)
		if err != nil {
			responce.JsonErrorResponce(w, err.Error(), http.StatusInternalServerError)
			return
		}
		responce.JsonOkResponce(w, nil, http.StatusOK)
	}
}

func getId(r *http.Request) (uint, bool) {
	idStr := r.PathValue("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if id > 0 {
		return uint(id), true
	}
	return 0, false
}

func (h *ProductHandler) getProductById(w http.ResponseWriter, id uint) *Product {
	p, err := h.ProductRepository.GetById(id)
	if err != nil {
		responce.JsonErrorResponce(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	if p == nil {
		responce.JsonErrorResponce(w, "Record not found", http.StatusNotFound)
		return nil
	}
	return p
}

func updateProduct(prod *Product, updated *ProductToUpdate) {
	fmt.Println(prod)
	if updated.Name != nil {
		prod.Name = *updated.Name
	}
	if updated.Description != nil {
		prod.Description = *updated.Description
	}
	if updated.Price != nil {
		prod.Price = *updated.Price
	}
	fmt.Println(prod)
}
