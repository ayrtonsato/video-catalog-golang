package helpers

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/protocols"
)

func HTTPCreated(body interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 201,
		Body: body,
		Error: nil,
	}
}

func HTTPOk(body interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 200,
		Body: body,
		Error: nil,
	}
}

func HTTPInternalError() protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 500,
		Body: nil,
		Error: errors.New("Internal Server Error"),
	}
}

func HTTPBadRequestError(err error) protocols.HttpResponse {
	return protocols.HttpResponse {
		Code: 400,
		Body: nil,
		Error: err,
	}
}