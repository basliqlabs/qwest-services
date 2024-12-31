package envelope

// Response represents the standard API response envelope
type Response struct {
	// Indicates if the request was successful
	Success bool `json:"success" example:"true"`

	// Contains the response data
	Data any `json:"data,omitempty"`

	// Contains error details if Success is false
	Error *ResponseError `json:"error,omitempty"`

	// Additional metadata about the response
	Meta *Meta `json:"meta,omitempty"`

	// Pagination information if applicable
	Pagination *Pagination `json:"pagination,omitempty"`
}

// ResponseError represents error details
type ResponseError struct {
	// Error code identifier
	// @example INVALID_INPUT
	Code ErrorCode `json:"code"`

	// Human-readable error message
	// @example Invalid input provided
	Message string `json:"message"`

	// Field-specific validation errors
	Fields map[string]string `json:"fields,omitempty"`
}

type Meta map[string]any

type Pagination struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
	Total      int `json:"total,omitempty"`
}

func New(success bool) *Response {
	return &Response{
		Success: success,
	}
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) WithMeta(meta *Meta) *Response {
	r.Meta = meta
	return r
}

func (r *Response) WithPagination(pagination *Pagination) *Response {
	r.Pagination = pagination
	return r
}

func (r *Response) WithError(err *ResponseError) *Response {
	r.Error = err
	return r
}
