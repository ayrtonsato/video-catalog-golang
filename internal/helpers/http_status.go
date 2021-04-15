package helpers

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/protocols"
)

func HTTPOk(body interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 200,
		Body: body,
	}
}

func HTTPCreated(body interface{}) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 201,
		Body: body,
	}
}

func HTTPOkNoContent() protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 204,
		Body: nil,
	}
}

func HTTPBadRequestError(err error) protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 400,
		Body: err,
	}
}

func HTTPNotFound() protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 404,
		Body: errors.New("Not Found"),
	}
}

func HTTPInternalError() protocols.HttpResponse {
	return protocols.HttpResponse{
		Code: 500,
		Body: errors.New("Internal Server Error"),
	}
}
