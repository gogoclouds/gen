package body

type PageRequest struct {
	PageNum  uint `json:"page_num"`
	PageSize int  `json:"page_size"`
}

func (r PageRequest) ToOffset() int {
	pNum := r.PageNum
	if r.PageNum == 0 {
		pNum = 1
	}
	return (int(pNum) - 1) * r.PageSize
}

func (r PageRequest) ToLimit() int {
	if r.PageSize == 0 {
		return 5
	}
	return r.PageSize
}
