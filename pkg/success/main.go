package pkg_success

import "fmt"

type ClientSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type PaginationData struct {
	Message string `json:"message"`
	Meta    *Meta  `json:"meta"`
	Data    any    `json:"data"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	TotalData   int `json:"total_data"`
	PerPage     int `json:"per_page"`
}

func NewMetaData(currentPage int, totalPage int, perPage int, totalData int) *Meta {
	return &Meta{
		CurrentPage: currentPage,
		TotalPage:   totalPage,
		TotalData:   totalData,
		PerPage:     perPage,
	}
}

func SuccessGetData(data any) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}

func SuccessPaginationData(data any, meta *Meta) *PaginationData {
	return &PaginationData{
		Message: "Success",
		Meta:    meta,
		Data:    data,
	}
}

func SuccessDeleteData(id any) *ClientSuccess {
	return &ClientSuccess{
		Message: fmt.Sprintf("Success delete data with ID %v", id),
		Data:    id,
	}
}

func SuccessUpdateData(id any, data any) *ClientSuccess {
	return &ClientSuccess{
		Message: fmt.Sprintf("Success update data with ID %v", id),
		Data:    data,
	}
}

func SuccessCreateData(data any) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}
