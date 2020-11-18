package main

import (
	"err_server/errserv"
)

func main() {
	//timeout 0.000125s
	errserv.Main(125000)
}
