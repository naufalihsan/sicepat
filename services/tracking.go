package services

import (
	v1 "github.com/naufalihsan/sicepat/models/v1"
	"net/http"
)

func Tracking(request *v1.TrackingRequest) (*v1.TrackingResponse, *error) {
	response := &v1.TrackingResponse{}
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
