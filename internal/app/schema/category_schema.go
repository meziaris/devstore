package schema

type GetCategoryResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateCategoryReq struct {
	Name        string `form:"name" json:"name" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
}

type UpdateCategoryReq struct {
	Name        string `form:"name" json:"name" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
}
