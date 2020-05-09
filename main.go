package main

import (
	data "grom_echo_demo/data"
	"grom_echo_demo/http"
)

func main() {
	data.InitSql()
	http.Start()
}
