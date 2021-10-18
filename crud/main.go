package main

import (
	"github.com/rastasi/learn-golang/crud/domain"
	"github.com/rastasi/learn-golang/crud/http"
)

func main() {
	domain := domain.NewDomain()
	http.StartHttpServer(domain)
}
