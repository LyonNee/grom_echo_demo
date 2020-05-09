package main

import (
	data "github.com/LyonNee/grom_echo_demo/data"
	"github.com/LyonNee/grom_echo_demo/http"
)

func main() {
	data.InitSql()
	http.Start()
}
