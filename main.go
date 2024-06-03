package main

import (
	"net/http"

	"github.com/uadmin/uadmin"
	"github.com/vincentroyce/html_templating/views"
)

func main() {
	uadmin.Register()

	http.HandleFunc("/", uadmin.Handler(views.TestTemplatingHandler))
	uadmin.RootURL = "/admin"
	uadmin.StartServer()
}
