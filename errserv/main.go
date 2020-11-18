package errserv

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	bindFailed  = 0
	bindSuccess = 1
)

type PortHandler interface {
	Port() Port
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type OneErrorHandler struct {
	port  Port
	error Errcode
}

func (handler *OneErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	error := int(handler.error)
	log.Printf("[REQUEST :%d] %s %s => %d", handler.port, request.Method, request.RequestURI, handler.error)
	http.Error(writer, fmt.Sprintf("Error %d - %s", error, http.StatusText(error)), error)
}

func (handler *OneErrorHandler) Port() Port {
	return handler.port
}

func Main(args ...interface{}) {
	//args
	//timeout
	var timeOut time.Duration = 0
	if firstArgInt, ok := args[0].(int); ok {
		timeOut = time.Duration(int64(firstArgInt))
	}

	opt := parseFlags()

	bindChan := make(chan int)
	bind := func(handler PortHandler, timeOut time.Duration) {
		port := handler.Port()
		if timeOut > 0 {
			http.Handle(fmt.Sprintf("/%d", port), http.TimeoutHandler(handler, timeOut, "TIMEOUT!!!!!!\n"))
		} else {
			http.Handle(fmt.Sprintf("/%d", port), handler)
		}
		err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
		if err != nil {
			log.Printf("[WARN] Port %d skipping: %s", port, err)
			bindChan <- bindFailed
		} else {
			bindChan <- bindSuccess
		}
	}

	if !opt.IsEnableErrorServ() && !opt.IsEnableTimeoutServ() {
		log.Fatal("Nothing to do")
	}

	needBinded := 0
	if opt.IsEnableErrorServ() {
		for port, errcode := range opt.ErrorsPorts {
			needBinded++
			go bind(&OneErrorHandler{error: errcode, port: port}, timeOut)
		}
	}

	anyBind := false
	var res int
	for i := 0; i < needBinded; i++ {
		res = <-bindChan
		anyBind = anyBind || res == bindSuccess
	}
	if !anyBind {
		log.Fatal("All ports are skipped")
	}
}
