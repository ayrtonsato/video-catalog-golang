package protocols

type Controller interface {
	handle() HttpResponse
}
