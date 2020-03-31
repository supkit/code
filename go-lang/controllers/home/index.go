package home

import "net/http"

func Index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Home Index"))
}
