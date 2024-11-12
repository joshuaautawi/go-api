package dto

type (
	Response[T any] struct {
		Data   T      `json:"data"`
		Status string `json:"status"`
		Meta   *Meta  `json:"meta,omitempty"`
		Error  *Error `json:"error,omitempty"`
	}
	Error struct {
		Code         string `json:"code"`
		Message      string `json:"message"`
		InternalCode string `json:"internalCode"`
	}
	GetOneById struct {
		ID int
	}
	GetAllRequest struct {
		Page  int
		Limit int
	}
	Meta struct {
		TotalCount int64 `json:"totalCount,omitempty"`
		Page       int   `json:"page,omitempty"`
		Limit      int   `json:"limit,omitempty"`
	}

	GetAllBaseResponse[T any] struct {
		Data T     `json:"data,omitempty"`
		Meta *Meta `json:"meta,omitempty"`
	}
)
