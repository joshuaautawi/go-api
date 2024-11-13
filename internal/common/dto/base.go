package dto

type (
	Response[T any] struct {
		Data   T      `json:"data"`
		Status string `json:"status"`
		Meta   *Meta  `json:"meta,omitempty"`
		Error  *Error `json:"error,omitempty"`
	}
	Error struct {
		Code         int    `json:"code"`
		Message      string `json:"message"`
		InternalCode string `json:"internalCode"`
	}
	Meta struct {
		TotalCount int64 `json:"totalCount,omitempty"`
		Page       int   `json:"page,omitempty"`
		Limit      int   `json:"limit,omitempty"`
	}

	GetOneByIdRequest struct {
		ID int `json:"id" validate:"required"`
	}
	GetAllRequest struct {
		Page  int `json:"page" validate:"required"`
		Limit int `json:"limit" validate:"required"`
	}
)
