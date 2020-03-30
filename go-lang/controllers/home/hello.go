package home

import "net/http"

func Hello(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User Index"))
}
