package pkg_success

import "fmt"

type ClientSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type PaginationData struct {
	Message     string `json:"message"`
	Data        any    `json:"data"`
	CurrentPage int    `json:"current_page"`
	TotalPage   int    `json:"total_page"`
	TotalData   int    `json:"total_data"`
	PerPage     int    `json:"per_page"`
}

func SuccessGetData(data any) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}

func SuccessPaginationData(data any, currentPage int, totalPage int, perPage int, totalData int) *PaginationData {
	return &PaginationData{
		Message:     "Success",
		Data:        data,
		CurrentPage: currentPage,
		TotalData:   totalData,
		TotalPage:   totalPage,
		PerPage:     perPage,
	}
}

func SuccessDeleteData(id any) *ClientSuccess {
	return &ClientSuccess{
		Message: fmt.Sprintf("Success delete data with ID %v", id),
		Data:    id,
	}
}

func SuccessCreateData(data any) *ClientSuccess {
	return &ClientSuccess{
		Message: "Success",
		Data:    data,
	}
}
