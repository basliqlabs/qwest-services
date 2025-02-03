package envelope

// Response represents the standard API response envelope
type Response struct {
	Success    bool           `json:"success" example:"true"`
	Data       any            `json:"data,omitempty"`
	Error      *ResponseError `json:"error,omitempty"`
	Meta       *Meta          `json:"meta,omitempty"`
	Pagination *Pagination    `json:"pagination,omitempty"`
}

// ResponseError represents error details
type ResponseError struct {
	Code    ErrorCode         `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"`
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
