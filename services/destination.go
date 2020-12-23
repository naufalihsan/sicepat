package services

import (
	v1 "github.com/naufalihsan/sicepat/models/v1"
	"net/http"
)

func GetDestinationList(request *v1.GetDestinationListRequest) (*v1.GetDestinationListResponse, *error) {
	response := &v1.GetDestinationListResponse{}
	header := &http.Header{}

	err := Call(
		header,
		request,
		response,
	)

	if err != nil {
		panic(err)
	}

	return response, nil

}
