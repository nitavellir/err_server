package errserv

import (
	"net/http"
)

type Port uint16
type Errcode uint16
type Timeout uint16

//port =>error_code
type ErrorPortPair map[Port]Errcode

type Options struct {
	ErrorsPorts ErrorPortPair
}

func (opt *Options) IsEnableErrorServ() bool {
	return len(opt.ErrorsPorts) > 0
}

func NewOptions() *Options {
	return &Options{
		ErrorsPorts: make(ErrorPortPair),
	}
}

func parseFlags() *Options {
	// TODO: parse config file or args
	opt := NewOptions()
	var baseErrorPort uint16 = 10000
	stdErrors := []uint16{
		//4xx
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusMethodNotAllowed,
		http.StatusNotAcceptable,
		http.StatusGone,
		http.StatusLengthRequired,
		http.StatusPreconditionFailed,
		http.StatusRequestEntityTooLarge,
		http.StatusRequestURITooLong,
		http.StatusUnsupportedMediaType,
		http.StatusRequestedRangeNotSatisfiable,

		//5xx
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout,
	}

	for _, errcode := range stdErrors {
		opt.ErrorsPorts[Port(baseErrorPort+uint16(errcode))] = Errcode(errcode)
	}

	return opt
}
