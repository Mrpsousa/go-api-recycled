package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/api-go/internal/usecase"
)

type ProductHandlers struct {
	CreateProductUseCase *usecase.CreateProductUseCase
	ListProductsUseCase  *usecase.ListProductUseCase
}

func NewProductHandlers(createProductUseCase *usecase.CreateProductUseCase, listProductsUseCase *usecase.ListProductUseCase) *ProductHandlers {
	return &ProductHandlers{
		CreateProductUseCase: createProductUseCase,
		ListProductsUseCase:  listProductsUseCase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := p.CreateProductUseCase.Execute((input))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_ := fmt.Sprintf("ERRO Create: %v", err)
		fmt.Println(error_)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

func (p *ProductHandlers) ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	output, err := p.ListProductsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error_ := fmt.Sprintf("ERRO List: %v", err)
		fmt.Println(error_)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
