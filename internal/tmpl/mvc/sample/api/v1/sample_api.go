package v1

type SampleListRequest struct{}

type SampleListResponse struct {
	Total int `json:"total"`
}

type SampleRequest struct {
	ID int64 `json:"id"`
}

type SampleResponse struct {
}

type SampleCreateRequest struct {
}

type SampleCreateResponse struct {
}

type SampleUpdateRequest struct {
	ID int64 `json:"id"`
}

type SampleUpdateResponse struct {
}

type SampleDeleteRequest struct {
	ID int64 `json:"id"`
}

type SampleDeleteResponse struct {
}
