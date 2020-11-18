package main

import (
	"errserv/errserv"
)

func main() {
	//timeout 0.000125s
	errserv.Main(125000)
}
