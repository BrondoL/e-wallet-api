package utils

import (
	"os"
	"strconv"
)

type SResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type EResponse struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func SuccessResponse(message string, code int, data interface{}) SResponse {
	return SResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "success",
		},
		Data: data,
	}
}

func ErrorResponse(message string, code int, err interface{}) EResponse {
	return EResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "error",
		},
		Error: err,
	}
}

type Link struct {
	URL    string
	Label  string
	Active bool
}

type PaginationResponse struct {
	CurrentPage  int         `json:"current_page"`
	Data         interface{} `json:"data"`
	FirstPageUrl string      `json:"first_page_url"`
	From         int         `json:"from"`
	LastPage     int         `json:"last_page"`
	LastPageUrl  string      `json:"last_page_url"`
	NextPageUrl  string      `json:"next_page_url"`
	Path         string      `json:"path"`
	PerPage      int         `json:"per_page"`
	PrevPageUrl  string      `json:"prev_page_url"`
	To           int         `json:"to"`
	Total        int         `json:"total"`
}

type Metadata struct {
	Resource string
	TotalAll int
	TotalNow int
	Page     int
	Limit    int
	Sort     string
}

func ResponseWithPagination(message string, code int, data interface{}, metadata Metadata) SResponse {
	APP_URL := os.Getenv("APP_URL")
	url := APP_URL + "/" + metadata.Resource
	totalPage := (metadata.TotalAll / metadata.Limit) + 1
	var next string
	var prev string
	if metadata.Page == 1 {
		prev = ""
	} else {
		prev = url + "?page=" + strconv.Itoa(metadata.Page-1)
	}
	if metadata.Page == totalPage {
		next = ""
	} else {
		next = url + "?page=" + strconv.Itoa(metadata.Page+1)
	}

	var from int
	var to int
	if metadata.TotalNow != 0 {
		from = ((metadata.Page - 1) * metadata.Limit) + 1
		to = from + metadata.TotalNow - 1
	}

	return SResponse{
		Meta: Meta{
			Message: message,
			Code:    code,
			Status:  "success",
		},
		Data: PaginationResponse{
			CurrentPage:  metadata.Page,
			Data:         data,
			FirstPageUrl: url + "?page=1",
			From:         from,
			LastPage:     totalPage,
			LastPageUrl:  url + "?page=" + strconv.Itoa(totalPage),
			NextPageUrl:  next,
			Path:         url,
			PerPage:      metadata.Limit,
			PrevPageUrl:  prev,
			To:           to,
			Total:        metadata.TotalAll,
		},
	}
}
