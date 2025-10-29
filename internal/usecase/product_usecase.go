package usecase

import (
	"context"
	"golang-redis/internal/delivery/http/request"
	"golang-redis/internal/delivery/http/response"
	"golang-redis/internal/entity"
	"golang-redis/internal/repository"
)

type ProductUseCase struct {
	Repository *repository.ProductRepository
}

func NewProductUseCase(repository *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{Repository: repository}
}

func (uc *ProductUseCase) CreateProduct(ctx context.Context, name string, price float64, category string) ([]response.ProductResponse, error) {
	product, _ := uc.Repository.Create(ctx, name, price, category)

	var result []response.ProductResponse
	result = append(result, response.ProductResponse{
		ID:       uint64(product.ID),
		Name:     product.Name,
		Price:    product.Price,
		Category: product.Category,
	})

	return result, nil
}

func (uc *ProductUseCase) CreateProductBatch(ctx context.Context, productPayload []request.ProductRequest) ([]response.ProductResponse, error) {
	products := make([]entity.Product, 0, len(productPayload))
	for _, p := range productPayload {
		products = append(products, entity.Product{
			Name:     p.Name,
			Price:    p.Price,
			Category: p.Category,
		})
	}

	if err := uc.Repository.CreateBatch(ctx, products); err != nil {
		return nil, err
	}

	productResponses := make([]response.ProductResponse, 0, len(products))
	for _, p := range products {
		productResponses = append(productResponses, response.ProductResponse{
			ID:       uint64(p.ID),
			Name:     p.Name,
			Price:    p.Price,
			Category: p.Category,
		})
	}

	return productResponses, nil
}

func (uc *ProductUseCase) GetAllProducts(ctx context.Context) ([]response.ProductResponse, error) {
	products, err := uc.Repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var result []response.ProductResponse
	for _, p := range products {
		result = append(result, response.ProductResponse{
			ID:       uint64(p.ID),
			Name:     p.Name,
			Price:    p.Price,
			Category: p.Category,
		})
	}

	return result, nil
}

func (uc *ProductUseCase) GetProductByID(ctx context.Context, id int) (*entity.Product, error) {
	product, err := uc.Repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil

}

func (uc *ProductUseCase) GetProductByCategory(ctx context.Context, productCategory string) ([]entity.Product, error) {
	products, err := uc.Repository.GetByCategory(ctx, productCategory)
	if err != nil {
		return nil, err
	}

	return products, nil

}
