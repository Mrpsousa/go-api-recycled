package usecase

import (
	"github.com/api-go/internal/entity"
)

// Orquestrador (terceira parte)
// a aplicacao nao precisa saber do dado na totalidade, so eh precisa transitar dados assencias pela app, por issoo uso de DTO
type ListProductOutputDto struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ListProductUseCase struct {
	ProductRepository entity.ProductRepository
}

// funcao construtora
func NewListProductUseCase(productRepository entity.ProductRepository) *ListProductUseCase {
	return &ListProductUseCase{ProductRepository: productRepository}
}

func (u *ListProductUseCase) Execute() ([]*ListProductOutputDto, error) {
	products, err := u.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	var productsOutPut []*ListProductOutputDto
	for _, product := range products {
		productsOutPut = append(productsOutPut, &ListProductOutputDto{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsOutPut, nil
}
