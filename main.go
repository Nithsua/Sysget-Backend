package main

import (
	"github.com/austinmarner/sysget/server"
	"github.com/austinmarner/sysget/top"
)

var topInfo top.Top

func main() {
	server.InitiateServer(topInfo)
}
