package service

import (
	"fmt"
	"time"

	"github.com/lucas-rech/sisinfo-ecommerce/internal/domain"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/dto"
	"github.com/lucas-rech/sisinfo-ecommerce/internal/repository"
)

// Análogo a uma classe que implementa a interface ProductService
type productService struct {
	productRepo repository.ProductRepository
}

// Análogo ao construtor
func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{productRepo: productRepo}
}

func (s *productService) CreateProduct(product dto.ProductCreateRequest) error {
	if product.Name == "" || product.Price <= 0 {
		return fmt.Errorf("invalid product data")
	}
	productEntity := domain.Product{
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		ImageURL:    product.ImageURL,
		Category:    product.Category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return s.productRepo.Create(&productEntity)
}

func (s *productService) FindProductByID(id uint) (*domain.Product, error) {
	if id == 0 {
		return nil, fmt.Errorf("invalid product ID")
	}

	product, err := s.productRepo.FindByID(&id)
	if err != nil {
		return nil, fmt.Errorf("error finding product: %w", err)
	}
	if product == nil {
		return nil, fmt.Errorf("product not found")
	}

	return product, nil
}

func (s *productService) FindAllProducts() ([]domain.Product, error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving products: %w", err)
	}
	if len(products) == 0 {
		return nil, fmt.Errorf("no products found")
	}

	return products, nil
}

// FETCH Update
func (s *productService) UpdateProduct(product dto.ProductUpdateRequest, id *uint) error {
	if id == nil || *id == 0 {
		return fmt.Errorf("invalid product ID")
	}
	existingProduct, err := s.productRepo.FindByID(id)
	if err != nil {
		return fmt.Errorf("error finding product by ID: %w", err)
	}
	if existingProduct == nil {
		return fmt.Errorf("product not found")
	}

	//Atualiza apenas os campos que foram fornecidos no request
	if product.Name != nil {
		existingProduct.Name = *product.Name
	}
	if product.Description != nil {
		existingProduct.Description = *product.Description
	}
	if product.Price != nil {
		existingProduct.Price = *product.Price
	}
	if product.Stock != nil {
		existingProduct.Stock = *product.Stock
	}
	if product.ImageURL != nil {
		existingProduct.ImageURL = *product.ImageURL
	}
	if product.Category != nil {
		existingProduct.Category = *product.Category
	}
	existingProduct.UpdatedAt = time.Now()

	return s.productRepo.Update(existingProduct)

}

func (s *productService) DeleteProduct(id uint) error {
	if id == 0 {
		return fmt.Errorf("invalid product ID")
	}

	err := s.productRepo.Delete(&id)
	if err != nil {
		return fmt.Errorf("error deleting product: %w", err)
	}

	return nil
}
