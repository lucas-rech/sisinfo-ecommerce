package repository

import (
	"fmt"

	"github.com/lucas-rech/sisinfo-ecommerce/backend/internal/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) Create(product *domain.Product) error {
	if product == nil {
		return fmt.Errorf("product cannot be nil")
	}

	return p.db.Create(product).Error
}

func (p *productRepository) Delete(id *uint) error {
	if id == nil {
		return fmt.Errorf("invalid product ID for deletion")
	}

	return p.db.Delete(&domain.Product{}, id).Error
}

func (p *productRepository) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	err := p.db.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *productRepository) FindByID(id *uint) (*domain.Product, error) {
	if id == nil {
		return nil, fmt.Errorf("invalid product ID")
	}
	
	var product domain.Product
	err := p.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) Update(product *domain.Product) error {
	if product == nil || product.ID == 0 {
		return fmt.Errorf("product to update cannot be nil and mus have a valid ID")
	}

	result := p.db.Save(product)

	if result.Error != nil {
		return fmt.Errorf("error updating product: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no product found with ID %d to update", product.ID)
	}

	return nil
}


