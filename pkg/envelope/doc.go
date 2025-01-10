package envelope

// These structs are meant to be used in OpenAPI comments to
// indicate the resulting type of different responses

type OpenAPIResponseSuccess struct {
	Success bool `json:"success" example:"true"`
	Data    any  `json:"data,omitempty"`
}

type OpenAPIResponseSuccessMeta struct {
	Success bool  `json:"success" example:"true"`
	Data    any   `json:"data,omitempty"`
	Meta    *Meta `json:"meta,omitempty"`
}

type OpenAPIResponseSuccessPagination struct {
	Success    bool        `json:"success" example:"true"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type OpenAPIResponseSuccessPaginationMeta struct {
	Success    bool        `json:"success" example:"true"`
	Data       any         `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
	Meta       *Meta       `json:"meta,omitempty"`
}

type OpenAPIResponseError struct {
	Success bool          `json:"success" example:"false"`
	Error   ResponseError `json:"error,omitempty" extensions:"x-nullable"`
}
