package main

import (
	"fiber-web/cmd"
)

// @title server
// @version 1.0
// @description server token格式:Bearer token

// @securityDefinitions.apikey jwt
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}

