package schema

import "mime/multipart"

type CreateProductReq struct {
	Name        string                `validate:"required" form:"name"`
	Description string                `validate:"required" form:"description"`
	Currency    string                `validate:"required" form:"currency"`
	TotalStock  int                   `validate:"required,number" form:"total_stock"`
	IsActive    bool                  `validate:"required,boolean" form:"is_active"`
	CategoryID  int                   `validate:"required,number" form:"category_id"`
	Image       *multipart.FileHeader `validate:"required,omitempty" form:"image"`
}

type CreateProductResp struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
	TotalStock  int    `json:"total_stock"`
	IsActive    bool   `json:"is_active"`
	CategoryID  int    `json:"category_id"`
	Image       string `json:"image"`
}

type BrowseProductResp struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	TotalStock  int     `json:"total_stock"`
	IsActive    bool    `json:"is_active"`
	ImageURL    *string `json:"image_url"`
}

type DetailProductResp struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Currency    string   `json:"currency"`
	TotalStock  int      `json:"total_stock"`
	IsActive    bool     `json:"is_active"`
	Category    Category `json:"category"`
	ImageURL    *string  `json:"image_url"`
}

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateProductReq struct {
	Name        string                `validate:"required" form:"name"`
	Description string                `validate:"required" form:"description"`
	Currency    string                `validate:"required" form:"currency"`
	TotalStock  int                   `validate:"required,number" form:"total_stock"`
	IsActive    bool                  `validate:"required,boolean" form:"is_active"`
	CategoryID  int                   `validate:"required,number" form:"category_id"`
	Image       *multipart.FileHeader `validate:"required,omitempty" form:"image"`
}

type UpdateProductResp struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Currency    string `form:"currency"`
	TotalStock  int    `form:"total_stock"`
	IsActive    bool   `form:"is_active"`
	CategoryID  int    `form:"category_id"`
	Image       string `form:"image"`
}
