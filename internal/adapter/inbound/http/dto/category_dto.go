package dto

import "kasir-api/internal/domain"

type CreateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateCategoryRequest) ToDomain() domain.Category {
	return domain.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}

func (r *UpdateCategoryRequest) ToDomain() domain.Category {
	return domain.Category{
		Name:        r.Name,
		Description: r.Description,
	}
}

func FromDomainCategory(c domain.Category) CategoryResponse {
	return CategoryResponse{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
