package helpers

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/protocols"
)

func HTTPCreated(body interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 201,
		Body: body,
	}
}

func HTTPOk(body interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 200,
		Body: body,
	}
}

func HTTPInternalError() protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 500,
		Body: errors.New("Internal Server Error"),
	}
}

func HTTPBadRequestError(err error) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 400,
		Body: err,
	}
}
