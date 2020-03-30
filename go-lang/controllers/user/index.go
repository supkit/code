package user

import "net/http"

func Index(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("User Index"))
}
