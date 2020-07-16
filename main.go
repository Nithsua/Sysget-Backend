package main

import (
	"github.com/austinmarner/system_admin_backend/server"
	"github.com/austinmarner/system_admin_backend/top"
)

var topInfo top.Top

func main() {
	server.InitiateServer(topInfo)
}
